package ProducerAbout

import (
	"github.com/Shopify/sarama"
)

type KafkaProducer struct {
	Sp    *sarama.SyncProducer
	Topic string
}

func (kp *KafkaProducer) Send(msg_str string) (partition int32, offset int64, err error) {

	kafka_msg := &sarama.ProducerMessage{
		Topic: kp.Topic,
		//newManualPartitioner: true  #是否手动分配分区
		//如果手动分区选择true ，需要手动设置分区 增加如下代码
		//Partition: int32(your_partition_number),  // 设置分区号
		Value: sarama.StringEncoder(msg_str),
	}
	partition, offset, err = (*kp.Sp).SendMessage(kafka_msg)
	return
}

// 批量发送
func (kp *KafkaProducer) SendBatch(msg_strs []string) (err error) {

	var kafka_msgs []*sarama.ProducerMessage

	for _, msg_str := range msg_strs {
		kafka_msg := &sarama.ProducerMessage{
			Topic: kp.Topic,
			Value: sarama.StringEncoder(msg_str),
		}
		kafka_msgs = append(kafka_msgs, kafka_msg)
	}

	err = (*kp.Sp).SendMessages(kafka_msgs)
	return
}
