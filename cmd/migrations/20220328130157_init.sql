-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS logs
(
    service_name String,
    log_level    String,
    span_id      UUID,
    tstamp       DateTime,
    src          String,
    message      String
) ENGINE Log;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS logs;
-- +goose StatementEnd
