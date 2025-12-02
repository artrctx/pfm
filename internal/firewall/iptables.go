package firewall

import (
	"fmt"
	"log"

	"github.com/coreos/go-iptables/iptables"
)

// https://www.ipserverone.info/knowledge-base/how-to-open-ports-in-iptables/
type IPTables struct{ client *iptables.IPTables }

var iptclient *IPTables

func newIptables() *IPTables {
	if iptclient != nil {
		return iptclient
	}
	// defualt ipv4 protocol
	client, err := iptables.New()
	if err != nil {
		log.Fatalf("Failed to initialize IPTables with error %v", err)
	}
	iptclient = &IPTables{client}
	return iptclient
}

func (ipt *IPTables) allowPort(port uint16) error {
	list, err := ipt.client.List("filter", "INPUT")
	if err != nil {
		return err
	}
	fmt.Println(list, "list here")
	return nil
}
