package clique

import (
	"math/big"

	"github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/trie"
)

// BlockReward is the reward in wei distributed each block.
var BlockReward = big.NewInt(5e+18)

// Finalize implements consensus.Engine, ensuring no uncles are set, but this does give rewards.
func (_ *Clique) Finalize(chain consensus.ChainHeaderReader, header *types.Header, state *state.StateDB, txs []*types.Transaction,
	receipts []*types.Receipt)  (*types.Block, error){
	cfg := chain.Config()
	signerReward := BlockReward
	// Reward the signer.
	state.AddBalance(header.Coinbase, signerReward)

	header.Root = state.IntermediateRoot(cfg.IsEIP158(header.Number))
	header.UncleHash = types.CalcUncleHash(nil)

	// Assemble and return the final block for sealing
	return types.NewBlock(header, txs, nil, receipts, trie.NewStackTrie(nil)), nil
}
