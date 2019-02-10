package main

import (
	"github.com/spf13/viper"
	"log"
	"fmt"
	"varnost-core/Varnost-Alert/conf"
)

// https://github.com/vsouza/go-kafka-example

type Alert struct {
    ID  string
    Severity int
    Source string
    Detail string
}

var alerts []Alert

// create a new item
//func CreateAlert() {
 //  return
//}

func (a Alert) GenerateAlertHash() (string){
    return "foo" //todo
}

// Handler for email
// Handler for slack
// Handler for pagerduty


func main() {
    viper.SetConfigName("config")
	viper.AddConfigPath(".")
	var configuration conf.Configuration

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	err := viper.Unmarshal(&configuration)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}
	foo := viper.Get("kafka.BrokerList")
	fmt.Printf("%s", foo)
    //router := mux.NewRouter()

    //router.HandleFunc("/alert", CreateAlert).Methods("POST")

    //log.Fatal(http.ListenAndServe(":8000", router))
	//config.json := sarama.NewConfig()
	//config.json.Consumer.Return.Errors = true
	//brokers := *configuration.kafka.BrokerList
	//master, err := sarama.NewConsumer(brokers, config.json)
	//if err != nil {
	//	panic(err)
	//}
	//defer func() {
	//	if err := master.Close(); err != nil {
	//		panic(err)
	//	}
	//}()
	//consumer, err := master.ConsumePartition(*topic, 0, sarama.OffsetOldest)
	//if err != nil {
	//	panic(err)
	//}
	//signals := make(chan os.Signal, 1)
	//signal.Notify(signals, os.Interrupt)
	//doneCh := make(chan struct{})
	//go func() {
	//	for {
	//		select {
	//		case err := <-consumer.Errors():
		//		fmt.Println(err)
		//	case msg := <-consumer.Messages():
	//			*messageCountStart++
	//			fmt.Println("Received messages", string(msg.Key), string(msg.Value))
	//		case <-signals:
	//			fmt.Println("Interrupt is detected")
	//			doneCh <- struct{}{}
	//		}
	//	}
	//}()
	//<-doneCh
//fmt.Println("Processed", *messageCountStart, "messages")

}