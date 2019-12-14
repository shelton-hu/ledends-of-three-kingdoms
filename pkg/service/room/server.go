package room

import (
	"context"

	"google.golang.org/grpc"

	"github.com/shelton-hu/legends-of-three-kingdoms/pkg/constants"
	"github.com/shelton-hu/legends-of-three-kingdoms/pkg/manager"
	"github.com/shelton-hu/legends-of-three-kingdoms/pkg/pb"
	"github.com/shelton-hu/legends-of-three-kingdoms/pkg/pi"
)

type Server struct {
}

func Serve() {
	s := Server{}
	manager.NewGrpcServer(constants.RoomManagerHost, constants.RoomManagerPort).
		ShowErrorCause(pi.Global().Cfg(context.Background()).Grpc.ShowErrorCause).
		WithChecker(s.Checker).
		WithBuilder(s.Builder).
		Serve(func(server *grpc.Server) {
			pb.RegisterRoomServiceServer(server, &s)
		})
}
