package service

import (
	"context"
	"encoding/json"
	"github.com/adaptive-scale/logchain/internal"
	"github.com/adaptive-scale/logchain/pkg/logchain"
	"gorm.io/driver/clickhouse"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"time"
)

type LogChainServiceImpl struct {
	logchain.UnimplementedLogChainServer
	*gorm.DB
}

func New(config internal.Configuration) *LogChainServiceImpl {

	var db *gorm.DB
	var err error
	switch config.DatabaseType {
	case "clickhouse", "":
		{
			db, err = gorm.Open(clickhouse.Open(config.DSN), &gorm.Config{})
			if err != nil {
				log.Fatal("failed to connect database")
			}
			err := db.AutoMigrate(&internal.LogStore{}, &internal.MetricStore{})
			if err != nil {
				log.Fatal(err)
			}
			break
		}
	case "postgres":
		{
			db, err = gorm.Open(postgres.Open(config.DSN), &gorm.Config{})
			if err != nil {
				log.Fatal("failed to connect database")
			}
			break
		}
	default:
		log.Fatal("database type is not supported yet")
	}

	return &LogChainServiceImpl{DB: db}
}

func (l *LogChainServiceImpl) Log(ctx context.Context, req *logchain.LogRequest) (*logchain.Response, error) {

	var logline internal.LogStore

	logline.Message = req.LogLine
	logline.LogLevel = req.LogLevel
	logline.Timestamp = time.UnixMicro(req.Timestamp)
	logline.AppName = req.AppName

	var val = map[string]string{}
	for _, v := range req.Labels {
		val[v.Name] = v.Value
	}
	d, _ := json.Marshal(val)

	logline.Labels = d

	err := l.WithContext(ctx).Create(&logline)
	if err != nil {
		log.Println(err)
	}

	return &logchain.Response{Success: true}, nil
}

func (l *LogChainServiceImpl) Metric(ctx context.Context, req *logchain.MetricRequest) (*logchain.Response, error) {

	var metric internal.MetricStore

	metric.MetricGroup = req.MetricGroup
	metric.MetricName = req.MetricName

	if req.Timestamp == 0 {
		metric.Timestamp = time.Now()
	} else {
		metric.Timestamp = time.UnixMicro(req.Timestamp)
	}

	metric.Value = req.MetricValue

	var val = map[string]string{}
	for _, v := range req.Labels {
		val[v.Name] = v.Value
	}
	d, _ := json.Marshal(val)
	metric.Labels = d

	err := l.WithContext(ctx).Create(&metric)
	if err != nil {
		log.Println(err)
	}

	return &logchain.Response{Success: true}, nil
}
