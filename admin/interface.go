package admin

import "github.com/Shopify/sarama"

type Admin interface {
	Client()
	Admin()
	ListTopic() (map[string]sarama.TopicDetail, error)
	//CreateTopic()
	//AddPartition()
	//CreateSasl()
	//CreateAcl()
	//Producer()
	//Consumer()
}
