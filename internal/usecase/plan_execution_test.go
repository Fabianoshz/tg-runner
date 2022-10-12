package usecase

import (
	"testing"

	"github.com/fabianoshz/tg-runner/internal/repository/mocks"
	"github.com/go-playground/assert"
	"github.com/google/uuid"
)

func TestPlanExecution(t *testing.T) {
	mockPersistenceRepository := new(mocks.Persistence)

	executionPlanner := NewExecutionPlannerService(mockPersistenceRepository)

	type args struct {
		changelist string
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
				changelist: "testdata/changelist-plan.yaml",
			},
			want:         true,
			expectdError: nil,
		},
		{
			name: "Test basic destroy plan",
			args: args{
				changelist: "testdata/changelist-destroy.yaml",
			},
			want:         true,
			expectdError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// given
			// TODO read this directly from the yaml file
			mockPersistenceRepository.On("SavePlanfile", uuid.New().String(), "planfile", uuid.New(), "internal/usecase/testdata/terragrunt/basic-terragrunt").Return(true)
			mockPersistenceRepository.On("SavePlanfile", uuid.New().String(), "planfile", uuid.New(), "internal/usecase/testdata/terragrunt/basic-terragrunt-2").Return(true)

			// when
			got := executionPlanner.PlanExecution(tt.args.changelist)

			// then
			assert.Equal(t, tt.want, got)
		})
	}

}
