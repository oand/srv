// Copyright (c) 2016 Øystein Andersen
// Use of this source code is governed by MIT license found at
// https://github.com/oand/srv/blob/master/LICENSE

package srv

import (
	"errors"
	"net"
	"strconv"
	"strings"
)

// Lookup resolves an SRV query of the given protocol and address
//
// Values for protocol are "tcp" or "udp"
// address is in the form: domain:service
//
// Examples:
//	Lookup("tcp", "example.com:xmpp-client")
//	Lookup("udp", "example.com:stun")
//
// The returned records are net.Addr containing the CNAME and port of the service,
// sorted by priority and randomized by weight within a priority.
func Lookup(protocol, address string) ([]net.Addr, error) {
	s := strings.Split(address, ":")
	if len(s) != 2 {
		return nil, errors.New("address is in the form: domain:service")
	}

	return LookupSRV(s[1], protocol, s[0])
}

// LookupSRV resolves an SRV query of the given service, protocol, and domain name.
// Values for protocol are "tcp" or "udp"
//
// LookupSRV constructs the DNS name to look up following RFC 2782.
// That is, it looks up _service._proto.name. To accommodate services publishing SRV records under non-standard names.
//
// The returned records are net.Addr containing the CNAME and port of the service,
// sorted by priority and randomized by weight within a priority.
func LookupSRV(service, protocol, domain string) ([]net.Addr, error) {
	_, srv, err := net.LookupSRV(service, protocol, domain)
	if err != nil {
		return nil, err
	}

	var adr []net.Addr
	for _, v := range srv {
		adr = append(adr, &lookupAddr{
			network: protocol,
			cname:   strings.Trim(v.Target, "."),
			port:    strconv.Itoa(int(v.Port))})
	}
	return adr, nil
}

type lookupAddr struct{ network, cname, port string }

func (l *lookupAddr) Network() string { return l.network }
func (l *lookupAddr) String() string  { return net.JoinHostPort(l.cname, l.port) }
