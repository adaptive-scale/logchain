# Logchain

Logchain is a grpc adapter for clickhouse for persisting logs and metrics.

#### Start Clickhouse server
```
docker run --name some-clickhouse-server --ulimit nofile=262144:262144 -e CLICKHOUSE_DB=logchain -e CLICKHOUSE_USER=logchain -e CLICKHOUSE_DEFAULT_ACCESS_MANAGEMENT=1 -e CLICKHOUSE_PASSWORD=L0gCh@in -p 9000:9000 -p 18123:8123 clickhouse/clickhouse-server
```