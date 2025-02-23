package apiserver

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	petname "yadro.com/course/internal/petname"
	petnamepb "yadro.com/course/proto"
)

var (
	namer = petname.NewNamer()
)

type StreamFlowCtrl struct {
	busy        atomic.Int32
	terminate   atomic.Bool //TODO: а нужен ли тут атомик 🤔
	cancelMux   sync.Mutex  //NOTE: overkill, but need to control cancel funcs
	cancelFuncs map[string]context.CancelFunc
}

func (s *grpcServer) Ping(_ context.Context, in *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, nil
}

func (s *grpcServer) GenerateMany(r *petnamepb.PetnameStreamRequest, stream petnamepb.PetnameGenerator_GenerateManyServer) error {
	s.flowCtrl.busy.Add(1)

	ctx, cancel := context.WithTimeout(stream.Context(), s.maxDuration)

	cancelKey := fmt.Sprintf("%p", cancel)
	if cancel != nil {
		s.addCancel(cancelKey, cancel)
	}

	nameChan, err := namer.GenerateMany(ctx, int(r.Words), r.Separator, int(r.Names))
	if err != nil { //NOTE: Only custom error(apiserver.InvalidArgument could apear here)
		return status.Error(codes.InvalidArgument, err.Error())
	}
	defer func() {
		cancel()
		s.removeCancel(cancelKey)
	}()
	defer s.flowCtrl.busy.Add(-1)

	for {
		if s.flowCtrl.terminate.Load() {
			return status.Error(codes.Aborted, "Server is shutting down...")
		}
		select {
		//NOTE: ctx.Done() вызовется и при отключении пользователя во время ответа
		// и при условии, что время закончилось

		case <-ctx.Done():
			log.Warn("Context is closed, stopping iteration...")
			if ctx.Err() == context.DeadlineExceeded {
				return status.Error(codes.DeadlineExceeded, "Timeout exceeded")
			}
			return status.Error(codes.Canceled, "Context cancelled") //NOTE: в случае, если контекст закрыт пользователем

		case val, ok := <-nameChan:
			if !ok {
				log.Info("nameChan is closed, stopping iteration")
				return nil
			}
			if err := stream.Send(&petnamepb.PetnameResponse{Name: val}); err != nil {
				log.Error(fmt.Sprintf("failed to send response: %v", err))
				return err
			}
		}
	}
}

// func (s *grpcServer) addCancelAtomic(c context.CancelFunc) {
// 	s.flowCtrl.cancelMux.Lock()
// 	s.flowCtrl.cancel = append(s.flowCtrl.cancel, c)
// 	s.flowCtrl.cancelMux.Unlock()
// }

func (s *grpcServer) addCancel(key string, c context.CancelFunc) {
	s.flowCtrl.cancelMux.Lock()
	s.flowCtrl.cancelFuncs[key] = c
	s.flowCtrl.cancelMux.Unlock()
}

// Функция для удаления CancelFunc из мапы по ключу
func (s *grpcServer) removeCancel(key string) {
	s.flowCtrl.cancelMux.Lock()
	defer s.flowCtrl.cancelMux.Unlock()

	delete(s.flowCtrl.cancelFuncs, key)
}

func (s *grpcServer) Generate(_ context.Context, r *petnamepb.PetnameRequest) (*petnamepb.PetnameResponse, error) {
	log.Info(r.String())
	name, err := namer.GenerateName(int(r.Words), r.Separator)
	if err != nil { //NOTE: Only custom error(apiserver.InvalidArgument could apear here)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return &petnamepb.PetnameResponse{Name: name}, nil //NOTE: fan fact, status.Error(codes.OK, "panic") equal to nil)

}
