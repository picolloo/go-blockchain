package main

import (
	"crypto/sha256"
	"fmt"
)

type Block struct {
  Hash []byte
  Data []byte
  PrevHash []byte
}

func (b *Block) GerenateHash() {
  payload := append(b.Data, b.PrevHash...)
  hash := sha256.Sum256(payload)
  b.Hash = hash[:]
}

func NewBlock(data []byte, prevHash []byte) *Block {
  block := Block{
    Data: data,
    PrevHash: prevHash,
  }
  block.GerenateHash()
  return &block
}

func (b *Block) Serialize() string {
  return fmt.Sprintf("PrevHash: %s\nData: %s\nHash: %s\n", b.PrevHash, b.Data, b.Hash)
}

type BlockChain struct {
  Blocks []*Block
}

func (c *BlockChain) AddBlock(block *Block) {
  lastBlock := c.Blocks[len(c.Blocks) - 1]
  block.PrevHash = lastBlock.Hash
  c.Blocks = append(c.Blocks, block)
}

func (c *BlockChain) ListBlock() string {
  var output string

  for _, block := range(c.Blocks) {
    output += block.Serialize()
    output += "----------------\n"
  }
  return output
}

func main() {

  firstBlock := NewBlock([]byte("first block"), []byte{})
  chain := BlockChain{
    Blocks: []*Block{firstBlock},
  }

  block1 := NewBlock([]byte("block A"),firstBlock.Hash)
  chain.AddBlock(block1)
  block2 := NewBlock([]byte("block B"),block1.Hash)
  chain.AddBlock(block2)
  block3 := NewBlock([]byte("block C"),block2.Hash)
  chain.AddBlock(block3)

  fmt.Print(chain.ListBlock())

}
