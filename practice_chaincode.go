package main

import (
	"fmt"
	"practice/chaincode/service"

	"github.com/hyperledger-labs/cckit/router"
	"github.com/hyperledger/fabric-chaincode-go/shim"
)

func NewPracticeChaincode() *router.Chaincode {
	r := router.New(`chaincode`)

	r.Query("get", service.GetMessage)

	r.Invoke("save", service.SaveMessage)

	r.Invoke("update", service.UpdateMessage)

	r.Invoke("delete", service.DeleteMessage)

	return router.NewChaincode(r)
}

func main() {
	err := shim.Start(NewPracticeChaincode())
	if err != nil {
		fmt.Printf(`error start new chaincode : %s`, err)
	}
}
