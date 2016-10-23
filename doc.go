// Copyright (c) 2016 Øystein Andersen
// Use of this source code is governed by MIT license found at
// https://github.com/oand/srv/blob/master/LICENSE

// Package srv simplify connecting to services by resolving SRV query and
// then connect to it in order of priority.
package srv

/*
A complete list of Service Name and Transport Protocol Port Number is found at:
http://www.iana.org/assignments/service-names-port-numbers
*/
