ATTACH TABLE _ UUID '5cefca4d-bb72-48ee-9bef-fd9c1f10128e'
(
    `metric_group` String,
    `metric_name` String,
    `value` Float64,
    `timestamp` DateTime64(3),
    `labels` String
)
ENGINE = MergeTree
ORDER BY tuple()
SETTINGS index_granularity = 8192
