package main

import (
	"log"

	"github.com/it-chain/engine/common/command"
	"github.com/it-chain/engine/common/rabbitmq/rpc"
	"github.com/it-chain/engine/conf"
	"github.com/it-chain/engine/txpool"
	"github.com/it-chain/midgard"
	"github.com/rs/xid"
)

func main() {

	config := conf.GetConfiguration()
	client := rpc.NewClient(config.Engine.Amqp)
	defer client.Close()

	txCreateCommand := command.CreateTransaction{
		CommandModel: midgard.CommandModel{
			ID: xid.New().String(),
		},
		ICodeID:  "1",
		Jsonrpc:  "2.0",
		Method:   "invoke",
		Args:     []string{},
		Function: "initA",
	}

	err := client.Call("transaction.create", txCreateCommand, func(transaction txpool.Transaction, err rpc.Error) {
		log.Printf("created transaction id [%s]", transaction.ID)
	})

	if err != nil {
		log.Println(err)
	}
}
