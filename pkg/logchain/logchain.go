package logchain

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/debarshibasak/logchain/internal"
	"gorm.io/driver/clickhouse"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"time"
)

type LogChainServiceImpl struct {
	UnimplementedLogChainServer
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
			err := db.AutoMigrate(&internal.LogStore{})
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

func (l *LogChainServiceImpl) Log(ctx context.Context, req *LogRequest) (*LogResponse, error) {

	var logline internal.LogStore

	fmt.Println(req.Timestamp)

	logline.LogLine = req.LogLine
	logline.LogLevel = req.LogLevel
	logline.Timestamp = time.UnixMicro(req.Timestamp)
	logline.AppName = req.AppName

	d, _ := json.Marshal(req.Labels)

	logline.Labels = d

	err := l.WithContext(ctx).Create(&logline)
	if err != nil {
		log.Println(err)
	}

	return &LogResponse{Status: true}, nil
}
