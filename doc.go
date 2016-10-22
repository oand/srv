// Copyright (c) 2016 Øystein Andersen
// Use of this source code is governed by MIT license found at
// https://github.com/oand/srv/blob/master/LICENSE

// Package srv simplify connecting to services by resolving SRV query and
// then connect to it in order priority, randomized by weight within a priority.
package srv
