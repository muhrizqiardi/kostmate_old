-- migrate:up
create extension if not exists "uuid-ossp";

create table if not exists public.rooms (
    id uuid default uuid_generate_v4() primary key, 
    unique_name text not null unique,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp,
);

-- migrate:down
drop table public.rooms;

