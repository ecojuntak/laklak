package application

import (
	"context"

	v1application "github.com/ecojuntak/laklak/gen/go/v1/application"
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

func (r repository) Create(ctx context.Context, application *v1application.Application) (err error) {
	ctx, span := tracer.Start(ctx, "application.repository.Create")
	defer span.End()

	result := r.db.WithContext(ctx).Create(application)
	return result.Error
}
