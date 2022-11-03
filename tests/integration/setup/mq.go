package setup

import (
	"fmt"
	"log"

	"github.com/Dharitri-org/tw-go-libs/network/mq"
	"github.com/ory/dockertest"
)

var (
	mqResource *dockertest.Resource
)

func runMQContainer() error {
	var err error
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	mqResource, err = pool.Run("rabbitmq", "latest", nil)
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	if err = pool.Retry(func() error {
		return mq.Init(fmt.Sprintf("amqp://localhost:%s", mqResource.GetPort("5672/tcp")))
	}); err != nil {
		return err
	}
	return nil
}

func stopMQContainer() error {
	mq.Close()
	return mqResource.Close()
}
