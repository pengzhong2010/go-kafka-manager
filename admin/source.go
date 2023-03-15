package admin

import (
	"log"
	"os"
	"os/signal"

	"github.com/Shopify/sarama"
)

type AdminSource struct {
	client sarama.Client
	admin  sarama.ClusterAdmin
}

func NewAdminSource(config *sarama.Config) (a Admin, err error) {
	client, err := sarama.NewClient([]string{
		"127.0.0.1:9092",
	}, config)
	if err != nil {
		return
	}
	admin, err := sarama.NewClusterAdminFromClient(client)
	a = &AdminSource{
		client: client,
		admin:  admin,
	}
	return
}
func (a *AdminSource) Client() {}
func (a *AdminSource) Admin()  {}
func (a *AdminSource) ListAcls() ([]sarama.ResourceAcls, error) {
	return a.admin.ListAcls(sarama.AclFilter{})
}
func (a *AdminSource) ListTopic() (map[string]sarama.TopicDetail, error) {
	return a.admin.ListTopics()
}
func (a *AdminSource) CreateTopic(topicName string, numPartition, numReplication int64) (err error) {
	return a.admin.CreateTopic(topicName, &sarama.TopicDetail{NumPartitions: int32(numPartition), ReplicationFactor: int16(numReplication)}, false)
}
func (a *AdminSource) DoProducer() {
	config := sarama.NewConfig()
	config.Version = sarama.V3_3_1_0
	config.Producer.Retry.Max = 10
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true
	config.ClientID = "abc123ok"
	client, err := sarama.NewClient([]string{
		"127.0.0.1:9092",
	}, config)
	producer, err := sarama.NewSyncProducerFromClient(client)
	if err != nil {
		return
	}
	defer func() {
		if err = producer.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	for i := 0; i < 100; i++ {
		msg := &sarama.ProducerMessage{
			Topic: "test1",
			Value: sarama.StringEncoder("testing 123 gogo"),
			Key:   sarama.StringEncoder("123"),
		}
		partition, offset, err := producer.SendMessage(msg)
		if err != nil {
			log.Printf("failed to send msg,msg:%s\n", err.Error())
		} else {
			log.Printf("> message send to partition %d at offset %d\n", partition, offset)
		}
	}
}
func (a *AdminSource) DoConsumer1() {
	config := sarama.NewConfig()
	config.Version = sarama.V3_3_1_0
	client, err := sarama.NewClient([]string{
		"127.0.0.1:9092",
	}, config)
	consumer, err := sarama.NewConsumerFromClient(client)
	if err != nil {
		return
	}
	defer func() {
		if err = consumer.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	partitionConsumer, err := consumer.ConsumePartition("my_topic", 0, sarama.OffsetNewest)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := partitionConsumer.Close(); err != nil {
			log.Fatalln(err)
		}
	}()
	// Trap SIGINT to trigger a shutdown.
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	consumed := 0
ConsumerLoop:
	for {
		select {
		case msg := <-partitionConsumer.Messages():
			log.Printf("Consumed message offset %d\n", msg.Offset)
			consumed++
		case <-signals:
			break ConsumerLoop
		}
	}

	log.Printf("Consumed: %d\n", consumed)
}
func (a *AdminSource) DoConsumer2() {
	config := sarama.NewConfig()
	config.Version = sarama.V3_3_1_0
	client, err := sarama.NewClient([]string{
		"127.0.0.1:9092",
	}, config)
	consumer, err := sarama.NewConsumerFromClient(client)
	if err != nil {
		return
	}
	defer func() {
		if err = consumer.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	partitions, err := consumer.Partitions("test1")
	if err != nil {
		panic(err)
	}
	// close signals
	var signalsCloses []chan int
	for _, partition := range partitions {
		partitionConsumer, err := consumer.ConsumePartition("my_topic", partition, sarama.OffsetNewest)
		if err != nil {
			panic(err)
		}
		signalsClose := make(chan int)
		signalsCloses = append(signalsCloses, signalsClose)
		go func() {
			//for msg:=range partitionConsumer.Messages(){
			//	log.Printf("message partition:%d,offset:%d,")
			//}
		ConsumerLoop:
			for {
				select {
				case msg := <-partitionConsumer.Messages():
					log.Printf("Consumed message offset %d\n", msg.Offset)
				case <-signalsClose:
					break ConsumerLoop
				}
			}
		}()
	}
	// Trap SIGINT to trigger a shutdown.
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)
	<-signals
	for _, signalsClose := range signalsCloses {
		signalsClose <- 0
	}

}
