package team

import (
	"context"
	"errors"
	"testing"

	v1team "github.com/ecojuntak/laklak/gen/go/v1/team"
	customError "github.com/ecojuntak/laklak/internal/error"
	teamMock "github.com/ecojuntak/laklak/mocks/team"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestServer_Create(t *testing.T) {
	type fields struct {
		Repository *teamMock.MockRepository
		Validator  *teamMock.MockValidator
	}
	type args struct {
		ctx     context.Context
		request *v1team.CreateTeamRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *v1team.CreateTeamResponse
		wantErr error
		mockFn  func(aa args, ff fields)
	}{
		{
			name: "should success create team",
			fields: fields{
				Repository: new(teamMock.MockRepository),
				Validator:  new(teamMock.MockValidator),
			},
			args: args{
				ctx: context.TODO(),
				request: &v1team.CreateTeamRequest{
					Name: "test",
				},
			},
			want: &v1team.CreateTeamResponse{
				Message: "new team created",
			},
			wantErr: nil,
			mockFn: func(aa args, ff fields) {
				ff.Repository.Mock.
					On("Create", mock.Anything, &v1team.Team{Name: aa.request.Name}).
					Return(nil)

				ff.Validator.On("Validate", aa.request).Return(nil)
			},
		},
		{
			name: "should return empty response if repository fails",
			fields: fields{
				Repository: new(teamMock.MockRepository),
				Validator:  new(teamMock.MockValidator),
			},
			args: args{
				ctx: context.TODO(),
				request: &v1team.CreateTeamRequest{
					Name: "test",
				},
			},
			want:    nil,
			wantErr: errors.New("repository failure"),
			mockFn: func(aa args, ff fields) {
				ff.Repository.Mock.
					On("Create", mock.Anything, &v1team.Team{Name: aa.request.Name}).
					Return(errors.New("repository failure"))

				ff.Validator.On("Validate", aa.request).Return(nil)
			},
		},
		{
			name: "should return error if validation failed",
			fields: fields{
				Repository: new(teamMock.MockRepository),
				Validator:  new(teamMock.MockValidator),
			},
			args: args{
				ctx: context.TODO(),
				request: &v1team.CreateTeamRequest{
					Name: "",
				},
			},
			want:    nil,
			wantErr: status.Error(codes.InvalidArgument, "team name is required"),
			mockFn: func(aa args, ff fields) {
				ff.Validator.On("Validate", aa.request).Return(errors.New("validation failed"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				Repository: tt.fields.Repository,
				Validator:  tt.fields.Validator,
			}

			tt.mockFn(tt.args, tt.fields)
			response, err := s.Create(tt.args.ctx, tt.args.request)
			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.want, response)
			tt.fields.Repository.AssertExpectations(t)
			tt.fields.Validator.AssertExpectations(t)
		})
	}
}

func TestServer_GetTeam(t *testing.T) {
	type fields struct {
		Repository *teamMock.MockRepository
		Validator  *teamMock.MockValidator
	}
	type args struct {
		ctx     context.Context
		request *v1team.GetTeamRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *v1team.GetTeamResponse
		wantErr error
		mockFn  func(aa args, ff fields)
	}{
		{
			name: "should success get team",
			fields: fields{
				Repository: new(teamMock.MockRepository),
				Validator:  new(teamMock.MockValidator),
			},
			args: args{
				ctx: context.TODO(),
				request: &v1team.GetTeamRequest{
					Id: 1,
				},
			},
			want: &v1team.GetTeamResponse{
				Team: &v1team.Team{
					Id:   1,
					Name: "test",
				},
			},
			wantErr: nil,
			mockFn: func(aa args, ff fields) {
				ff.Repository.Mock.
					On("GetTeam", mock.Anything, aa.request.Id).
					Return(&v1team.Team{
						Id:   1,
						Name: "test",
					}, nil)

				ff.Validator.On("Validate", aa.request).Return(nil)
			},
		},
		{
			name: "should return empty response if repository fails",
			fields: fields{
				Repository: new(teamMock.MockRepository),
				Validator:  new(teamMock.MockValidator),
			},
			args: args{
				ctx: context.TODO(),
				request: &v1team.GetTeamRequest{
					Id: 1,
				},
			},
			want:    nil,
			wantErr: status.Error(codes.Internal, "error getting team"),
			mockFn: func(aa args, ff fields) {
				ff.Repository.Mock.
					On("GetTeam", mock.Anything, aa.request.Id).
					Return(nil, errors.New("repository failure"))

				ff.Validator.On("Validate", aa.request).Return(nil)
			},
		},
		{
			name: "should return not found response if repository return not found",
			fields: fields{
				Repository: new(teamMock.MockRepository),
				Validator:  new(teamMock.MockValidator),
			},
			args: args{
				ctx: context.TODO(),
				request: &v1team.GetTeamRequest{
					Id: 1,
				},
			},
			want:    nil,
			wantErr: status.Error(codes.NotFound, "team not found"),
			mockFn: func(aa args, ff fields) {
				ff.Repository.Mock.
					On("GetTeam", mock.Anything, aa.request.Id).
					Return(nil, customError.RecordNotFoundError)

				ff.Validator.On("Validate", aa.request).Return(nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				Repository: tt.fields.Repository,
				Validator:  tt.fields.Validator,
			}

			tt.mockFn(tt.args, tt.fields)
			response, err := s.GetTeam(tt.args.ctx, tt.args.request)
			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.want, response)
			tt.fields.Repository.AssertExpectations(t)
			tt.fields.Validator.AssertExpectations(t)
		})
	}
}
