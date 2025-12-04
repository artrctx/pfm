package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/artrctx/pfm/internal/firewall"
	"github.com/artrctx/pfm/internal/upnp"
	"github.com/spf13/cobra"
)

// https://github.com/spf13/cobra
var rootCmd = &cobra.Command{
	Use:   "pfm",
	Short: "Port forward designated port and make it accesible",
	Long: `Port forward given port so your friend can access your destination. 
	(ONLY SUPPORTS LINUX FOR NOW)
	pfm --port 25565 --protocol tcp --firewall iptables --stun stun:stun1.l.google.com:3478`,
	Run: portForwardMe,
}

var (
	// port to open
	port uint16
	// protocol tcp || udp
	protocol string
	// for stun server
	stunAddr string
	// for firewall config
	firewallProvider string
)

func portForwardMe(cmd *cobra.Command, args []string) {
	provider, err := firewall.GetProvider(firewallProvider)
	if err != nil {
		log.Panicln(err)
	}
	protocol, err := firewall.GetProtocol(protocol)
	if err != nil {
		log.Panicln(err)
	}

	log.Printf("Initializing firewall client for %v\n", provider)
	fw, err := firewall.New(provider)
	if err != nil {
		log.Panicln(err)
	}
	log.Printf("Adding firewall ruleset for protocol %v | port %v\n", protocol, port)
	ruleset, err := fw.AllowPort(port, protocol)
	if err != nil {
		log.Panicln(err)
	}
	defer ruleset.Close()

	fmt.Println("Initializing UPNP")
	upnpClient, err := upnp.NewClient()
	if err != nil {
		log.Panicln(err)
	}

	externIp, err := upnpClient.GetExternalIPAddress()
	if err != nil {
		log.Panicln(err)
	}

	fmt.Println(externIp)

	for {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := scanner.Text()
		fmt.Printf("input: %v\n", input)
		break
	}
	// TODO
	// CONFIG FIREWALL TO PROVIDED OPEN PORT (Defer reset firewall setting)
	// CONFIG UPNP TO add specified port forwarding setup (Defer reset)
	// IN FOR LOOP
	// - SHOULD GET CURRENTLY CONFIGURED DNS IP
	// - GET ACTUAL IP
	// -- IF MISMATCH UPDATE CONFIG
	// MAYBE PANIC NOTIFIER
}

func init() {
	rootCmd.Flags().Uint16VarP(&port, "port", "p", 0, "Port to forward request to")
	rootCmd.MarkFlagRequired("port")

	rootCmd.Flags().StringVar(&protocol, "protocol", "tcp", "Protocol to listen to (supports tcp or udp)")

	rootCmd.Flags().StringVarP(&firewallProvider, "firewall", "f", "iptables", "Firewall provider name (currently supports iptables only)")

	rootCmd.Flags().StringVarP(&stunAddr, "stun", "s", "stun:stun1.l.google.com:3478", "STUN address to use")
}

func main() {
	rootCmd.Execute()
}
