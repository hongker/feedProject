create database feed charset=utf8;

drop table if exists temp;
create table temp (
        id int unsigned not null primary key auto_increment comment '主键',
        created_at  timestamp    default CURRENT_TIMESTAMP not null comment '创建时间',
        updated_at  timestamp    default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP comment '更新时间',
        key `idx_updated_at`(`updated_at`)
)charset=utf8mb4;

drop table if exists users;
create table users (
        id int unsigned not null primary key auto_increment comment '主键',
        username varchar(100) not null default '' comment '用户名',
        role_type tinyint not null default 1 comment '用户角色(1:普通用户,2:大V)',
        status tinyint unsigned not null default 1 comment '状态(1:活跃,2:不活跃的)',
        created_at  timestamp    default CURRENT_TIMESTAMP not null comment '创建时间',
        updated_at  timestamp    default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP comment '更新时间',
        key `idx_updated_at`(`updated_at`)
)charset=utf8mb4 comment='用户信息';
insert into users(username, role_type, status)
values('test01',1,1),('test02',1,2),('test03',2,1)

drop table if exists relations;
create table relations (
        id int unsigned not null primary key auto_increment comment '主键',
        user_id int unsigned not null default 0 comment '用户ID',
        target_id int unsigned not null default 0 comment '关注目标ID',
        status tinyint unsigned not null default 1 comment '状态(1:正常,2:删除)',
        created_at  timestamp    default CURRENT_TIMESTAMP not null comment '创建时间',
        updated_at  timestamp    default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP comment '更新时间',
        key `idx_user_id`(`user_id`),
        key `idx_status`(`status`),
        key `idx_updated_at`(`updated_at`)
)charset=utf8mb4 comment='关注信息';

drop table if exists feed;
create table feed (
        id int unsigned not null primary key auto_increment comment '主键',
        user_id int unsigned not null default 0 comment '所属用户UID',
        creator_id int unsigned not null default 0 comment '创建者UID',
        feed_id  bigint unsigned not null default 0 comment 'feed唯一ID',
        content text not null  comment '内容',
        status tinyint unsigned not null default 1 comment '状态(1:正常,2:停用)',
        created_at  timestamp    default CURRENT_TIMESTAMP not null comment '创建时间',
        updated_at  timestamp    default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP comment '更新时间',
        key `idx_creator_id`(`creator_id`),
        key `idx_status`(`status`),
        key `idx_updated_at`(`updated_at`)
)charset=utf8mb4;

drop table if exists temp;
create table temp (
        id int unsigned not null primary key auto_increment comment '主键',
        created_at  timestamp    default CURRENT_TIMESTAMP not null comment '创建时间',
        updated_at  timestamp    default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP comment '更新时间',
        key `idx_updated_at`(`updated_at`)
)charset=utf8mb4;