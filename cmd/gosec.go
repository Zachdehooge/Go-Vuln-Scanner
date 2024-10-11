/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/spf13/cobra"
)

// gosecCmd represents the gosec command
var gosecCmd = &cobra.Command{
	Use:   "gosec",
	Short: "gosecs the current project directory using the gosec command to find application vulnerabilities locally",
	Long:  `gosecs the current project directory using the gosec command to find application vulnerabilities locally`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("\nRUNNING GOSEC SCAN...")

		gosec := exec.Command("gosec", ".")
		output, err := gosec.Output()
		if err != nil {
			log.Fatalf("Error executing command: %s", err)
		}

		fmt.Println(string(output))
	},
}

func init() {
	rootCmd.AddCommand(gosecCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// gosecCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// gosecCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
