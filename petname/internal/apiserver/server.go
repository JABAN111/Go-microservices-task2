package apiserver

import (
	"context"
	"fmt"
	"net"
	"sync/atomic"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"yadro.com/course/internal/logger"
	petnamepb "yadro.com/course/proto"
)

var (
	log = logger.GetInstance()
)

const timeout = 10 * time.Second

type ServerManipulator interface {
	Run() error
	GracefulShutDown() error
}

type grpcServer struct {
	petnamepb.UnimplementedPetnameGeneratorServer
	port        int
	maxDuration time.Duration //NOTE: default in seconds
	flowCtrl    StreamFlowCtrl
	grpcSrv     *grpc.Server
	listener    net.Listener
}

func CreateServer(port int) *grpcServer {
	return &grpcServer{
		port:        port,
		maxDuration: timeout,
		flowCtrl: StreamFlowCtrl{
			busy:        atomic.Int32{},
			cancelFuncs: make(map[string]context.CancelFunc),
		},
		grpcSrv: grpc.NewServer(),
	}
}

func (s *grpcServer) Run() error {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", s.port))
	if err != nil {
		log.Error(fmt.Sprintf("failed to listen: %v", err))
		return err
	}
	s.listener = listener

	reflection.Register(s.grpcSrv)
	petnamepb.RegisterPetnameGeneratorServer(s.grpcSrv, s)

	log.Info(fmt.Sprintf("gRPC server is running on port %d", s.port))
	if err := s.grpcSrv.Serve(listener); err != nil {
		log.Error(fmt.Sprintf("failed to serve: %v", err))
		return err
	}
	return nil
}

func (s *grpcServer) GracefulShutDown() error {
	log.Info("Initiating graceful shutdown...")

	s.flowCtrl.terminate.Store(true)

	if s.flowCtrl.cancelFuncs != nil {
		for _, f := range s.flowCtrl.cancelFuncs {
			f()
		}
	}
	if s.flowCtrl.busy.Load() != 0 {
		log.Info(fmt.Sprintf("Waiting %d seconds for active streams to finish...", timeout))
		time.Sleep(timeout)
	} else {
		log.Info("No active streams, shutting down immediately...")
	}

	log.Info("Shutting down gRPC server gracefully...")
	s.grpcSrv.GracefulStop()

	return nil
}
