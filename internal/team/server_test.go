package team

import (
	"context"
	"errors"
	"testing"

	v1team "github.com/ecojuntak/laklak/gen/go/v1/team"
	teamMock "github.com/ecojuntak/laklak/mocks/team"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestServer_Create(t *testing.T) {
	type fields struct {
		Repository *teamMock.MockRepository
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
			},
			args: args{
				ctx: context.TODO(),
				request: &v1team.CreateTeamRequest{
					Name: "test",
				},
			},
			want:    &v1team.CreateTeamResponse{},
			wantErr: nil,
			mockFn: func(aa args, ff fields) {
				ff.Repository.Mock.
					On("Create", mock.Anything, &v1team.Team{Name: aa.request.Name}).
					Return(nil)
			},
		},
		{
			name: "should return empty response if repository fails",
			fields: fields{
				Repository: new(teamMock.MockRepository),
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
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				Repository: tt.fields.Repository,
			}

			tt.mockFn(tt.args, tt.fields)
			response, err := s.Create(tt.args.ctx, tt.args.request)
			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, tt.want, response)
			tt.fields.Repository.AssertExpectations(t)
		})
	}
}
