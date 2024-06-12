package app

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net"
	"net/http"

	"github.com/bufbuild/protovalidate-go"
	v1team "github.com/ecojuntak/laklak/gen/go/v1/team"
	"github.com/ecojuntak/laklak/internal/team"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"gorm.io/gorm"

	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
)

type app struct {
	grpcServer *grpc.Server
	httpServer *runtime.ServeMux
	db         *gorm.DB
}

func New(db *gorm.DB) app {
	grpcServer := grpc.NewServer(
		grpc.StatsHandler(otelgrpc.NewServerHandler()),
	)
	reflection.Register(grpcServer)

	validator, err := protovalidate.New()
	if err != nil {
		log.Fatal(err)
	}

	teamRepository := team.NewRepository(db)
	v1team.RegisterTeamServiceServer(grpcServer, &team.Server{Repository: teamRepository, Validator: validator})

	return app{
		grpcServer: grpcServer,
		httpServer: runtime.NewServeMux(),
		db:         db,
	}
}

func (a app) StartGrpcServer(port string) {
	err := setupOTelSDK(context.Background())
	if err != nil {
	}

	address := fmt.Sprintf("0.0.0.0:%s", port)
	slog.Info(fmt.Sprintf("grpc app start on %s", address))

	if l, err := net.Listen("tcp", address); err != nil {
		slog.Error(fmt.Sprintf("error in listening on %s", address), "err", err)
	} else {
		if err := a.grpcServer.Serve(l); err != nil {
			slog.Error("unable to start grpcServer", "err", err)
		}
	}
}

func (a app) StartHTTPServer(grpcPort, httpPort string) {
	grpcAddress := fmt.Sprintf("0.0.0.0:%s", grpcPort)
	err := v1team.RegisterTeamServiceHandlerFromEndpoint(context.Background(), a.httpServer, grpcAddress, []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())})
	if err != nil {
		slog.Error("error when registering team service handler", "err", err)
		panic(err)
	}

	httpAddress := fmt.Sprintf("0.0.0.0:%s", httpPort)
	server := &http.Server{
		Addr:    httpAddress,
		Handler: a.httpServer,
	}

	slog.Info(fmt.Sprintf("http app start on %s", httpAddress))
	if err = server.ListenAndServe(); err != nil {
		slog.Error("error when starting http app", "err", err)
	}
}
