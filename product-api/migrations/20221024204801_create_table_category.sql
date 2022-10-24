-- +goose Up
-- +goose StatementBegin
create table if not exists category
(
    id          bigint primary key,
    name        text,
    description text,
    parent_id   bigint,
    level       int
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists category;
-- +goose StatementEnd
