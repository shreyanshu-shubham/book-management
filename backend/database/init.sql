create table authors (
    id integer primary key,
    first_name text not null,
    last_name text
);

create table books (
    isbn bigint primary key,
    title text not null,
    author_id int not null,
    is_deleted boolean not null,
    constraint fk_author foreign key (author_id) references authors (id)
);