-- +goose Up
-- +goose StatementBegin
create table if not exists characteristic
(
    id          bigserial primary key,
    name        text,
    ch_type     text,
    description text
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists characteristic;
-- +goose StatementEnd
