// Copyright (c) 2016 Øystein Andersen
// Use of this source code is governed by MIT license found at
// https://github.com/oand/srv/blob/master/LICENSE

package srv

import "net"

// Dial resolves a DNS SRV query of the given address.
// Then it will attempt to connect to the service in order priority.
//
// Values for protocol are "tcp" or "udp",
// address is in the form: domain:service
//
// Examples:
//	conn, err := Dial("tcp", "example.com:xmpp-client")
//	conn, err := Dial("udp", "example.com:stun")
func Dial(protocol, address string) (net.Conn, error) {
	addr, err := Lookup(protocol, address)
	if err != nil {
		return nil, err
	}
	return dial(addr)
}

// DialSRV resolves a DNS SRV query of the given service, protocol and domain name.
// Then it will attempt to connect to the service in order priority.
//
// Values for protocol are "tcp" or "udp".
//
// Examples:
//	conn, err := DialSRV("xmpp-client", "tcp", "example.com")
//	conn, err := DialSRV("stun", "udp", "example.com")
func DialSRV(service, protocol, domain string) (net.Conn, error) {
	addr, err := LookupSRV(service, protocol, domain)
	if err != nil {
		return nil, err
	}
	return dial(addr)
}

// DialTCP resolves and connects the given address.
//
// Address is in the form: domain:service.
// If localAddr is not nil, it is used as the local address for the connection.
//
// Examples:
//	conn, err := DialTCP(nil, "example.com:xmpp-client")
func DialTCP(localAddr *net.TCPAddr, address string) (*net.TCPConn, error) {
	addr, err := Lookup("tcp", address)
	if err != nil {
		return nil, err
	}
	return dialTCP(localAddr, addr)
}

// DialUDP resolves and connects the given address.
//
// Address is in the form: domain:service.
// If localAddr is not nil, it is used as the local address for the connection.
//
// Examples:
//	conn, err := DialUDP(nil, "example.com:stun")
func DialUDP(localAddr *net.UDPAddr, address string) (*net.UDPConn, error) {
	addr, err := Lookup("udp", address)
	if err != nil {
		return nil, err
	}
	return dialUDP(localAddr, addr)
}

func dial(addr []net.Addr) (conn net.Conn, err error) {
	for _, v := range addr {
		if conn, err = net.Dial(v.Network(), v.String()); err == nil {
			return
		}
	}
	return
}

func dialTCP(localAddr *net.TCPAddr, addr []net.Addr) (conn *net.TCPConn, err error) {
	var a *net.TCPAddr

	for _, v := range addr {
		if a, err = net.ResolveTCPAddr(v.Network(), v.String()); err == nil {
			if conn, err = net.DialTCP(a.Network(), localAddr, a); err == nil {
				return
			}
		}
	}
	return
}

func dialUDP(localAddr *net.UDPAddr, addr []net.Addr) (conn *net.UDPConn, err error) {
	var a *net.UDPAddr

	for _, v := range addr {
		if a, err = net.ResolveUDPAddr(v.Network(), v.String()); err == nil {
			if conn, err = net.DialUDP(a.Network(), localAddr, a); err == nil {
				return
			}
		}
	}
	return
}
