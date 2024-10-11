/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

// snykCmd represents the snyk command
var snykCmd = &cobra.Command{
	Use:   "snyk",
	Short: "Runs a snyk vulnerability scan on the project",
	Long:  `Runs a snyk vulnerability scan on the project`,
	Run: func(cmd *cobra.Command, args []string) {

		style := lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FAFAFA")). // Text color
			Background(lipgloss.Color("#1D3557")). // Background color
			Padding(1).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#A8DADC"))

		snyk := exec.Command("snyk", "test")
		output, err := snyk.Output()
		if err != nil {
			log.Fatalf("Error executing command: %s", err)
		}

		cmdOutput := string(output)

		styledCommand := style.Render("Snyk Scan\n" + cmdOutput)

		// Print the styled output
		fmt.Println(styledCommand)
	},
}

func init() {
	rootCmd.AddCommand(snykCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// snykCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// snykCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
