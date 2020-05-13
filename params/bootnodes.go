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

package params

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main Taichain network.
var MainnetBootnodes = []string{
	"enode://fb331ff6aded86b393d9de2f9c449d313b356af0c4c0b9500e0f6c51bcb4ed31ca45dc2ab64c6182d1876eb9e3fd073d488277a40a6d357bc6e63350a2e00ffc@101.132.183.35:30313", // CN
}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Ropsten test network.
var TestnetBootnodes = []string{
	"enode://a395d2799c1e63307b9a5ecc44729e9ba2fb8fa6d64e362e8498ce9aba85b7b405755ad28bd662a9a48d941bbbfe18d29e0ea46105258110e2318fd6faab8c09@39.108.212.229:30313", // CN
}

// DevnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the dev Taichain network.
var DevnetBootnodes = []string{
	"enode://f1ce2725b0e5cf403293be25ce94c222d8f4e6e7e4e2881559382a8fbfb64934923467ca182985f8391c6f65d79a717c13df4fb2a53ccd8aba51e5638d6da6a7@39.98.202.190:30314",
}
