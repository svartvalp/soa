-- +goose Up
-- +goose StatementBegin
create table if not exists product_characteristic
(
    id                bigserial primary key,
    product_id        bigint references product (id),
    characteristic_id bigint references characteristic (id),
    value             text,
    ch_type           text
);
-- +goose StatementEnd

-- +goose NO TRANSACTION
create unique index concurrently unique_baking_plan_idx
    on product_characteristic (product_id, characteristic_id);

-- +goose Down
-- +goose StatementBegin
drop table if exists product_characteristic;
-- +goose StatementEnd
