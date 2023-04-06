package lib

import (
	"github.com/chelnak/ysmrr"
	"github.com/go-git/go-git/v5"
)

type Repo struct {
	Name string
	Path string
}

func Status(repo Repo, r git.Repository, spinner *ysmrr.Spinner, status string) {
	ref, _ := r.Head()
	commit, _ := r.CommitObject(ref.Hash())

	current := spinner.GetMessage()

	var message = FmtMessage(
		repo.Name,
		ref.Name().Short(),
		commit.Hash.String()[0:8],
		status,
		commit.Message,
		commit.Author.Name,
		commit.Author.Email,
	)

	if current != message {
		spinner.UpdateMessage(message)
	}
}
