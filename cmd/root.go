package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var (
	rootCmd = &cobra.Command{
		Use: "server",
		Run: StartServer,
	}
	adminCmd = &cobra.Command{
		Use: "admin",
		Run: StartAdminServer,
	}
)

func init() {
	rootCmd.AddCommand(adminCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
