-- +goose Up
-- +goose StatementBegin
create table if not exists product_characteristic
(
    id                bigserial primary key,
    product_id        bigint references product (id) unique,
    characteristic_id bigint references characteristic (id) unique,
    value             text,
    ch_type           text
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists product_characteristic;
-- +goose StatementEnd
