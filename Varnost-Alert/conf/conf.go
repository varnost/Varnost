package conf

type KafkaConfiguration struct {
	Brokerlist []string
	Topic string
	Partition int
	offsetType int
	messageCountStart int
}

type AlertConfiguration struct {
	FatalEmail string
	FatalSlackChannel string
	WarnEmail string
	WarnSlackChannel string
	InfoEmail string
	InfoSlackChannel string
}

type Configuration struct {
	kafka KafkaConfiguration
	alert AlertConfiguration
}