// main.go
package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "demo",
		Short: "A demo CLI tool",
		Run: func(cmd *cobra.Command, args []string) {
			host, _ := cmd.Flags().GetString("host")
			fmt.Printf("Connecting to %s...\n", host)
			// 在这里可以编写与主机交互的业务逻辑
		},
	}

	rootCmd.Flags().StringP("host", "H", "", "Specify the host")
	rootCmd.MarkFlagRequired("host")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}