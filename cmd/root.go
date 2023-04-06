/*
Copyleft (x) 2023 Des <des@riseup.net>
*/

package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/chelnak/ysmrr"
	"github.com/spf13/cobra"
	"gitlab.com/desyncr/exads-go/lib"
	"gopkg.in/yaml.v2"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "exads-go",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}


func CmdMain(f func(lib.Repo, string, ysmrr.SpinnerManager), ccmd *cobra.Command, args []string) {
	yfile, err := ioutil.ReadFile("config.yaml")

	if err != nil {
		log.Fatal(err)
	}

	repos := make(map[string]lib.Repo)

	err2 := yaml.Unmarshal(yfile, &repos)

	if err2 != nil {
		log.Fatal(err2)
	}

	// Create a new spinner manager
	sm := ysmrr.NewSpinnerManager(
		ysmrr.WithFrameDuration(time.Millisecond * 200),
	)
	sm.Start()

	var branch = ""
	if len(args) > 0 {
		branch = args[0]
	}

	start := time.Now()

	fmt.Println("Starting...")

	for _, repo := range repos {
		wg.Add(1)
		go f(repo, branch, sm)
	}

	wg.Wait()

	sm.Stop()

	elapsed := time.Since(start)
	log.Printf("Took %s", elapsed)
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.exads-go.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
