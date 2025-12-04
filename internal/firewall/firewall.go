package firewall

import (
	"fmt"
	"io"
)

type Provider string

const (
	IPTables Provider = "iptables"
	NFTables Provider = "nptables"
)

func GetProvider(prov string) (Provider, error) {
	switch prov {
	case string(IPTables):
		return IPTables, nil
	case string(NFTables):
		return NFTables, nil
	default:
		return "", fmt.Errorf("provider name: %v is not supported", prov)
	}
}

type Ruleset interface {
	io.Closer
}

type Firewall interface {
	AllowPort(port uint16) (Ruleset, error)
}

var fw Firewall

func New(provider Provider) (Firewall, error) {
	if fw != nil {
		return fw, nil
	}
	switch provider {
	case IPTables:
		return newIptables()
	case NFTables:
		return nil, fmt.Errorf("%v as firewall provider is currently not supported", NFTables)
	default:
		return nil, fmt.Errorf("specified provider is not suppoerted")
	}
}
