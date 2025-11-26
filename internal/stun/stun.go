package stun

import (
	"log"

	stn "github.com/pion/stun"
)

// https://dev.to/alakkadshaw/google-stun-server-list-21n4
// https://datatracker.ietf.org/doc/html/rfc5389
type Client struct{ client *stn.Client }

var client *Client

func NewClient(stunAddr string) *Client {
	if client != nil {
		return client
	}

	// stunAddr := "stun:stun1.l.google.com:3478"
	addr, err := stn.ParseURI(stunAddr)
	if err != nil {
		log.Panicf("tcp addr resolve error: %v", err)
	}

	c, err := stn.DialURI(addr, &stn.DialConfig{})
	if err != nil {
		log.Panicf("tcp conn error: %v", err)
	}

	client = &Client{c}
	return client
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
