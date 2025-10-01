create table books (
    isbn bigint primary key,
    title text not null,
    author text not null,
    is_deleted boolean not null
);