package main

import (
	"log"

	"github.com/arthurDiff/pfm/internal/stun"
	"github.com/spf13/cobra"
)

// https://github.com/spf13/cobra
var rootCmd = &cobra.Command{
	Use:   "pfm",
	Short: "Port forward designated port and make it accesible",
	Long: `Port forward given port so your friend can access your destination.
	pfm --stun stun:stun1.l.google.com:3478`,
	Run: portForwardMe,
}

func portForwardMe(cmd *cobra.Command, args []string) {
	flags := cmd.Flags()
	stunAddr, err := flags.GetString("stun")
	if err != nil {
		log.Fatalf("Failed to get stun flag value error %v", err)
	}
	stunClient := stun.NewClient(stunAddr)
	defer stunClient.Close()

	// TODO
	// CONFIG FIREWALL TO PROVIDED OPEN PORT
	// IN FOR LOOP
	// - SHOULD GET CURRENTLY CONFIGURED DNS IP
	// - GET ACTUAL IP
	// -- IF MISMATCH UPDATE CONFIG
	// MAYBE PANIC NOTIFIER
}

func init() {
	rootCmd.Flags().StringP("stun", "s", "stun:stun1.l.google.com:3478", "STUN address to use. (Default: stun:stun1.l.google.com:3478)")
}

func main() {
	rootCmd.Execute()
}
