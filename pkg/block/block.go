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
	Data         []transaction.Transaction
}

func (b *Block) CalculateHash() string {
	blockData := struct {
		Index        int
		Timestamp    int64
		PreviousHash string
		Nonce        int
		Data         []transaction.Transaction
	}{
		Index:        b.Index,
		Timestamp:    b.Timestamp,
		PreviousHash: b.PreviousHash,
		Nonce:        b.Nonce,
		Data:         b.Data,
	}

	blockBytes, err := json.Marshal(blockData)
	if err != nil {
		panic(err)
	}

	hash := sha256.Sum256(blockBytes)

	return hex.EncodeToString(hash[:])
}

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
		Data:         genesisData,
	}

	genesisBlock.Hash = genesisBlock.CalculateHash()

	return genesisBlock
}
