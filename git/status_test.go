package git

import (
	"reflect"
	"testing"
)

func TestStatus(t *testing.T) {
	var tests = []struct {
		makeTestRepo func(*testing.T) *Repo
		want         *RepoStatus
	}{
		{newRepoEmpty, &RepoStatus{
			HasUntrackedFiles:     false,
			HasUncommittedChanges: false,
			CurrentBranch:         StatusUnknown,
			Branches:              nil,
		}},
		{newRepoWithUntracked, &RepoStatus{
			HasUntrackedFiles:     true,
			HasUncommittedChanges: false,
			CurrentBranch:         StatusUnknown,
			Branches:              nil,
		}},
		{newRepoWithStaged, &RepoStatus{
			HasUntrackedFiles:     false,
			HasUncommittedChanges: true,
			CurrentBranch:         StatusUnknown,
			Branches:              nil,
		}},
		{newRepoWithCommit, &RepoStatus{
			HasUntrackedFiles:     false,
			HasUncommittedChanges: false,
			CurrentBranch:         "master",
			Branches: []*BranchStatus{
				{
					Name:     "master",
					Upstream: "",
					Behind:   0,
					Ahead:    0,
				},
			},
		}},
		{newRepoWithModified, &RepoStatus{
			HasUntrackedFiles:     false,
			HasUncommittedChanges: true,
			CurrentBranch:         "master",
			Branches: []*BranchStatus{
				{
					Name:     "master",
					Upstream: "",
					Behind:   0,
					Ahead:    0,
				},
			},
		}},
		{newRepoWithIgnored, &RepoStatus{
			HasUntrackedFiles:     false,
			HasUncommittedChanges: false,
			CurrentBranch:         "master",
			Branches: []*BranchStatus{
				{
					Name:     "master",
					Upstream: "",
					Behind:   0,
					Ahead:    0,
				},
			},
		}},
		{newRepoWithLocalBranch, &RepoStatus{
			HasUntrackedFiles:     false,
			HasUncommittedChanges: false,
			CurrentBranch:         "master",
			Branches: []*BranchStatus{
				{
					Name:     "master",
					Upstream: "",
					Behind:   0,
					Ahead:    0,
				}, {
					Name:     "local",
					Upstream: "",
					Behind:   0,
					Ahead:    0,
				},
			},
		}},
		{newRepoWithClonedBranch, &RepoStatus{
			HasUntrackedFiles:     false,
			HasUncommittedChanges: false,
			CurrentBranch:         "local",
			Branches: []*BranchStatus{
				{
					Name:     "master",
					Upstream: "origin/master",
					Behind:   0,
					Ahead:    0,
				}, {
					Name:     "local",
					Upstream: "",
					Behind:   0,
					Ahead:    0,
				},
			},
		}},
		{newRepoWithDetachedHead, &RepoStatus{
			HasUntrackedFiles:     false,
			HasUncommittedChanges: false,
			CurrentBranch:         StatusDetached,
			Branches: []*BranchStatus{
				{
					Name:     "master",
					Upstream: "",
					Behind:   0,
					Ahead:    0,
				},
			},
		}},
		{newRepoWithBranchAhead, &RepoStatus{
			HasUntrackedFiles:     false,
			HasUncommittedChanges: false,
			CurrentBranch:         "master",
			Branches: []*BranchStatus{
				{
					Name:     "master",
					Upstream: "origin/master",
					Behind:   0,
					Ahead:    1,
				},
			},
		}},
		{newRepoWithBranchBehind, &RepoStatus{
			HasUntrackedFiles:     false,
			HasUncommittedChanges: false,
			CurrentBranch:         "master",
			Branches: []*BranchStatus{
				{
					Name:     "master",
					Upstream: "origin/master",
					Behind:   1,
					Ahead:    0,
				},
			},
		}},
		{newRepoWithBranchAheadAndBehind, &RepoStatus{
			HasUntrackedFiles:     false,
			HasUncommittedChanges: false,
			CurrentBranch:         "master",
			Branches: []*BranchStatus{
				{
					Name:     "master",
					Upstream: "origin/master",
					Behind:   3,
					Ahead:    2,
				},
			},
		}},
	}

	for _, test := range tests {
		repo := test.makeTestRepo(t)

		err := repo.LoadStatus()
		checkFatal(t, err)

		if !reflect.DeepEqual(repo.Status, test.want) {
			t.Errorf("Wrong repo status, got: %+v; want: %+v", repo.Status, test.want)
		}
	}
}

// TODO: test branch status when tracking a local branch
// TODO: test head pointing to a tag
// TODO: newRepoWithGlobalGitignore
// TODO: newRepoWithGlobalGitignoreSymlink
