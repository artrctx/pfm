package firewall

import (
	"fmt"
	"log"

	ipts "github.com/coreos/go-iptables/iptables"
)

// https://www.ipserverone.info/knowledge-base/how-to-open-ports-in-iptables/
type iptables struct{ binding *ipts.IPTables }

func newIptables() *iptables {
	// defualt ipv4 protocol
	binding, err := ipts.New()
	if err != nil {
		log.Fatalf("Failed to initialize IPTables with error %v", err)
	}
	return &iptables{binding}
}

func (ipt *iptables) AllowPort(port uint16) error {
	// append input rule
	err := ipt.binding.AppendUnique("filter", "INPUT", "-p all", fmt.Sprintf("--dport %v", port), "-j ACCEPT")
	if err != nil {
		return err
	}

	err = ipt.binding.AppendUnique("filter", "OUTPUT", "-p all", fmt.Sprintf("--dport %v", port), "-j ACCEPT")
	if err != nil {
		return err
	}

	return nil
}
