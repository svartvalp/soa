-- +goose Up
-- +goose StatementBegin
create sequence category_id_seq start 1 increment by 1;
alter table category alter column id set default nextval('category_id_seq');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
alter table category alter column id set default null;
drop sequence category_id_seq;
-- +goose StatementEnd
