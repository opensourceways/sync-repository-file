package app

import "github.com/opensourceways/sync-repository-file/server/domain"

type CmdToFetchRepoBranch = domain.OrgRepo

type CmdToFetchRepoFile struct {
	domain.OrgRepo

	Branch string
}

type CmdToFetchFileContent struct {
	domain.OrgRepo

	Branch   string
	FilePath string
}
