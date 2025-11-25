package stun

import "net/http"

type STUNClient http.Client

var STUN *STUNClient

// func NewSTUNClient() *STUNClient {
// 	if STUN != nil {
// 		return STUN
// 	}
// 	// todo
// }
