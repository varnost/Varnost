package main

import (
    "encoding/json"
    "github.com/gorilla/mux"
    "log"
    "net/http"
    "github.com/Shopify/sarama"
)

var (
	brokerList        = kingpin.Flag("brokerList", "List of brokers to connect").Default("localhost:9092").Strings()
	topic             = kingpin.Flag("topic", "Topic name").Default("important").String()
	partition         = kingpin.Flag("partition", "Partition number").Default("0").String()
	offsetType        = kingpin.Flag("offsetType", "Offset Type (OffsetNewest | OffsetOldest)").Default("-1").Int()
	messageCountStart = kingpin.Flag("messageCountStart", "Message counter start from:").Int()
)

// Read config file passed as arg
//  Config file contains default mailer per severity
//    Also contains granular source => destination mapping if necessary

// Consume alerts off kafka
// https://github.com/vsouza/go-kafka-example

type Alert struct {
    ID  string
    Severity int
    Source string
    Detail string
}

var alerts []Alert

// create a new item
func CreateAlert(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    var alert Alert
    _ = json.NewDecoder(r.Body).Decode(&alert)
    alert.ID = GenerateAlertHash(alert)
    //alerts = append(alerts, alert)
    //json.NewEncoder(w).Encode(alerts)
}

func GenerateAlertHash(alert *Alert)(string){
    return "foo" //todo
}

// Handler for email
// Handler for slack
// Handler for pagerduty

// 
// main function to boot up everything
func main() {
    //router := mux.NewRouter()

    //router.HandleFunc("/alert", CreateAlert).Methods("POST")

    //log.Fatal(http.ListenAndServe(":8000", router))
    kingpin.Parse()
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	brokers := *brokerList
	master, err := sarama.NewConsumer(brokers, config)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := master.Close(); err != nil {
			panic(err)
		}
	}()
	consumer, err := master.ConsumePartition(*topic, 0, sarama.OffsetOldest)
	if err != nil {
		panic(err)
	}
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)
	doneCh := make(chan struct{})
	go func() {
		for {
			select {
			case err := <-consumer.Errors():
				fmt.Println(err)
			case msg := <-consumer.Messages():
				*messageCountStart++
				fmt.Println("Received messages", string(msg.Key), string(msg.Value))
			case <-signals:
				fmt.Println("Interrupt is detected")
				doneCh <- struct{}{}
			}
		}
	}()
	<-doneCh
fmt.Println("Processed", *messageCountStart, "messages")

}