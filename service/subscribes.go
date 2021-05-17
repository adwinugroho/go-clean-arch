package service

import (
	"encoding/json"
	"go-clean-arch/config"
	"go-clean-arch/pkg"
	"go-clean-arch/repository"
	"log"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
)

// service subs to channel
type (
	subscribe struct {
		repo repository.OrderRepository
	}

	payload struct {
		PaymentID   string `json:"paymentId"`
		AlternateID string `json:"altId"`
	}
)

func ListenNats(repository repository.OrderRepository) {
	var opts = config.GetNatsConfig()
	var nc, err = opts.Connect()
	if err != nil {
		panic(err)
	}

	sc, err := stan.Connect(config.CLUSTERID, config.HOSTNAME, stan.NatsConn(nc), stan.SetConnectionLostHandler(func(s stan.Conn, reason error) {
		log.Printf("[SetReconnectHandler]: Connection lost, cause: %+v \n", reason)
		panic(reason)
	}))
	if err != nil {
		panic(err)
	}

	nc.SetReconnectHandler(func(ncr *nats.Conn) {
		log.Println("Reconnecting...")
		nsc, err := stan.Connect(config.CLUSTERID, config.HOSTNAME, stan.NatsConn(ncr), stan.SetConnectionLostHandler(func(s stan.Conn, reason error) {
			log.Printf("[SetReconnectHandler]: Connection lost, cause: %+v \n", err)
			panic(err)
		}))
		if err != nil {
			panic(err)
		}

		subscribe := &subscribe{repository}
		subscribe.SubsCHPayment(nsc)
	})

	nc.SetDisconnectErrHandler(func(nc *nats.Conn, err error) {
		log.Println("Disconnected.. ", err)
	})

	nc.SetClosedHandler(func(nc *nats.Conn) {
		log.Println("Connection closed.")
		ListenNats(repository)
	})

	subscribe := &subscribe{repository}
	subscribe.SubsCHPayment(sc)
}

func (conn *subscribe) SubsCHPayment(sc stan.Conn) {
	var subs, err = sc.Subscribe(config.CH_PAYMENT, func(msg *stan.Msg) {
		var payload payload
		if err := json.Unmarshal(msg.Data, &payload); err != nil {
			log.Printf(" Error unmarshaling object, cause: %+v \n", err)
			return
		}

		response, err := pkg.Post(&pkg.Request{
			Protocol: "https",
			Port:     443,
			Path:     "/example/path",
			Host:     "example-host",
			Body: map[string]interface{}{
				"alternateId": payload.AlternateID,
				"paymentId":   payload.PaymentID,
			},
		})
		if err != nil {
			log.Printf("error when requesting to PHRCAREVOAPI, cause %v", err)
			return
		}
		if !response.Status {
			log.Printf("error when requesting to PHRCAREVOAPI, cause %+v", response)
			return
		}

		if err := conn.repo.DeleteByID(payload.AlternateID); err != nil {
			log.Printf("error when deleting business account, cause %+v", err)
			return
		}
		msg.Ack()
	}, stan.SetManualAckMode(), stan.DurableName("payment"))
	if err != nil {
		panic(err)
	}

	log.Printf("Subcribe %s %v", config.CH_PAYMENT, subs.IsValid())
}
