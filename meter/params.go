// Copyright (c) 2018 The VeChainThor developers

// Distributed under the GNU Lesser General Public License v3.0 software license, see the accompanying
// file LICENSE or <https://www.gnu.org/licenses/lgpl-3.0.html>

package meter

import (
	"time"

	"github.com/ethereum/go-ethereum/params"
)

// Constants of block chain.
const (
	BlockInterval uint64 = 10 // time interval between two consecutive blocks.

	TxGas                     uint64 = 5000
	ClauseGas                 uint64 = params.TxGas - TxGas
	ClauseGasContractCreation uint64 = params.TxGasContractCreation - TxGas

	MinGasLimit          uint64 = 1000 * 1000
	InitialGasLimit      uint64 = 10 * 1000 * 1000 // InitialGasLimit gas limit value int genesis block.
	GasLimitBoundDivisor uint64 = 1024             // from ethereum
	GetBalanceGas        uint64 = 400              //EIP158 gas table
	SloadGas             uint64 = 200              // EIP158 gas table
	SstoreSetGas         uint64 = params.SstoreSetGas
	SstoreResetGas       uint64 = params.SstoreResetGas

	MaxTxWorkDelay uint32 = 30 // (unit: block) if tx delay exceeds this value, no energy can be exchanged.

	MaxBlockProposers uint64 = 101

	TolerableBlockPackingTime = 500 * time.Millisecond // the indicator to adjust target block gas limit

	MaxStateHistory = 65535 // max guaranteed state history allowed to be accessed in EVM, presented in block number
)
