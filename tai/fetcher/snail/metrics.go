// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

// Contains the metrics collected by the sfetcher.

package snailfetcher

import (
	"github.com/taiyuechain/taipublicchain/metrics"
)

var (
	propAnnounceInMeter    = metrics.NewRegisteredMeter("tai/sfetcher/prop/announces/in", nil)
	propAnnounceOutTimer   = metrics.NewRegisteredTimer("tai/sfetcher/prop/announces/out", nil)
	propAnnounceDropMeter  = metrics.NewRegisteredMeter("tai/sfetcher/prop/announces/drop", nil)
	propAnnounceDOSMeter   = metrics.NewRegisteredMeter("tai/sfetcher/prop/announces/dos", nil)
	propBroadcastInMeter   = metrics.NewRegisteredMeter("tai/sfetcher/prop/broadcasts/in", nil)
	propBroadcastOutTimer  = metrics.NewRegisteredTimer("tai/sfetcher/prop/broadcasts/out", nil)
	propBroadcastDropMeter = metrics.NewRegisteredMeter("tai/sfetcher/prop/broadcasts/drop", nil)
	propBroadcastDOSMeter  = metrics.NewRegisteredMeter("tai/sfetcher/prop/broadcasts/dos", nil)
	headerFetchMeter       = metrics.NewRegisteredMeter("tai/sfetcher/fetch/headers", nil)
	bodyFetchMeter         = metrics.NewRegisteredMeter("tai/sfetcher/fetch/bodies", nil)

	headerFilterInMeter  = metrics.NewRegisteredMeter("tai/sfetcher/filter/headers/in", nil)
	headerFilterOutMeter = metrics.NewRegisteredMeter("tai/sfetcher/filter/headers/out", nil)
	bodyFilterInMeter    = metrics.NewRegisteredMeter("tai/sfetcher/filter/bodies/in", nil)
	bodyFilterOutMeter   = metrics.NewRegisteredMeter("tai/sfetcher/filter/bodies/out", nil)
)
