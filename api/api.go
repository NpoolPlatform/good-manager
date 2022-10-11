package api

import (
	"context"
	"github.com/NpoolPlatform/good-manager/api/appgood"
	"github.com/NpoolPlatform/message/npool/good/mgr/v1"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	v1.UnimplementedManagerServer
}

func Register(server grpc.ServiceRegistrar) {
	v1.RegisterManagerServer(server, &Server{})
	appgood.Register(server)
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	if err := v1.RegisterManagerHandlerFromEndpoint(context.Background(), mux, endpoint, opts); err != nil {
		return err
	}
	return nil
}