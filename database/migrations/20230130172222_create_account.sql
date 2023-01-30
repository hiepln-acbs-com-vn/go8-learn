-- +goose Up
-- +goose StatementBegin
create table IF not exists accounts
(
    id bigserial
    constraint accounts_pk
    primary key,
    account_name text not null,
    created_at timestamp with time zone default current_timestamp,
    updated_at timestamp with time zone default current_timestamp,
    deleted_at timestamp with time zone
);


CREATE TABLE IF NOT EXISTS ekycs
(
    id bigserial,
    ekyc_name varchar(255) not null,
    created_at timestamp with time zone default current_timestamp,
    updated_at timestamp with time zone default current_timestamp,
    deleted_at timestamp with time zone,
                                 primary key (id)
);

create table if not exists ekyc_accounts
(
    ekyc_id bigserial not null
    constraint ekyc_accounts_ekycs_id_fk
    references ekycs
    on delete cascade,
    account_id bigserial not null
    constraint ekyc_accounts_account_id_fk
    references accounts
    on delete cascade,
    constraint ekyc_authors_pk
    primary key (ekyc_id, account_id)
    );

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table ekyc_accounts
drop table ekycs
drop table accounts
-- +goose StatementEnd
