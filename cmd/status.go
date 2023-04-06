/*
Copyleft (x) 2023 Des <des@riseup.net>
*/
package cmd

import (
	"github.com/chelnak/ysmrr"
	"github.com/spf13/cobra"
	"gitlab.com/desyncr/exads-go/lib"

	"github.com/go-git/go-git/v5"
)

// checkoutCmd represents the checkout command
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Display current branch and last update",
	Run: func (ccmd *cobra.Command, args []string) {
		CmdMain(getStatus, ccmd, args)
	},
}

func getStatus(repo lib.Repo, branch string, sm ysmrr.SpinnerManager) {
	defer wg.Done()

	// STATUS
	spinner := sm.AddSpinner(
		lib.FmtMessage(
			repo.Name,
			"Initializing",
			"",
			"",
			"",
			"",
			"",
		),
	)

	r, _ := git.PlainOpen(repo.Path)

	// STATUS
	lib.Status(repo, *r, spinner, "Complete")

	spinner.Complete()
}

func init() {
	rootCmd.AddCommand(statusCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// statusCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// statusCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
