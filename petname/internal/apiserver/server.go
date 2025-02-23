package apiserver

type grpcServer struct {
	port int
}

func NewGrpcServer() *grpcServer {

	return &grpcServer{
		port: 8080, //TODO 8080
	}

}
