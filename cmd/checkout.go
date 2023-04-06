/*
Copyleft (x) 2023 Des <des@riseup.net>
*/
package cmd

import (
	"sync"

	"github.com/chelnak/ysmrr"
	"github.com/spf13/cobra"
	"gitlab.com/desyncr/exads-go/lib"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

// checkoutCmd represents the checkout command
var checkoutCmd = &cobra.Command{
	Use:   "checkout <branch>",
	Short: "Checkout and pull the branch given as argument",
	Run: func (ccmd *cobra.Command, args []string) {
		CmdMain(checkout, ccmd, args)
	},
}

var wg sync.WaitGroup
var mu sync.Mutex

func checkout(repo lib.Repo, branch string, sm ysmrr.SpinnerManager) {
	defer wg.Done()

	// STATUS
	spinner := sm.AddSpinner(
		lib.FmtMessage(
			repo.Name,
			"Initializing",
			branch,
			"",
			"",
			"",
			"",
		),
	)

	r, _ := git.PlainOpen(repo.Path)
	w, _ := r.Worktree()

	ref, _ := r.Head()
	if ref.Name().Short() != branch {
		lib.Status(repo, *r, spinner, "Checking out")
		// ACTION: checkout
		w.Checkout(&git.CheckoutOptions{
			Branch: plumbing.NewBranchReferenceName(branch),
		})
	}

	// STATUS
	lib.Status(repo, *r, spinner, "Pulling")

	// ACTION: pull
	w.Pull(&git.PullOptions{RemoteName: "origin"})

	// STATUS
	lib.Status(repo, *r, spinner, "Complete")
	spinner.Complete()
}


func init() {
	rootCmd.AddCommand(checkoutCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// checkoutCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// checkoutCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
