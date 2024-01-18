# Logchain

Logchain is a grpc adapter for clickhouse and postgres for persisting logs.

##### Use cases

- When applications are deployment on-premise or on customer side. Logchain provides visibility into Errors for application support without additional infrastructure
- For small/medium applications that do not need huge logging infrastructure but still need visibility.

Install logchain using the following command:
```
go install github.com/adaptive-scale/logchain@1.0.3
```

#### Setup Logchain

```
docker stack deploy logchain --compose-file ./dev/docker-compose.yaml
```

This will start the logchain server on port 9090, clickhouse on port 8123 and grafana at port 3000.

For golang client, import the logchain client as follows:

```go
package main

import (
	"github.com/adaptive-scale/logchain/pkg/logchain"
	"github.com/adaptive-scale/logchain/pkg/logchainhook"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"github.com/sirupsen/logrus"
)

func main()  {
	conn, err := grpc.Dial("localhost:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logrus.Fatalf("did not connect: %v", err)
	}
	c := logchain.NewLogChainClient(conn)
	hook := logchainhook.NewLogChainHook("app=name", c, logrus.ErrorLevel)
	logrus.AddHook(hook)
	
	logrus.WithField("key", "value").Error("error message")
}
```

#### Setup Logging Dashboards in Grafana

You can select the clickhouse as datasource and then use the following query to get the logs:

```
SELECT app_name, date(timestamp) dt,  message, labels, count(message) cnt FROM logchain.log_stores  group by app_name, dt, message, labels order by dt desc, cnt desc LIMIT 100
```