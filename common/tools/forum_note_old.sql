create table note
(
    id           bigint auto_increment primary key,
    post_id      bigint                              not null comment '帖子id',
    title        varchar(128)                        not null comment '标题',
    content      varchar(8192)                       not null comment '内容',
    author_id    bigint                              not null comment '作者的用户id',
    community_id bigint                              not null comment '所属社区',
    status       tinyint   default 1                 not null comment '帖子状态',
    create_time  timestamp default CURRENT_TIMESTAMP null comment '创建时间',
    update_time  timestamp default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP comment '更新时间',
    constraint idx_post_id
        unique (post_id)
)collate = utf8mb4_general_ci;

create index idx_author_id
    on note (author_id);

create index idx_community_id
    on note (community_id);

INSERT INTO note (id, post_id, title, content, author_id, community_id, status, create_time, update_time) VALUES (1, 14283784123846656, 'golang', 'go', 28018727488323585, 1, 1, '2022-12-01 09:09:10', '2022-12-01 10:09:10');
INSERT INTO note (id, post_id, title, content, author_id, community_id, status, create_time, update_time) VALUES (2, 14373128436191232, 'CSGO', 'c', 28018727488323585, 2, 1, '2022-12-01 11:09:10', '2022-12-01 12:09:10');
