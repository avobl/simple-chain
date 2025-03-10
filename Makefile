CHAIN_ID=simple-chain
BINARY=simple-chaind

SENDER ?= alice
RECEIVER ?= bob

# Start the blockchain
serve:
	ignite chain serve

# Default kudos transaction
send-kudos:
	@$(BINARY) tx kudos send-kudos $(SENDER) $(RECEIVER) "Great job!" --from $(SENDER) --chain-id $(CHAIN_ID)

.PHONY: serve send-kudos
