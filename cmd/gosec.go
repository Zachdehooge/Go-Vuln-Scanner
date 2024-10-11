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

// gosecCmd represents the gosec command
var scanCMD = &cobra.Command{
	Use:   "gosec",
	Short: "gosecs the current project directory using the gosec command to find application vulnerabilities locally",
	Long:  `gosecs the current project directory using the gosec command to find application vulnerabilities locally`,

	Run: func(cmd *cobra.Command, args []string) {

		style := lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FAFAFA")). // Text color
			Background(lipgloss.Color("#1D3557")). // Background color
			Padding(1).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#A8DADC"))

		scan := exec.Command("gosec", ".")
		output, err := scan.Output()
		if err != nil {
			log.Fatalf("Error executing command: %s", err)
		}

		cmdOutput := string(output)

		styledCommand := style.Render("Gosec Scan\n\n" + cmdOutput)

		// Print the styled output
		fmt.Println(styledCommand)

	},
}

func init() {
	rootCmd.AddCommand(scanCMD)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// gosecCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// gosecCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
