package main

import (
	"fmt"
	"rythorium/pkg/block"
)

func main() {
	genesisBlock := block.CreateGenesisBlock()
	fmt.Printf("Genesis Block: %+v\n", genesisBlock)
}
