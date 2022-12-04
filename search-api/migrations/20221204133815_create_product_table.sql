-- +goose Up
-- +goose StatementBegin
create table if not exists product
(
    id              bigserial primary key,
    category_id     bigint,
    price           bigint,
    name            text,
    description     text,
    brand           text,
    image           text,
    characteristics jsonb,
    categorys       jsonb
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists product;
-- +goose StatementEnd
