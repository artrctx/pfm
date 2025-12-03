package firewall

import (
	"fmt"
	"io"
	"log"

	ipts "github.com/coreos/go-iptables/iptables"
	"golang.org/x/sync/errgroup"
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

type rule struct {
	chain      string
	rulespecs  []string
	preexisted bool
}

func (r *rule) append(ipt *iptables) error {
	exists, err := ipt.binding.Exists("filter", r.chain, r.rulespecs...)
	if err != nil {
		return err
	}
	if exists {
		r.preexisted = true
		return nil
	}
	return ipt.binding.Append("filter", r.chain, r.rulespecs...)
}

func (r *rule) delete(ipt *iptables) error {
	if r.preexisted {
		return nil
	}
	return ipt.binding.DeleteIfExists("filter", r.chain, r.rulespecs...)
}

type rules []rule

func (rs rules) deleteAll(ipt *iptables) error {
	g := errgroup.Group{}
	for _, r := range rs {
		g.Go(func() error { return r.delete(ipt) })
	}
	return g.Wait()
}

type ruleset struct {
	ipt   *iptables
	rules rules
}

func (rs ruleset) appendAll() error {
	applied := make(rules, 0, len(rs.rules))
	for _, r := range rs.rules {
		err := r.append(rs.ipt)
		if err != nil {
			cleanErr := applied.deleteAll(rs.ipt)
			return fmt.Errorf("append error: %w; cleanup error: %w", err, cleanErr)
		}
	}
	return nil
}
func (rs ruleset) Close() error {
	return rs.rules.deleteAll(rs.ipt)
}

func (ipt *iptables) AllowPort(port uint16) (io.Closer, error) {
	ruleset := ruleset{ipt, []rule{
		{chain: "INPUT", rulespecs: []string{"-p tcp", fmt.Sprintf("--dport %v", port), "-j ACCEPT"}},
		{chain: "INPUT", rulespecs: []string{"-p udp", fmt.Sprintf("--dport %v", port), "-j ACCEPT"}},
		{chain: "OUTPUT", rulespecs: []string{"-p tcp", fmt.Sprintf("--dport %v", port), "-j ACCEPT"}},
		{chain: "OUTPUT", rulespecs: []string{"-p tcp", fmt.Sprintf("--dport %v", port), "-j ACCEPT"}},
	}}

	if err := ruleset.appendAll(); err != nil {
		return nil, err
	}

	return ruleset, nil
}
