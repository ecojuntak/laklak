package application

import (
	"context"

	v1application "github.com/ecojuntak/laklak/gen/go/v1/application"
	"go.opentelemetry.io/otel"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

var tracer = otel.Tracer("go.opentelemetry.io/otel")

type Server struct {
	v1application.UnimplementedApplicationServiceServer
	Repository Repository
	Validator  Validator
}

type Repository interface {
	Create(ctx context.Context, team *v1application.Application) error
}

type Validator interface {
	Validate(msg proto.Message) error
}

func (s Server) Create(ctx context.Context, request *v1application.CreateRequest) (*v1application.CreateResponse, error) {
	ctx, span := tracer.Start(ctx, "application.server.Create")
	defer span.End()

	if err := s.Validator.Validate(request); err != nil {
		return nil, status.Error(codes.InvalidArgument, "request not valid")
	}

	err := s.Repository.Create(ctx, &v1application.Application{
		TeamId: request.TeamId,
		Name:   request.Name,
	})

	if err != nil {
		return nil, status.Error(codes.Internal, "error creating application")
	}
	return &v1application.CreateResponse{
		Message: "new application created",
	}, status.Error(codes.OK, "")
}
