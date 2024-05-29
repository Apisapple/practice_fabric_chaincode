package service

import (
	"practice/chaincode/entity"
	"practice/chaincode/repository"

	"github.com/hyperledger-labs/cckit/router"
)

func GetMessage(ctx router.Context) (interface{}, error) {
	title := ctx.ParamString(`title`)
	return repository.ReadMessage(ctx, title)
}

func SaveMessage(ctx router.Context) (interface{}, error) {

	messageByte := ctx.ParamBytes(`msg`)
	message := new(entity.Message)
	err := message.ToObject(messageByte)
	if err != nil {
		return nil, err
	}

	err = repository.CreateMessage(ctx, message)
	if err != nil {
		return nil, err
	}

	return message, nil
}

func UpdateMessage(ctx router.Context) (interface{}, error) {

	messageByte := ctx.ParamBytes(`msg`)
	message := new(entity.Message)
	err := message.ToObject(messageByte)
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
