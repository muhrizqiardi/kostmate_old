-- migrate:up
create extension if not exists "uuid-ossp";

create type user_roles as enum ('USER', 'ADMIN');
create type payment_status as enum ('CANCELLED', 'PENDING', 'SUCCESS');
create type booking_status as enum ('ONGOING', 'ENDED');

create table if not exists public.users (
    id uuid default uuid_generate_v4() primary key, 
    email text not null unique,
    full_name text not null,
    role user_roles not null,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp,
);

create table if not exists public.kos (
    id uuid default uuid_generate_v4() primary key, 
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp,
);

create table if not exists public.rooms (
    id uuid default uuid_generate_v4() primary key, 
    unique_name text not null unique,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp,
);

create table if not exists public.payments (
    id uuid default uuid_generate_v4() primary key, 
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp
)

create table if not exists public.bookings (
    id uuid default uuid_generate_v4() primary key, 
    status booking_status not null,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp
)

-- migrate:down
drop table public.users;
drop table public.kos;
drop table public.rooms;
drop table public.payments;

