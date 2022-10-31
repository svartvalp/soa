-- +goose Up
-- +goose StatementBegin
create table if not exists product
(
    id          bigserial primary key,
    name        text,
    description  text,
    category_id bigint references category (id),
    brand       text,
    price       bigint,
    image       text
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists product;
-- +goose StatementEnd
