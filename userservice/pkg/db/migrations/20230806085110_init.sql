-- migrate:up
create extension if not exists "uuid-ossp";

create type user_roles as enum ('USER', 'ADMIN');

create table if not exists public.users (
    id uuid default uuid_generate_v4() primary key, 
    email text not null unique,
    full_name text not null,
    role user_roles not null,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp,
);

-- migrate:down
drop table public.users;
