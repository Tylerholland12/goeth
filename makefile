start-node:
	ganache-cli

transaction:
	go run transaction/block_subscribe.go
	go run transaction/blocks.go
	go run transaction/transaction_raw_create.go
	go run transaction/transactions.go
	go run transaction/transfer_eth.go
	go run transaction/transfer_tokens.go

account:
	go run account/account_balance.go
	go run account/address.go
	go run account/address_check.go
	go run account/generate_wallet.go
	go run account/keystore.go