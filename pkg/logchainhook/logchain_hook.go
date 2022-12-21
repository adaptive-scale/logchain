package logchainhook

import (
	"context"
	"errors"
	"fmt"
	"github.com/debarshibasak/logchain/pkg/logchain"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

type LogChainHook struct {
	logChainClient    logchain.LogChainClient
	acceptedLogLevels []log.Level
	appName           string
}

func NewLogChainHook(appName, serverLocation string, allowedLevels ...log.Level) *LogChainHook {

	conn, err := grpc.Dial(serverLocation, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	c := logchain.NewLogChainClient(conn)

	return &LogChainHook{
		logChainClient:    c,
		acceptedLogLevels: allowedLevels,
		appName:           appName,
	}
}

func (hook *LogChainHook) Fire(entry *log.Entry) error {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var labels []string
	for key, values := range entry.Data {
		labels = append(labels, fmt.Sprintf("%v:%v", key, values))
	}
	resp, err := hook.logChainClient.Log(ctx, &logchain.LogRequest{
		AppName:   hook.appName,
		LogLine:   []byte(entry.Message),
		Labels:    labels,
		LogLevel:  entry.Level.String(),
		Timestamp: entry.Time.UnixMicro(),
	})
	if !resp.Status {
		return errors.New("response was invalid from the logchain server")
	}
	return err
}

// Levels define on which log levels this hook would trigger
func (hook *LogChainHook) Levels() []log.Level {
	return hook.acceptedLogLevels
}
