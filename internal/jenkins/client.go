package jenkins

import "github.com/shirbental/jenkins-envoy/internal/model"

type Client interface {
	GetBuild(job string, number int) (*model.Build, error)
}
