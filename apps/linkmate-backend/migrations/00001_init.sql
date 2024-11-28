-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users
(
    id          SERIAL PRIMARY KEY,
    email       TEXT UNIQUE NOT NULL,
    name        TEXT,
    password    TEXT        NOT NULL,
    create_time TIMESTAMP   NOT NULL DEFAULT NOW(),
    update_time TIMESTAMP   NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS links
(
    id          SERIAL PRIMARY KEY,
    key         TEXT UNIQUE,
    url         TEXT      NOT NULL,
    user_id     INT       NOT NULL,
    create_time TIMESTAMP NOT NULL DEFAULT NOW(),
    update_time TIMESTAMP NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS tags
(
    id          SERIAL PRIMARY KEY,
    name        TEXT      NOT NULL,
    color       TEXT      NOT NULL,
    user_id     INT       NOT NULL,
    create_time TIMESTAMP NOT NULL DEFAULT NOW(),
    update_time TIMESTAMP NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);

-- TODO: Rename to tags_settings
CREATE TABLE IF NOT EXISTS links_tags
(
    link_id INT REFERENCES links (id) ON UPDATE CASCADE ON DELETE CASCADE,
    tag_id  INT REFERENCES tags (id) ON UPDATE CASCADE ON DELETE CASCADE,
    CONSTRAINT pk_links_tags PRIMARY KEY (link_id, tag_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS links_tags;
DROP TABLE IF EXISTS links;
DROP TABLE IF EXISTS tags;
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
