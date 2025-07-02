package cmd

import (
	"testing"

	"github.com/jackchuka/gh-oss-watch/services"
	mock_services "github.com/jackchuka/gh-oss-watch/services/mock"
	"go.uber.org/mock/gomock"
)

func TestHandleConfigAdd_Success(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockConfig := mock_services.NewMockConfigService(ctrl)
	mockCache := mock_services.NewMockCacheService(ctrl)
	mockGitHub := mock_services.NewMockGitHubService(ctrl)
	mockOutput := mock_services.NewMockOutput(ctrl)

	cli := NewCLI(mockConfig, mockCache, mockGitHub, mockOutput)

	config := &services.Config{Repos: []services.RepoConfig{}}

	// Set up expectations
	mockConfig.EXPECT().Load().Return(config, nil)
	mockConfig.EXPECT().Save(gomock.Any()).DoAndReturn(func(c *services.Config) error {
		// Verify the repo was added
		if len(c.Repos) != 1 {
			t.Errorf("Expected 1 repo, got %d", len(c.Repos))
		}
		if c.Repos[0].Repo != "owner/repo" {
			t.Errorf("Expected 'owner/repo', got %s", c.Repos[0].Repo)
		}
		return nil
	})
	mockOutput.EXPECT().Printf(gomock.Any(), gomock.Any()).AnyTimes()

	err := cli.handleConfigAdd("owner/repo", []string{"stars", "issues"})

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestHandleConfigAdd_InvalidEvents(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockConfig := mock_services.NewMockConfigService(ctrl)
	mockCache := mock_services.NewMockCacheService(ctrl)
	mockGitHub := mock_services.NewMockGitHubService(ctrl)
	mockOutput := mock_services.NewMockOutput(ctrl)

	cli := NewCLI(mockConfig, mockCache, mockGitHub, mockOutput)

	config := &services.Config{Repos: []services.RepoConfig{}}
	mockConfig.EXPECT().Load().Return(config, nil)

	err := cli.handleConfigAdd("owner/repo", []string{"invalid_event"})

	if err == nil {
		t.Error("Expected error for invalid events, got nil")
	}
}
