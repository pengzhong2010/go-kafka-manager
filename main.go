package main

import (
	"github.com/Shopify/sarama"
	"github.com/pengzhong2010/go-kafka-manager/admin"
)

func main() {
	config := sarama.NewConfig()
	config.Version = sarama.V3_3_1_0
	//config.Net.SASL.Enable = true
	//config.Net.SASL.Mechanism = "PLAIN"
	//config.Net.SASL.User = "admin"
	//config.Net.SASL.Password = "admin"
	admin, err := admin.NewAdminSource(config)
	if err != nil {
		panic(err)
	}
	admin.ListTopic()
}
