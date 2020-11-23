package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
)

func main() {
	address := common.HexToAddress("0xbeE0900C04159CBa4e154968047914487974d3f0")

	fmt.Println(address.Hex()) // 0x71C7656EC7ab88b098defB751B7401B5f6d8976F

	fmt.Println(address.Hash().Hex()) // 0x000000000000000000000000bee0900c04159cba4e154968047914487974d3f0

	fmt.Println(address.Bytes())      // [190 224 144 12 4 21 156 186 78 21 73 104 4 121 20 72 121 116 211 240]
}