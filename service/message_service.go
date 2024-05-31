package service

import (
	"fmt"
	"practice/chaincode/entity"
	"practice/chaincode/repository"

	"github.com/hyperledger-labs/cckit/router"
)

func GetMessage(ctx router.Context) (interface{}, error) {
	args := ctx.GetArgs()
	if args == nil {
		return nil, fmt.Errorf("args value is nil")
	}

	data, err := repository.ReadMessage(ctx, string(args[1]))
	if err != nil {
		return nil, err
	}

	return data, nil
}

func SaveMessage(ctx router.Context) (interface{}, error) {

	args := ctx.GetArgs()
	if args == nil {
		return nil, fmt.Errorf("args value is nil")
	}

	message := new(entity.Request)
	err := message.ToObject(args[1])
	if err != nil {
		fmt.Printf("convert error %s\n", err.Error())
		return nil, err
	}

	err = repository.CreateMessage(ctx, message.Msg)
	if err != nil {
		return nil, err
	}

	return message, nil
}

func UpdateMessage(ctx router.Context) (interface{}, error) {

	args := ctx.GetArgs()
	if args == nil {
		return nil, fmt.Errorf("args value is nil")
	}

	message := new(entity.Message)
	err := message.ToObject(args[1])
	if err != nil {
		return nil, err
	}

	err = repository.UpdateMessage(ctx, message)
	if err != nil {
		return nil, err
	}

	return message, nil
}

func DeleteMessage(ctx router.Context) (interface{}, error) {

	title := ctx.ParamString(`title`)

	err := repository.DeleteMessage(ctx, title)
	if err != nil {
		return nil, err
	}

	return "delete success", nil
}
