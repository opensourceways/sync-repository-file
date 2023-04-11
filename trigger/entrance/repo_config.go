package main

import "fmt"

// repos
type repos []repoConfig

func (cfg repos) Validate() error {
	for i := range cfg {
		if err := cfg[i].validate(); err != nil {
			return err
		}
	}

	return nil
}

// repoConfig
type repoConfig struct {
	// Platform is the code platform.
	Platform string `json:"platform"        required:"true"`

	// FileNames is the list of files to be synchronized.
	FileNames []string `json:"file_names"   required:"true"`

	OrgRepos []orgRepos `json:"org_repos,omitempty"`
}

func (s *repoConfig) validate() error {
	if s.Platform == "" {
		return fmt.Errorf("must set platform")
	}

	if len(s.FileNames) == 0 {
		return fmt.Errorf("must set file_names")
	}

	for _, item := range s.OrgRepos {
		if err := item.validate(); err != nil {
			return err
		}
	}

	return nil
}

type orgRepos struct {
	Org           string   `json:"org" required:"true"`
	Repos         []string `json:"repos,omitempty"`
	ExcludedRepos []string `json:"excluded_repos,omitempty"`
}

func (o *orgRepos) validate() error {
	if o.Org == "" {
		return fmt.Errorf("must set org")
	}

	if len(o.Repos) > 0 && len(o.ExcludedRepos) > 0 {
		return fmt.Errorf("can't set repos and excluded_repos for org:%s at same time", o.Org)
	}

	return nil
}
