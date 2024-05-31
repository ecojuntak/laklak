package team

import (
	"context"
	"errors"

	v1team "github.com/ecojuntak/laklak/gen/go/v1/team"
	customError "github.com/ecojuntak/laklak/internal/error"
	"go.opentelemetry.io/otel"
	"google.golang.org/grpc/status"
)

var tracer = otel.Tracer("go.opentelemetry.io/otel")

type Server struct {
	v1team.UnimplementedTeamServiceServer
	Repository Repository
}

type Repository interface {
	Create(ctx context.Context, team *v1team.Team) error
	Teams(ctx context.Context) ([]*v1team.Team, error)
	Team(ctx context.Context, id int32) (*v1team.Team, error)
}

func (s *Server) Create(ctx context.Context, request *v1team.CreateTeamRequest) (*v1team.CreateTeamResponse, error) {
	ctx, span := tracer.Start(ctx, "server.Create")
	defer span.End()

	err := s.Repository.Create(ctx, &v1team.Team{
		Name: request.Name,
	})

	if err != nil {
		return nil, err
	}
	return &v1team.CreateTeamResponse{}, err
}

func (s *Server) Teams(ctx context.Context, request *v1team.GetTeamsRequest) (*v1team.GetTeamsResponse, error) {
	ctx, span := tracer.Start(ctx, "server.Teams")
	defer span.End()

	teams, err := s.Repository.Teams(ctx)
	return &v1team.GetTeamsResponse{Teams: teams}, err
}

func (s *Server) Team(ctx context.Context, request *v1team.GetTeamRequest) (*v1team.GetTeamResponse, error) {
	ctx, span := tracer.Start(ctx, "server.Team")
	defer span.End()

	team, err := s.Repository.Team(ctx, request.Id)
	if errors.Is(err, customError.RecordNotFoundError) {
		return &v1team.GetTeamResponse{}, status.Error(5, "team not found")
	}
	return &v1team.GetTeamResponse{Team: team}, err
}
