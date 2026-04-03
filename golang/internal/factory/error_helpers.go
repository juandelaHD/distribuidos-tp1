package factory

import (
	m "github.com/7574-sistemas-distribuidos/tp-mom/golang/internal/middleware"
	amqp "github.com/rabbitmq/amqp091-go"
)

func middlewareError(err error) error {
	if err == amqp.ErrClosed {
		return m.ErrMessageMiddlewareDisconnected
	}
	return m.ErrMessageMiddlewareMessage
}

func closeConnection(conn *amqp.Connection, err error) error {
	_ = conn.Close()
	return middlewareError(err)
}

func closeChannelAndConnection(ch *amqp.Channel, conn *amqp.Connection, err error) error {
	_ = ch.Close()
	_ = conn.Close()
	return middlewareError(err)
}
