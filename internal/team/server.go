package team

import (
	"context"
	"errors"

	v1team "github.com/ecojuntak/laklak/gen/go/v1/team"
	customError "github.com/ecojuntak/laklak/internal/error"
	"go.opentelemetry.io/otel"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

var tracer = otel.Tracer("go.opentelemetry.io/otel")

type Server struct {
	v1team.UnimplementedTeamServiceServer
	Repository Repository
	Validator  Validator
}

type Repository interface {
	Create(ctx context.Context, team *v1team.Team) error
	GetTeams(ctx context.Context) ([]*v1team.Team, error)
	GetTeam(ctx context.Context, id int32) (*v1team.Team, error)
}

type Validator interface {
	Validate(msg proto.Message) error
}

func (s *Server) Create(ctx context.Context, request *v1team.CreateTeamRequest) (*v1team.CreateTeamResponse, error) {
	ctx, span := tracer.Start(ctx, "server.Create")
	defer span.End()

	if err := s.Validator.Validate(request); err != nil {
		return nil, status.Error(codes.InvalidArgument, "team name is required")
	}

	err := s.Repository.Create(ctx, &v1team.Team{
		Name: request.Name,
	})

	if err != nil {
		return nil, status.Error(codes.Internal, "error creating team")
	}
	return &v1team.CreateTeamResponse{
		Message: "new team created",
	}, status.Error(codes.OK, "")
}

func (s *Server) GetTeams(ctx context.Context, request *v1team.GetTeamsRequest) (*v1team.GetTeamsResponse, error) {
	ctx, span := tracer.Start(ctx, "server.GetTeams")
	defer span.End()

	teams, err := s.Repository.GetTeams(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "error getting teams")
	}
	return &v1team.GetTeamsResponse{Teams: teams}, nil
}

func (s *Server) GetTeam(ctx context.Context, request *v1team.GetTeamRequest) (*v1team.GetTeamResponse, error) {
	ctx, span := tracer.Start(ctx, "server.GetTeam")
	defer span.End()

	if err := s.Validator.Validate(request); err != nil {
		return nil, status.Error(codes.InvalidArgument, "team ID is required")
	}

	team, err := s.Repository.GetTeam(ctx, request.Id)
	if errors.Is(err, customError.RecordNotFoundError) {
		return nil, status.Error(codes.NotFound, "team not found")
	} else if err != nil {
		return nil, status.Error(codes.Internal, "error getting team")
	}
	return &v1team.GetTeamResponse{Team: team}, status.Error(codes.OK, "")
}
