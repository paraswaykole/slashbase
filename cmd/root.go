package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "slashbase",
	Short: "slashbase is a cli & an api server for quering databases",
	Long:  `slashbase is a cli & an api server for quering databases`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("slashbase > ")
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			text := scanner.Text()
			handleCmd(text)
			if cliApp.CurrentDB == nil {
				fmt.Print("slashbase > ")
			} else {
				fmt.Printf("%s > ", cliApp.CurrentDB.Name)
			}
		}

		if scanner.Err() != nil {
			fmt.Println("unexpected error occurred")
			os.Exit(1)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
