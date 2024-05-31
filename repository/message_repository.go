package repository

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"practice/chaincode/entity"

	"github.com/hyperledger-labs/cckit/router"
)

func CreateMessage(ctx router.Context, msg *entity.Message) error {

	jsonData, err := msg.ToJson()
	if err != nil {
		ctx.State().Logger().Error(
			fmt.Sprintf("Failed to convert message to JSON: %v", err),
		)
		return err
	}
	ctx.State().Logger().Info(fmt.Sprintf("Message JSON: %s", jsonData))

	key := makeStateKey(msg.Title)
	ctx.State().Logger().Info(fmt.Sprintf("Generated state plantext: %s", msg.Title))
	ctx.State().Logger().Info(fmt.Sprintf("Generated state key: %s", key))

	err = ctx.State().Insert(key, jsonData)
	if err != nil {
		ctx.State().Logger().Error(fmt.Sprintf("Insert error for key %s with data %s: %v", key, jsonData, err))
		return err
	}

	return nil
}

func ReadMessage(ctx router.Context, title string) (interface{}, error) {

	key := makeStateKey(title)
	data, err := ctx.State().Get(key)
	if err != nil {
		return nil, fmt.Errorf("message_reopsitory error => {%w}", err)
	}

	return data, nil
}

func UpdateMessage(ctx router.Context, msg *entity.Message) error {
	key := makeStateKey(msg.Title)

	isExist, err := ctx.State().Exists(key)
	if err != nil {
		return err
	}

	if !isExist {
		return fmt.Errorf("cannot found message with key=%s", msg.Title)
	}

	jsonData, err := msg.ToJson()
	if err != nil {
		return err
	}

	ctx.State().Put(key, jsonData)

	return nil
}

func DeleteMessage(ctx router.Context, title string) error {
	key := makeStateKey(title)

	isExist, err := ctx.State().Exists(key)
	if err != nil {
		return err
	}

	if !isExist {
		return fmt.Errorf("cannot found message with key=%s", title)
	}

	return ctx.State().Delete(key)
}

func makeStateKey(title string) string {

	hash := sha256.New()
	hash.Write([]byte(title))
	md := hash.Sum(nil)

	return hex.EncodeToString(md)
}
