package block

import (
	"fmt"
	"log"
	"time"
)

type Block struct {
	nonce        int
	transactions []string
	previousHash string
	timestamp    int64
}

func (b *Block) Print() {
	fmt.Printf("Previous Hash: %s\n", b.timestamp)
	fmt.Printf("Nonce: %d\n", b.nonce)
	fmt.Printf("Timestamp : %d\n", b.timestamp)
	fmt.Printf("Transactions: %s\n", b.transactions)
}

// Chain Creation of the blockchain
type Chain struct {
	transactionPool []string
	chain           []*Block
}

func (bc *Chain) CreateBlock(nonce int, previousHash string) *Block {
	b := NewBlock(nonce, previousHash)

	bc.chain = append(bc.chain, b)
	return b
}

func (bc *Chain) Chain() []*Block {
	return bc.chain
}

func NewBlockChain() *Chain {
	bc := Chain{}
	bc.CreateBlock(0, "init hash")

	return &bc

}

func NewBlock(nonce int, previousHash string) *Block {
	return &Block{
		timestamp:    time.Now().UnixNano(),
		nonce:        nonce,
		previousHash: previousHash,
	}
}

func (bc *Chain) Print() {
	for i, block := range bc.chain {
		fmt.Printf("Chain %d \n", i)
		block.Print()
	}
}

func init() {
	log.SetPrefix("Blockchain: ")
}
func main() {
	bc := NewBlockChain()
	bc.CreateBlock(5, "hash 1")
	bc.CreateBlock(6, "hash 2")
	bc.Print()
}
