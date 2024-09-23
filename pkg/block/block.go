package block

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"rythorium/pkg/transaction"
	"time"
)

type Block struct {
	Index        int
	Timestamp    int64
	PreviousHash string
	Hash         string
	Nonce        int
	Transactions []transaction.Transaction
}

func (b *Block) CalculateHash() string {
	blockData := struct {
		Index        int
		Timestamp    int64
		PreviousHash string
		Nonce        int
		Transactions []transaction.Transaction
	}{
		Index:        b.Index,
		Timestamp:    b.Timestamp,
		PreviousHash: b.PreviousHash,
		Nonce:        b.Nonce,
		Transactions: b.Transactions,
	}

	blockBytes, err := json.Marshal(blockData)
	if err != nil {
		panic(err)
	}

	hash := sha256.Sum256(blockBytes)

	return hex.EncodeToString(hash[:])
}

/*
A genesis block serves as the initial block created in the chain, hardcoded values are used
as we have no other blocks to provide hashes etc
*/
func CreateGenesisBlock() *Block {
	genesisData := []transaction.Transaction{
		{
			Sender:    "Genesis",
			Recipient: "MasterAddress",
			Amount:    50.0,
			Timestamp: time.Now().UnixMilli(),
			Signature: []byte{},
		},
	}

	genesisBlock := &Block{
		Index:        0,
		Timestamp:    time.Now().UnixMilli(),
		PreviousHash: "0",
		Nonce:        0,
		Transactions: genesisData,
	}

	genesisBlock.Hash = genesisBlock.CalculateHash()

	return genesisBlock
}

/*
Create a new block following on from the previous block in the chain
*/
func CreateNewBlock(previousBlock Block, data []transaction.Transaction) Block {
	newBlock := Block{
		Index:        previousBlock.Index + 1,
		Timestamp:    time.Now().UnixMilli(),
		PreviousHash: previousBlock.PreviousHash,
		Nonce:        0,
		Transactions: data,
	}

	return newBlock
}
