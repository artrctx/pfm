package firewall

import (
	"fmt"
	"io"
)

type Protocol string

const (
	TCP Protocol = "tcp"
	UDP Protocol = "udp"
)

func GetProtocol(p string) (Protocol, error) {
	switch p {
	case string(TCP):
		return TCP, nil
	case string(UDP):
		return UDP, nil
	default:
		return "", fmt.Errorf("protocol name: %v is not valid", p)
	}
}

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
	AllowPort(port uint16, protocol Protocol) (Ruleset, error)
}

func New(provider Provider) (Firewall, error) {
	switch provider {
	case IPTables:
		return newIptables()
	case NFTables:
		return nil, fmt.Errorf("%v as firewall provider is currently not supported", NFTables)
	default:
		return nil, fmt.Errorf("specified provider is not suppoerted")
	}
}
