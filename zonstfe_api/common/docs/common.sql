--用户表  公共表
create table user_user(
    id serial primary key,
    email varchar(320) NOT NULL, --用户名邮箱
    password text, --用户密码
    role SMALLINT NOT NULL DEFAULT 1, --1.开发者 2开发者后台,3 广告商4广告商后台
    type VARCHAR (20) NOT NULL DEFAULT '', -- person company
    real_name VARCHAR (20) NOT NULL DEFAULT '',
    company_name VARCHAR (100) NOT NULL DEFAULT '',
    -- type SMALLINT NOT NULL DEFAULT 1, -- ((1,个人),(2,公司))
    status SMALLINT NOT null default 1, --0 账号异常 1 待审核  2 已审核
    ip_address VARCHAR(30) NOT NULL DEFAULT '',
    reg_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(email,role)  --
);

select setval('user_user_id_seq',10000,false);

create table log_actions(
id serial PRIMARY key,
user_id INTEGER NOT NULL, -- 操作人
action_user_id  INTEGER NOT NULL, -- 操作对象
action_type varchar(30) NOT NULL DEFAULT '',
action_module  varchar(30) NOT NULL DEFAULT '',
action_id INTEGER NOT NULL DEFAULT 0,
action_sql TEXT NOT NULL DEFAULT '',
request_path TEXT NOT NULL DEFAULT '',
request_method VARCHAR(10) NOT NULL DEFAULT '',
request_data TEXT NOT NULL DEFAULT '',
platform_id SMALLINT NOT NULL,
ip_address VARCHAR(30) NOT NULL DEFAULT '',
create_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

create table app_category(
id serial PRIMARY key,
name VARCHAR(200) NOT NULL,
code VARCHAR(100) NOT NULL,
level_num SMALLINT NOT NULL,
parent_code VARCHAR(100) NOT NULL
);
create table log_event(
id serial PRIMARY key,
name VARCHAR(100) NOT NULL DEFAULT '',
event_id VARCHAR(300) NOT NULL DEFAULT '',
event_obj TEXT NOT NULL DEFAULT '',
error_msg text not NULL DEFAULT '',
start_time INTEGER NOT NULL,
end_time INTEGER NOT NULL DEFAULT 0,
status SMALLINT NOT NULL DEFAULT 0,  --0 发起 1 完成 -1 错误
create_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

create table spider_app(
id serial primary key,
NAME TEXT NOT NULL,
app_id VARCHAR(100) not NULL DEFAULT '',
os VARCHAR(30) not NULL DEFAULT '',
genre VARCHAR(100) not NULL DEFAULT '',
sub_genre VARCHAR(100) not NULL DEFAULT '',
hot SMALLINT NOT NULL DEFAULT 0,
info jsonb not null DEFAULT '{}',
create_date  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
update_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
create table spider_identifier(
id serial primary key,
device VARCHAR (300) not NULL  DEFAULT '',
device_detail VARCHAR (300) not NULL  DEFAULT '',
identifier VARCHAR (300) not NULL  DEFAULT ''
);



create table log_spider
(
id serial PRIMARY key,
name VARCHAR(300) NOT NULL DEFAULT '',
signals VARCHAR(300) NOT NULL DEFAULT '',
info jsonb NOT NULL DEFAULT '{}',
report_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);