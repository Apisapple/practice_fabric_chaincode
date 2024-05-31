package main

import (
	"practice/chaincode/entity"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	testcc "github.com/hyperledger-labs/cckit/testing"
	expectcc "github.com/hyperledger-labs/cckit/testing/expect"
)

func TestOpenDID(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, `Practice Chaincode`)
}

var (
	Payload = []*entity.Message{
		{
			Title:   "title_1",
			Content: "content_1",
		},
		{
			Title:   "title_2",
			Content: "content_2",
		},
		{
			Title:   "title_3",
			Content: "content_3",
		},
	}
	Req = []*entity.Request{
		{
			Msg: Payload[0],
		},
	}
)

var _ = Describe(`PracticeChaincode`, func() {

	cc := testcc.NewMockStub(`practice_chaincode`, NewPracticeChaincode())

	Describe("PracticeChaincode", func() {
		It("Allow everyone to save message information", func() {
			expectcc.ResponseOk(cc.Invoke("save", Req[0]))
		})

		It("Allow everyone to retrieve message information", func() {
			message := expectcc.PayloadIs(cc.Query("get",
				Payload[0].Title), &entity.Message{}).(entity.Message)

			Expect(message.Title).To(Equal(Payload[0].Title))
		})
	})
})
