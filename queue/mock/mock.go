package mock

import (
	"main/queue"

	"github.com/IBM/sarama"
)

func NewMockProducerHandler() (queue.ProducerHandler) {
	return queue.ProducerHandler{
		Producer: &MockProducer{},
	}
}

type MockProducer struct{}


func (*MockProducer) SendMessage(msg *sarama.ProducerMessage) (partition int32, offset int64, err error) {
	return 0, 0, nil
}

func (*MockProducer) SendMessages(msgs []*sarama.ProducerMessage) error {
	return nil
}

func (*MockProducer) Close() error {
	return nil
}

func (*MockProducer) TxnStatus() sarama.ProducerTxnStatusFlag {
	return 0
}

func (*MockProducer) IsTransactional() bool {
	return false
}

func (*MockProducer) BeginTxn() error {
	return nil
}

func (*MockProducer) CommitTxn() error {
	return nil
}

func (*MockProducer) AbortTxn() error {
	return nil
}

func (*MockProducer) AddOffsetsToTxn(offsets map[string][]*sarama.PartitionOffsetMetadata, groupId string) error {
	return nil
}

func (*MockProducer) AddMessageToTxn(msg *sarama.ConsumerMessage, groupId string, metadata *string) error {
	return nil
}