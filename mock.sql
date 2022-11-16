create database book;

drop database book;

create table book_info (
    id serial not null,
    title varchar(150) not null,
    author_name varchar(100) not null,
    price numeric(18, 2) not null,
    amount int not null DEFAULT 0,
    created_at TIMESTAMP default current_timestamp
);