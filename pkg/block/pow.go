package block

import (
	"fmt"
	"rythorium/pkg/transaction"
	"time"
)

/*
A miner creates a new block and begins to attempt to mine said block
*/
func CreateAndMineBlock(previousBlock Block, data []transaction.Transaction, difficulty int) Block {
	newBlock := Block{
		Index:        previousBlock.Index + 1,
		Timestamp:    time.Now().UnixMilli(),
		PreviousHash: previousBlock.PreviousHash,
		Nonce:        0,
		Transactions: data,
	}

	target := ""

	for i := 0; i < difficulty; i++ {
		target += "0" // difficulty of mining represented by number of leading 0s
	}

	for {
		newBlock.Hash = newBlock.CalculateHash()

		if newBlock.Hash[:difficulty] == target {
			fmt.Printf("New block mined with hash: %s\n", newBlock.Hash)
			break
		}

		newBlock.Nonce++
	}

	return newBlock
}
