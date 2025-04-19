create table book(
    id bigint unsigned not null auto_increment comment '自增id',
    name varchar(255) not null comment '书名',
    price int(11) not null default 0 comment '价格',
    primary key (id)
)engine=innodb default charset=utf8mb4 comment '书籍表';