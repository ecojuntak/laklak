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

func (r repository) Teams(ctx context.Context) (teams []*v1team.Team, err error) {
	ctx, span := tracer.Start(ctx, "repository.Teams")
	defer span.End()

	result := r.db.WithContext(ctx).Find(&teams)
	return teams, result.Error
}

func (r repository) Team(ctx context.Context, id int32) (team *v1team.Team, err error) {
	ctx, span := tracer.Start(ctx, "repository.Team")
	defer span.End()

	result := r.db.WithContext(ctx).Find(&v1team.Team{Id: id}).First(&team)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, customError.RecordNotFoundError
	}
	return team, result.Error
}
