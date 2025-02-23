package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"os"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	// petname "yadro.com/course/proto"
	"yadro.com/course/internal/logger"
	petname "yadro.com/course/internal/petname"
	petnamepb "yadro.com/course/proto"
)

var log = logger.GetInstance()

var namer = petname.NewNamer()

type server struct {
	petnamepb.UnimplementedPetnameGeneratorServer
}

func (s *server) Ping(_ context.Context, in *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, nil
}

const maxDuration = 2 * time.Second

func (s *server) Generate(_ context.Context, r *petnamepb.PetnameRequest) (*petnamepb.PetnameResponse, error) {
	log.Info(r.String())
	return nil, status.Error(1, "error")
	// name := namer.GenerateName(int(r.Words), r.Separator)
	// return &petnamepb.PetnameResponse{Name: name}, status.Error(0, "error")
	// return &petnamepb.PetnameResponse{
	// Name: name,
	// }, nil
}

func (s *server) GenerateMany(r *petnamepb.PetnameStreamRequest, stream petnamepb.PetnameGenerator_GenerateManyServer) error {
	ctx, cancel := context.WithTimeout(stream.Context(), maxDuration)

	nameChan := namer.GenerateMany(ctx, int(r.Words), r.Separator, int(r.Names))
	defer cancel()

	for {
		select {
		//NOTE: ctx.Done() вызовется и при отключении пользователя во время ответа
		// и при условии, что время закончилось.
		case <-ctx.Done():
			log.Warn("Context is closed, stopping iteration")
			return status.Error(codes.DeadlineExceeded, "timeout exceeded")
			// return stream.Context().Err()
		case val, ok := <-nameChan:
			if !ok {
				log.Info("nameChan is closed, stopping iteration")
				// cancel()
				return nil
			}
			if err := stream.Send(&petnamepb.PetnameResponse{Name: val}); err != nil {
				log.Error(fmt.Sprintf("failed to send response: %v", err))
				return err
			}
		}
	}
}

func main() {
	var address string
	flag.StringVar(&address, "address", ":8081", "server address")
	flag.Parse()
	log.Info("Server is trying to start on address: 8081")
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Error(fmt.Sprintf("failed to listen: %v", err))
		os.Exit(-1)
	}

	s := grpc.NewServer()
	petnamepb.RegisterPetnameGeneratorServer(s, &server{})
	reflection.Register(s)

	if err := s.Serve(listener); err != nil {
		log.Error(fmt.Sprintf("failed to serve: %v", err))
		os.Exit(-1)
	}
}
