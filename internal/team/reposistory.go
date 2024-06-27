package team

import (
	"context"
	"errors"

	v1team "github.com/ecojuntak/laklak/gen/go/v1/team"
	customError "github.com/ecojuntak/laklak/internal/error"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) repository {
	return repository{
		db: db,
	}
}

func (r repository) Create(ctx context.Context, team *v1team.Team) (err error) {
	ctx, span := tracer.Start(ctx, "repository.Create")
	defer span.End()

	result := r.db.WithContext(ctx).Create(team)
	return result.Error
}

func (r repository) GetTeams(ctx context.Context) (teams []*v1team.Team, err error) {
	ctx, span := tracer.Start(ctx, "repository.GetTeams")
	defer span.End()

	result := r.db.WithContext(ctx).Find(&teams)
	return teams, result.Error
}

func (r repository) GetTeam(ctx context.Context, id int32) (team *v1team.Team, err error) {
	ctx, span := tracer.Start(ctx, "repository.GetTeam")
	defer span.End()

	result := r.db.WithContext(ctx).Preload("Applications").Find(&v1team.Team{Id: id}).First(&team)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, customError.RecordNotFoundError
	}
	return team, result.Error
}
