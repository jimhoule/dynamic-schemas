package queue

import (
	"context"
	"fmt"

	"github.com/IBM/sarama"
)

type Handler = func(body []byte) error
type ConsumerGroup = sarama.ConsumerGroup
type ConsumerGroupSession = sarama.ConsumerGroupSession
type ConsumerGroupClaim = sarama.ConsumerGroupClaim

type ConsumerGroupHandler struct {
	ConsumerGroup ConsumerGroup
	Handlers      map[string]Handler
}

func NewConsumerGroupHandler(addresses []string, id string) (ConsumerGroupHandler, error) {
	consumerGroupHandler := ConsumerGroupHandler{}

	config := sarama.NewConfig()
	config.Consumer.Offsets.AutoCommit.Enable = false
	config.Consumer.Return.Errors = true

	consumerGroup, err := sarama.NewConsumerGroup(addresses, id, config)
	if err != nil {
		return consumerGroupHandler, err
	}

	consumerGroupHandler.ConsumerGroup = consumerGroup

	return consumerGroupHandler, nil
}

/**
 * NOTES:
 *  - Satisfies the sarama ConsumerGroupHandler interface
 *  - Runs at the beginning of a new session, before ConsumeClaim
 */
func (*ConsumerGroupHandler) Setup(ConsumerGroupSession) error {
	return nil
}

/**
 * NOTES:
 *  - Satisfies the sarama ConsumerGroupHandler interface
 *  - Runs at the end of a session, once all ConsumeClaim goroutines have exited but before the
 *    offsets are committed for the very last time
 */
func (*ConsumerGroupHandler) Cleanup(ConsumerGroupSession) error {
	return nil
}

/**
 * NOTES:
 *  - Satisfies the sarama ConsumerGroupHandler interface
 *  - Starts a consumer loop of ConsumerGroupClaim's Messages()
 *  - Once the Messages() channel is closed, the Handler must finish its processing
 *    loop and exit
 */
func (cgh *ConsumerGroupHandler) ConsumeClaim(session ConsumerGroupSession, claim ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		handler, ok := cgh.Handlers[message.Topic]

		// If not topic, continue
		if !ok {
			continue
		}

		// If topic, call associated handler
		err := handler(message.Value)
		if err != nil {
			fmt.Println("error: ", err)
		}

		session.MarkMessage(message, "")
	}

	return nil
}

/**
 * NOTES:
 *  - Custom add
 *  - Starts listening for incoming messages
 */
func (cgh *ConsumerGroupHandler) Listen(topics []string) error {
	for {
		err := cgh.ConsumerGroup.Consume(context.Background(), topics, cgh)
		if err != nil {
			return err
		}
	}
}
