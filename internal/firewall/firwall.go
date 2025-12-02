package firewall

import (
	"log"
)

type Provider string

const (
	IPTables Provider = "iptables"
	NFTables Provider = "nptables"
)

type Firewall interface {
	AllowPort(port uint16) error
}

var fw Firewall

func New(provider Provider) Firewall {
	if fw != nil {
		return fw
	}
	switch provider {
	case IPTables:
		fw = newIptables()
	case NFTables:
		log.Fatalf("%v as firewall provider is currently not supported", NFTables)
	default:
		log.Fatalf("specified provider is not suppoerted")
	}
	return fw
}
