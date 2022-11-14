create table book_style (
    id int REFERENCES book_info(id),
    color varchar(50) not null 
);