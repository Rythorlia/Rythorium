package main

import (
	"fmt"
	"rythorium/pkg/block"
	"rythorium/pkg/transaction"
)

func main() {
	difficulty := 3

	genesisBlock := block.CreateGenesisBlock()
	fmt.Printf("Genesis Block: %+v\n", genesisBlock)

	transactions := []transaction.Transaction{
		{Sender: "Alice", Recipient: "Bob", Amount: 10.0, Timestamp: 1725971732000, Signature: []byte("test")},
		{Sender: "Niklas", Recipient: "Daniel", Amount: 25.0, Timestamp: 1725621731000, Signature: []byte("test")},
	}

	secondBlock := block.CreateAndMineBlock(*genesisBlock, transactions, difficulty)

	fmt.Printf("Finished mining block: %+v\n", secondBlock)
}
