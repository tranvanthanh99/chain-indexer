package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/kien6034/chain-indexer/bitcoin/indexer"
)

var PRIVATE_KEY string

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	PRIVATE_KEY = os.Getenv("PRIVATE_KEY")
	if PRIVATE_KEY == "" {
		log.Fatal("PRIVATE_KEY environment variable is not set.")
	}
}

func main() {
	// create indexer
	indexer := indexer.NewBitcoinClient(false)

	//Get address transactions
	utxos, err := indexer.GetAddressTransactions("tb1qsehk8q2v6d833dnf8fgsr5p8klfhysvd7cega5")

	if err != nil {
		panic(err)
	}

	log.Printf("response: %v", utxos)

	// w := wallet.NewBtcWallet(PRIVATE_KEY, false)
	// walletPubkey, _ := w.GetWifPubkeyAddress()

	// log.Printf("wallet pubkey: %s", walletPubkey)

	// r, err := w.SendTxWithMemo(*indexer, "tb1qjfaa5vvxt9m4sp9kqkcpzypkzydz2vcywqx9tm", 800, "0x5b5fDd1510F817Ece8bBD911d7028144522c4429", "97")
	// if err != nil {
	// 	panic(err)
	// }

	// log.Printf("response: %s", r)
}
