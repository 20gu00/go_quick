create table community
(
    id             int auto_increment primary key,
    community_id   int unsigned                        not null,
    community_name varchar(128)                        not null,
    introduction   varchar(250)                        not null,
    create_time    timestamp default CURRENT_TIMESTAMP not null,
    update_time    timestamp default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP,
    constraint idx_community_id
        unique (community_id),
    constraint idx_community_name
        unique (community_name)
)collate = utf8mb4_general_ci;

INSERT INTO forum.community (id, community_id, community_name, introduction, create_time, update_time) VALUES (1, 1, 'Go', 'Golang', '2022-12-02 09:19:19', '2022-12-02 10:19:19');
INSERT INTO community (id, community_id, community_name, introduction, create_time, update_time) VALUES (2, 2, 'rust', 'c', '2022-12-02 12:19:19', '2022-12-02 13:19:19');


