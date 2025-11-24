package main

import (
	"github.com/spf13/cobra"
)

// https://github.com/spf13/cobra
var rootCmd = &cobra.Command{
	Use:   "pfm",
	Short: "Port forward designated port and make it accesible",
	Long: `Port forward given port so your friend can access your destination.
	pfm --help`,
}

func main() {
	rootCmd.Execute()
}
