ATTACH TABLE _ UUID '64bd3b14-7e73-4091-963a-84d6c080fcdc'
(
    `app_name` String,
    `log_level` String,
    `timestamp` DateTime64(3),
    `message` String,
    `labels` String
)
ENGINE = MergeTree
ORDER BY tuple()
SETTINGS index_granularity = 8192
