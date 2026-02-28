package jenkins

import (
	"time"

	"github.com/shirbental/jenkins-envoy/internal/model"
)

type MockClient struct{}

func NewMockClient() *MockClient { return &MockClient{} }

func (m *MockClient) GetBuild(job string, number int) (*model.Build, error) {
	return &model.Build{
		JobName:    job,
		Number:     number,
		Status:     "FAILED",
		StartedAt:  time.Now().Add(-2 * time.Minute),
		DurationMs: 45000,
		URL:        "https://jenkins.example/job/" + job,
	}, nil
}
