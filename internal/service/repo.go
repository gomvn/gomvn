package service

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/gomvn/gomvn/internal/config"
	"github.com/gomvn/gomvn/internal/entity"
)

func NewRepoService(conf *config.App, storage *Storage, ps *PathService) *RepoService {
	return &RepoService{
		repository: conf.Repository,
		storage:    storage,
		ps:         ps,
	}
}

type RepoService struct {
	repository []string
	storage    *Storage
	ps         *PathService
}

func (s *RepoService) GetRepositories() map[string][]*entity.Artifact {
	result := map[string][]*entity.Artifact{}
	for _, repo := range s.repository {
		result[repo] = []*entity.Artifact{}
		repoPath := s.storage.File(repo)
		_ = filepath.Walk(repoPath, func(path string, info os.FileInfo, err error) error {
			if strings.HasSuffix(path, ".pom") {
				path = strings.Replace(path, "\\", "/", -1)
				path = strings.Replace(path, repoPath+"/", "", 1)
				artifact := entity.NewArtifact(path, info.ModTime())
				result[repo] = append(result[repo], artifact)
			}
			return nil
		})
	}
	return result
}
