-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS link_analytics_events
(
    id           uuid               DEFAULT gen_random_uuid() PRIMARY KEY,
    link_id      INT       NOT NULL,
    create_time  TIMESTAMP NOT NULL DEFAULT NOW(),
    user_agent   TEXT               DEFAULT NULL,
    browser_name TEXT               DEFAULT NULL,
    device_type  TEXT               DEFAULT NULL,
    os_name      TEXT               DEFAULT NULL,
    source       TEXT               DEFAULT NULL,
    ip_address   TEXT               DEFAULT NULL,
    geolocation  jsonb              DEFAULT NULL,


    CONSTRAINT fk_link FOREIGN KEY (link_id) REFERENCES links (id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS link_analytics_events;
-- +goose StatementEnd
