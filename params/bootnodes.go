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

import "github.com/ethereum/go-ethereum/common"

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main Ethereum network.
var MainnetBootnodes = []string{
	// Wanshi mainnet Bootnodes
	"enode://09f8e13869dbb4171700b4a3f45ed8ff2a1aa4205ce18512b9f26d85f7c0e195ca7a52e2225f8de4166d2c208e9b905f53262c0c02961d95edb06216b302d8a1@146.190.36.212:30303",   // bootnode-aws-ap-southeast-1-001
	"enode://0010958b9e070cc56cc640a522e0f24faae50815f53ea2c4119db0460606bf04f39be1a7ed5fd2de47c99774c891ecf01d66a9ecb1c3fc81f8de105d8af6f6f0@164.92.68.57:30303",     // bootnode-aws-us-east-1-001
	"enode://cfa84b723e63a954db4b2b91b03556ca58264a33f7646020b5c1fff974c7057d5d8f6f7accfdadfff611cf39a0d72110baaa31749c8ebb0e9dd7ed3fc6f3a8a3@127.0.0.1:30306",     // bootnode-aws-us-east-1-001
	"enode://96c3b589a2d2a052c481586fdb509ca80baa60cddcd37f82f2b193fe272907681f1feadad18436b07c322f3488a13bc1c7b0a44d1e7fb7cfdebe780ea7a7a150@127.0.0.1:30307",     // bootnode-aws-us-east-1-001
}

// RopstenBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Ropsten test network.
var TestnetBootnodes = []string{
	"enode://30b7ab30a01c124a6cceca36863ece12c4f5fa68e3ba9b0b51407ccc002eeed3b3102d20a88f1c1d3c3154e2449317b8ef95090e77b312d5cc39354f86d5d606@52.176.7.10:30303",    // US-Azure geth
	"enode://865a63255b3bb68023b6bffd5095118fcc13e79dcf014fe4e47e065c350c7cc72af2e53eff895f11ba1bbb6a2b33271c1116ee870f266618eadfc2e78aa7349c@52.176.100.77:30303",  // US-Azure parity
	"enode://6332792c4a00e3e4ee0926ed89e0d27ef985424d97b6a45bf0f23e51f0dcb5e66b875777506458aea7af6f9e4ffb69f43f3778ee73c81ed9d34c51c4b16b0b0f@52.232.243.152:30303", // Parity
	"enode://94c15d1b9e2fe7ce56e458b9a3b672ef11894ddedd0c6f247e0f1d3487f52b66208fb4aeb8179fce6e3a749ea93ed147c37976d67af557508d199d9594c35f09@192.81.208.223:30303", // @gpip
}



const dnsPrefix = "enrtree://AKA3AM6LPBYEUDMVNU3BSVQJ5AD45Y7YPOHJLEF6W26QOE4VTUDPE@"

// KnownDNSNetwork returns the address of a public DNS-based node list for the given
// genesis hash and protocol. See https://github.com/ethereum/discv4-dns-lists for more
// information.
func KnownDNSNetwork(genesis common.Hash, protocol string) string {
	var net string
	switch genesis {
	case MainnetGenesisHash:
		net = "mainnet"
	case TestnetGenesisHash:
		net = "testnet"
	default:
		return ""
	}
	return dnsPrefix + protocol + "." + net + ".ethdisco.net"
}
