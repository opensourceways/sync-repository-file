package app

import (
	"github.com/opensourceways/sync-repository-file/server/domain"
	"github.com/opensourceways/sync-repository-file/server/domain/codeplatform"
	"github.com/opensourceways/sync-repository-file/server/domain/message"
	"github.com/opensourceways/sync-repository-file/server/domain/repository"
)

type RepoFileService interface {
	FetchRepoBranch(codeplatform.CodePlatform, *CmdToFetchRepoBranch) error
	FetchRepoFile(codeplatform.CodePlatform, *CmdToFetchRepoFile) error
	FetchFileContent(codeplatform.CodePlatform, *CmdToFetchFileContent) error
}

func NewRepoFileService(
	repo repository.RepoFile,
	message message.RepoFile,
) repoFileService {
	return repoFileService{
		repo:    repo,
		message: message,
	}
}

type repoFileService struct {
	repo    repository.RepoFile
	message message.RepoFile
}

func (s repoFileService) FetchRepoBranch(
	p codeplatform.CodePlatform,
	cmd *CmdToFetchRepoBranch,
) error {
	v, err := p.ListBranches(*cmd)
	if err != nil {
		return err
	}

	task := domain.RepoBranchFetchedEvent{
		Platform: p.Platform(),
		Org:      cmd.Org,
		Repo:     cmd.Repo,
	}

	for i := range v {
		task.Branch = v[i].Name

		if err := s.message.SendRepoBranchFetchedEvent(&task); err != nil {
			return err
		}
	}

	return nil
}

func (s repoFileService) FetchRepoFile(
	p codeplatform.CodePlatform,
	cmd *CmdToFetchRepoFile,
) error {
	v, err := p.ListFiles(cmd.OrgRepo, cmd.Branch)
	if err != nil {
		return err
	}

	task := domain.RepoFileFetchedEvent{
		Platform: p.Platform(),
		Org:      cmd.Org,
		Repo:     cmd.Repo,
		Branch:   cmd.Branch,
	}

	for i := range v {
		task.FilePath = v[i].Path

		if err := s.message.SendRepoFileFetchedEvent(&task); err != nil {
			return err
		}
	}

	return nil
}

func (s repoFileService) FetchFileContent(
	p codeplatform.CodePlatform,
	cmd *CmdToFetchFileContent,
) error {
	v, err := p.GetFile(cmd.OrgRepo, cmd.Branch, cmd.FilePath)
	if err != nil {
		return err
	}

	return s.repo.SaveFile(
		domain.PlatformOrgRepo{
			Platform: p.Platform(),
			OrgRepo:  cmd.OrgRepo,
		},
		cmd.Branch,
		v,
	)
}
