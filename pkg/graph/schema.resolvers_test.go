package graph

import (
	"errors"
	"testing"

	"github.com/b-open/jobbuzz/pkg/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockService struct {
	mock.Mock
}

func (s *MockService) GetJobs() ([]*model.Job, error) {
	args := s.Called()

	jobs := args.Get(0)
	if jobs == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).([]*model.Job), args.Error(1)
}

func TestJobs(t *testing.T) {
	t.Run("test return 1 job", func(t *testing.T) {
		mockService := MockService{}
		mockService.On("GetJobs").Return([]*model.Job{
			{
				BaseModel: model.BaseModel{
					ID: 1,
				},
				Title: "test job",
			},
		}, nil)

		r := Resolver{Service: &mockService}

		result, err := r.Query().Jobs(nil)
		if err != nil {
			t.Fatal(err)
		}

		assert.NotEmpty(t, result, "Jobs should not be empty")
		assert.Len(t, result, 1)
	})

	t.Run("test return 20 jobs", func(t *testing.T) {
		mockService := MockService{}
		var mockJobs []*model.Job
		for i := 0; i < 20; i++ {
			mockJobs = append(mockJobs, &model.Job{
				BaseModel: model.BaseModel{
					ID: uint(i),
				},
				Title: "test job",
			})
		}
		mockService.On("GetJobs").Return(mockJobs, nil)

		r := Resolver{Service: &mockService}

		result, err := r.Query().Jobs(nil)
		if err != nil {
			t.Fatal(err)
		}

		assert.NotEmpty(t, result, "Jobs should not be empty")
		assert.Len(t, result, 20, "Jobs length is not correct")
	})

	t.Run("test return no jobs", func(t *testing.T) {
		mockService := MockService{}
		mockService.On("GetJobs").Return([]*model.Job{}, nil)

		r := Resolver{Service: &mockService}

		result, err := r.Query().Jobs(nil)
		if err != nil {
			t.Fatal(err)
		}

		assert.Empty(t, result, "Jobs should be empty")
	})

	t.Run("test error", func(t *testing.T) {
		mockService := MockService{}
		mockService.On("GetJobs").Return(nil, errors.New("error"))

		r := Resolver{Service: &mockService}

		_, err := r.Query().Jobs(nil)
		assert.NotNil(t, err, "Error was expected but not found.")
	})
}
