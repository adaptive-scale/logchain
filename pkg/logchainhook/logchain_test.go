package logchainhook_test

import (
	"context"
	"github.com/adaptive-scale/logchain/pkg/logchain"
	"github.com/adaptive-scale/logchain/pkg/logchainhook"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"testing"
)

func TestLogEvents(t *testing.T) {

	conn, err := grpc.Dial("localhost:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	c := logchain.NewLogChainClient(conn)

	hook := logchainhook.NewLogChainHook("testapp", c, logrus.ErrorLevel)
	logrus.AddHook(hook)
	logrus.WithField("test", "Test").Error("test123")
}

func TestEventMetrics(t *testing.T) {
	conn, err := grpc.Dial("localhost:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	c := logchain.NewLogChainClient(conn)

	resp, err := c.Metric(context.Background(), &logchain.MetricRequest{
		MetricGroup: "usage",
		MetricName:  "metric_session",
		MetricValue: 20,
		Labels: []*logchain.Labels{{
			Name:  "workspace_url",
			Value: []byte("data"),
		}},
	})
	if !resp.Success {
		log.Println(resp.Error)
		return
	}

	if err != nil {
		log.Println(err)
		return
	}
}
