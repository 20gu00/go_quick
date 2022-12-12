create table user
(
    id          bigint auto_increment primary key,
    user_id     bigint                              not null,
    username    varchar(64)                         not null,
    password    varchar(64)                         not null,
    email       varchar(64)                         null,
    gender      tinyint   default 0                 not null,
    create_time timestamp default CURRENT_TIMESTAMP null,
    update_time timestamp default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP,
    constraint idx_user_id
        unique (user_id),
    constraint
        unique (username)
)
    collate = utf8mb4_general_ci;

INSERT INTO forum.user (id, user_id, username, password, email, gender, create_time, update_time) VALUES (1, 28018727488323585, 'aaa', '313233343536639a9119599647d841b1bef6ce5ea293', null, 0, '2022-12-2 09:09:09', '2022-12-2 09:09:09');
INSERT INTO forum.user (id, user_id, username, password, email, gender, create_time, update_time) VALUES (2, 4183532125556736, 'cjq', '313233639a9119599647d841b1bef6ce5ea293', null, 0, '2022-12-2 19:09:09', '2022-12-2 19:09:09');
