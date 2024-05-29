package main

import (
	"fmt"
	"practice/chaincode/service"

	"github.com/hyperledger-labs/cckit/router"
	"github.com/hyperledger/fabric-chaincode-go/shim"
)

func NewPracticeChaincode() *router.Chaincode {
	r := router.New(`practice_chaincode`)

	r.Query("get_msg", service.GetMessage)

	r.Invoke("save_msg", service.SaveMessage)

	r.Invoke("update_msg", service.UpdateMessage)

	r.Invoke("delete_msg", service.DeleteMessage)

	return router.NewChaincode(r)
}

func main() {
	err := shim.Start(NewPracticeChaincode())
	if err != nil {
		fmt.Printf(`error start new chaincode : %s`, err)
	}
}
