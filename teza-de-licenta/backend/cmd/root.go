package cmd

import "github.com/spf13/cobra"

func Exec() error {
	var rootCmd = &cobra.Command{Use: "teza", Short: "`teza` este un program care indeplineste cerintele tezei de licenta."}

	rootCmd.AddCommand(ServeCMD())
	rootCmd.AddCommand(WorkerCMD())

	return rootCmd.Execute()
}
