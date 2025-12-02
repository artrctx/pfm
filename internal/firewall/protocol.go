package firewall

type Protocol string

const (
	TCP  Protocol = "tcp"
	UDP  Protocol = "udp"
	BOTH Protocol = "tcp/udp"
)
