package main

import (
	"blockChain/blockchain"
	"fmt"
	"strconv"
)

func main() {

	chain := blockchain.InitBlockChain()

	//Creating a wallet for Alice.
	aliceWallet, err := blockchain.NewWallet()

	if err != nil {
		fmt.Println("Error creating Alice's wallet: ", err)
		return
	}

	fmt.Println("Aliced's wallet created successfully")

	//Creating a wallet for Bob.

	bobWallet, err := blockchain.NewWallet()

	if err != nil {
		fmt.Println("Error creating Bob's wallet :", err)
		return
	}

	fmt.Println("Bob's wallet created successfully")

	//Create a transaction from Alice to Bob
	tx := &blockchain.Transaction{Sender: aliceWallet.PublicKey.N.String(), Receiver: bobWallet.PublicKey.N.String(), Amount: 5.0}

	fmt.Println("Alice to Bob Transaction created successfully")

	// sign the transaction with Alice's wallet.
	signature, err := aliceWallet.SignTransaction(tx)

	if err != nil {
		fmt.Println("Error signing the transaction:", err)
		return
	}

	//Verity the transaction using Alice's wallet, public key, and the signature.

	err = blockchain.VerifyTransaction(tx, aliceWallet.PublicKey, signature)

	if err != nil {
		fmt.Println("Transaction verification failed:", err)
		return
	}

	fmt.Println("Transaction Verified Successfully")

	// add the verified transaction to the blockchain.

	chain.AddBlock("Block2", "Alice", []*blockchain.Transaction{tx})

	fmt.Println()

	for _, block := range chain.Blocks {
		fmt.Printf("previous hash : %X\n", block.PrevHash)
		fmt.Printf("Data in Bokck : %s\n", block.Data)
		fmt.Printf("Hash of block : %x\n", block.Hash)

		pow := blockchain.NewProofOfWork(block)
		fmt.Printf("IsValidPoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()

		fmt.Println("Transactions:")

		for _, tx := range block.Transactions {
			fmt.Printf("Sender: %s\n", tx.Sender)
			fmt.Printf("Receiver: %s\n", tx.Receiver)
			fmt.Printf("Amount: %f\n", tx.Amount)
			fmt.Printf("Coinbase: %t\n", tx.Coinbase)
			fmt.Println()
		}
		fmt.Println()
	}

}
