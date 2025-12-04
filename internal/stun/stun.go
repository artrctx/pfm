package stun

import (
	"fmt"

	stn "github.com/pion/stun"
)

// https://dev.to/alakkadshaw/google-stun-server-list-21n4
// https://datatracker.ietf.org/doc/html/rfc5389
type Client struct{ client *stn.Client }

func NewClient(stunAddr string) (*Client, error) {
	addr, err := stn.ParseURI(stunAddr)
	if err != nil {
		return nil, fmt.Errorf("tcp addr resolve error: %v", err)
	}

	c, err := stn.DialURI(addr, &stn.DialConfig{})
	if err != nil {
		return nil, fmt.Errorf("tcp conn error: %v", err)
	}

	return &Client{c}, nil
}

func (c *Client) GetIP() (stn.XORMappedAddress, error) {
	var err error
	var xorAddr stn.XORMappedAddress
	if err = c.client.Do(stn.MustBuild(stn.TransactionID, stn.BindingRequest), func(res stn.Event) {
		//do stuff
		if res.Error != nil {
			err = res.Error
			return
		}

		if err = xorAddr.GetFrom(res.Message); err != nil {
			return
		}
	}); err != nil {
		return xorAddr, err
	}

	return xorAddr, nil
}

func (c *Client) Close() error {
	return c.client.Close()
}
