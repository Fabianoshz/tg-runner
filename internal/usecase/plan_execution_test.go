package usecase

import (
	"testing"

	"github.com/fabianoshz/iflantis/internal/entity"
	"github.com/fabianoshz/iflantis/internal/repository/mocks"
	"github.com/go-playground/assert"
	"github.com/google/uuid"
)

func TestPlanExecution(t *testing.T) {
	mockProjectRepository := new(mocks.Persistence)

	executionPlanner := NewExecutionPlannerService(mockProjectRepository)

	type args struct {
		changelist entity.Changelist
	}

	tests := []struct {
		name         string
		args         args
		want         bool
		expectdError error
	}{
		{
			name: "Test basic plan",
			args: args{
				changelist: entity.Changelist{
					Id: uuid.New(),
					Resources: []entity.Resource{
						{
							ID:     "abc",
							Path:   "internal/usecase/testdata/terragrunt/basic-terragrunt",
							Action: entity.Plan,
						},
						{
							ID:     "123",
							Path:   "internal/usecase/testdata/terragrunt/basic-terragrunt-2",
							Action: entity.Plan,
						},
					},
				},
			},
			want:         true,
			expectdError: nil,
		},
		{
			name: "Test basic destroy plan",
			args: args{
				changelist: entity.Changelist{
					Id: uuid.New(),
					Resources: []entity.Resource{
						{
							ID:     "abc",
							Path:   "internal/usecase/testdata/terragrunt/basic-terragrunt",
							Action: entity.PlanDestroy,
						},
						{
							ID:     "123",
							Path:   "internal/usecase/testdata/terragrunt/basic-terragrunt-2",
							Action: entity.PlanDestroy,
						},
					},
				},
			},
			want:         true,
			expectdError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// given
			for _, r := range tt.args.changelist.Resources {
				mockProjectRepository.On("SavePlanfile", r.ID, "planfile", tt.args.changelist.Id, r.Path).Return(true)
			}

			// when
			got := executionPlanner.PlanExecution(tt.args.changelist)

			// then
			assert.Equal(t, tt.want, got)
		})
	}

}
