package products_lib

import (
	"encoding/json"
	"errors"
	"github.com/djumanoff/amqp"
)

type AMQPEndpointFactory struct {
	productService ProductService
}

func NewAMQPEndpointFactory(productService ProductService) *AMQPEndpointFactory {
	return &AMQPEndpointFactory{productService: productService}
}

type ErrorSt struct {
	Text string `json:"text"`
}

func (fac *AMQPEndpointFactory) GetProductByIdAMQPEndpoint() amqp.Handler {
	return func(message amqp.Message) *amqp.Message {
		cmd := &GetProductByIdCommand{}

		if err := json.Unmarshal(message.Body, cmd); err != nil {
			return AMQPError(&ErrorSt{err.Error()})
		}

		if cmd.Id == 0 {
			return AMQPError(&ErrorSt{errors.New("not product id").Error()})
		}

		resp, err := cmd.Exec(fac.productService)
		if err != nil {
			return AMQPError(&ErrorSt{err.Error()})
		}
		return OK(resp)
	}
}

func (fac *AMQPEndpointFactory) CreateProductAMQPEndpoint() amqp.Handler {
	return func(message amqp.Message) *amqp.Message {
		cmd := &CreateProductCommand{}
		if err := json.Unmarshal(message.Body, cmd); err != nil {
			return AMQPError(err)
		}
		resp, err := cmd.Exec(fac.productService)
		if err != nil {
			return AMQPError(err)
		}
		return OK(resp)
	}
}

func OK(d interface{}) *amqp.Message {
	data, _ := json.Marshal(d)
	return &amqp.Message{Body: data}
}

func AMQPError(e interface{}) *amqp.Message {
	errObj, _ := json.Marshal(e)
	return &amqp.Message{Body: errObj}
}
