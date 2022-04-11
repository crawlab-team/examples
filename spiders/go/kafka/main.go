package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/crawlab-team/go-trace"
	"github.com/segmentio/kafka-go"
	"time"
)

var topic = "crawlab"

func main() {
	conn, err := kafka.Dial("tcp", "localhost:9092")
	if err != nil {
		panic(err)
	}

	if err := conn.CreateTopics(kafka.TopicConfig{
		Topic:             topic,
		NumPartitions:     -1,
		ReplicationFactor: -1,
	}); err != nil {
		panic(err)
	}
	_ = conn.Close()

	connLeader, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, 0)
	if err != nil {
		panic(err)
	}

	for {
		m, err := connLeader.ReadMessage(1024 * 1024)
		if err != nil {
			trace.PrintError(err)
			time.Sleep(1 * time.Second)
			continue
		}
		var d map[string]interface{}
		if err := json.Unmarshal(m.Value, &d); err != nil {
			trace.PrintError(err)
			continue
		}
		fmt.Println(d)
	}
}
