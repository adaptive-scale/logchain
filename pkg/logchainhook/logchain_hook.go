package logchainhook

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/adaptive-scale/logchain/pkg/logchain"
	log "github.com/sirupsen/logrus"
	"time"
)

type LogChainHook struct {
	logChainClient    logchain.LogChainClient
	acceptedLogLevels []log.Level
	appName           string
}

func NewLogChainHook(appName string, c logchain.LogChainClient, allowedLevels ...log.Level) *LogChainHook {
	return &LogChainHook{
		logChainClient:    c,
		acceptedLogLevels: allowedLevels,
		appName:           appName,
	}
}

func (hook *LogChainHook) Fire(entry *log.Entry) error {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var labels []*logchain.Labels
	for key, values := range entry.Data {
		d, _ := json.Marshal(values)
		labels = append(labels, &logchain.Labels{
			Name: key, Value: string(d),
		})
	}
	resp, err := hook.logChainClient.Log(ctx, &logchain.LogRequest{
		AppName:   hook.appName,
		LogLine:   []byte(entry.Message),
		Labels:    labels,
		LogLevel:  entry.Level.String(),
		Timestamp: entry.Time.UnixMicro(),
	})

	if resp == nil {
		return errors.New("response was empty")
	}
	if !resp.Success {
		return errors.New("response was invalid from the logchain server")
	}
	return err
}

// Levels define on which log levels this hook would trigger
func (hook *LogChainHook) Levels() []log.Level {
	return hook.acceptedLogLevels
}
