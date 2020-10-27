package kubelet

import (
	"context"

	kubeletTypes "github.com/mackerelio/mackerel-container-agent/internal/k8s-apis/stats/v1alpha1"
	kubernetesTypes "k8s.io/api/core/v1"
)

// MockClient represents a mock client of Kubelet APIs
type MockClient struct {
	getPodCallback      func(context.Context) (*kubernetesTypes.Pod, error)
	getPodStatsCallback func(context.Context) (*kubeletTypes.PodStats, error)
}

// MockClientOption represents an option of mock client of Kubelet APIs
type MockClientOption func(*MockClient)

// NewMockClient creates a new mock client of Kubelet APIs
func NewMockClient(opts ...MockClientOption) *MockClient {
	c := &MockClient{}
	for _, o := range opts {
		c.ApplyOption(o)
	}
	return c
}

// ApplyOption applies a mock client option
func (c *MockClient) ApplyOption(opt MockClientOption) {
	opt(c)
}

type errCallbackNotFound string

func (err errCallbackNotFound) Error() string {
	return string(err) + " callback not found"
}

// GetPod ...
func (c *MockClient) GetPod(ctx context.Context) (*kubernetesTypes.Pod, error) {
	if c.getPodCallback != nil {
		return c.getPodCallback(ctx)
	}
	return nil, errCallbackNotFound("GetPod")
}

// MockGetPod returns an option to set the callback of GetPod
func MockGetPod(callback func(context.Context) (*kubernetesTypes.Pod, error)) MockClientOption {
	return func(c *MockClient) {
		c.getPodCallback = callback
	}
}

// GetPodStats ...
func (c *MockClient) GetPodStats(ctx context.Context) (*kubeletTypes.PodStats, error) {
	if c.getPodStatsCallback != nil {
		return c.getPodStatsCallback(ctx)
	}
	return nil, errCallbackNotFound("GetPodStats")
}

// MockGetPodStats returns an option to set the callback of GetPodStats
func MockGetPodStats(callback func(context.Context) (*kubeletTypes.PodStats, error)) MockClientOption {
	return func(c *MockClient) {
		c.getPodStatsCallback = callback
	}
}
