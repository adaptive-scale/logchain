ATTACH TABLE _ UUID '37ed1bf0-5b82-46af-af8c-9a2719972653'
(
    `event_date` Date CODEC(Delta(2), ZSTD(1)),
    `event_time` DateTime CODEC(Delta(4), ZSTD(1)),
    `metric` LowCardinality(String) CODEC(ZSTD(1)),
    `value` Float64 CODEC(ZSTD(3))
)
ENGINE = MergeTree
PARTITION BY toYYYYMM(event_date)
ORDER BY (metric, event_date, event_time)
SETTINGS index_granularity = 8192
