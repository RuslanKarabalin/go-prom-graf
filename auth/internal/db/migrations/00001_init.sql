-- +goose Up
create type "user_role" as enum (
    'user'
    , 'admin'
);

create table "users" (
    id uuid primary key default gen_random_uuid()
    , email varchar(255) not null unique
    , password_hash varchar(255) not null
    , enabled boolean not null default true
);

create table "user_roles" (
    user_id uuid references users(id) on delete cascade
    , role user_role not null
    , primary key (user_id, role)
);

-- +goose Down
drop table if exists "user_roles";
drop table if exists "users";

drop type if exists "user_role";

