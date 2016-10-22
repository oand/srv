// Copyright (c) 2016 Øystein Andersen
// Use of this source code is governed by MIT license found at
// https://github.com/oand/srv/blob/master/LICENSE

package srv

import "net"

// Dial resolves an SRV query of the given address on the named network.
// Then it will attempt to connect to the service in order priority and randomized by weight within a priority.
//
// Values for network are "tcp" or "udp"
// address is in the form: domain:service
//
// Examples:
//	Dial("tcp", "example.com:xmpp-client")
//	Dial("udp", "example.com:stun")
func Dial(network, address string) (net.Conn, error) {
	addr, err := Lookup(network, address)
	if err != nil {
		return nil, err
	}
	return dial(addr)
}

// DialSRV resolves an SRV query of the given service, protocol, and domain name.
// Then it will attempt to connect to the service in order priority and randomized by weight within a priority.
func DialSRV(service, protocol, domain string) (net.Conn, error) {
	addr, err := LookupSRV(service, protocol, domain)
	if err != nil {
		return nil, err
	}
	return dial(addr)
}

func dial(addr []net.Addr) (conn net.Conn, err error) {
	for _, v := range addr {
		conn, err = net.Dial(v.Network(), v.String())
		if err == nil {
			return
		}
	}
	return
}
