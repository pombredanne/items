

create table app_app(
  id serial primary key,
  name varchar(300) NOT NULL, --应用名称
  bundle_id varchar(300) NOT NULL, --应用ID
  os VARCHAR(10) NOT NULL, --应用OS
  category VARCHAR(10) NOT NULL, --应用分类 默认两级 例如一级 001 二级 001001
  sub_category VARCHAR(10) NOT NULL, --应用分类 默认两级 例如一级 001 二级 001001
  keywords VARCHAR(100) not NULL DEFAULT '',
  store_name VARCHAR (300) not NULL DEFAULT '',
  store_url VARCHAR (300) not NULL DEFAULT '',
  describtion VARCHAR (200) not NULL DEFAULT '', --应用介绍
  category_limit jsonb not NULL  DEFAULT '[]',
  reward jsonb NOT NULL DEFAULT '{}',
  slots  jsonb NOT NULL DEFAULT '{}',
  user_id INTEGER NOT NULL , --用户ID
  zonst_user_id INTEGER NOT NULL DEFAULT 0, --业务
  create_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP ,
  update_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP ,
  status SMALLINT NOT NULL DEFAULT 0, --((-1,审核失败)(0,待审核),(1,已审核))
  UNIQUE(bundle_id,os,user_id)
);

{
"reward":{
"open":0,
"currency_name":"",
"amount":0,
"callback":1
"callback_url":""
}

}



-- APP 奖励机制 是否需要创建默认？
-- create table app_reward(
--   id serial primary key,
--   app_id varchar(30) NOT NULL , --应用ID
--   app_os SMALLINT  NOT NULL, --应用OS
--   currency_name varchar(30) NOT NULL DEFAULT '', --开发者自定义货币名称 不填则不显示单位
--   num INTEGER NOT NULL DEFAULT 0, --奖励个数
--   type SMALLINT NOT NULL DEFAULT 1, --奖励频率 ((1,平均),)
--   callback SMALLINT NOT NULL DEFAULT 1, --奖励回调 ((1,仅客户端回调),(2,客户端和服务器端回调))
--   callback_url VARCHAR(300) NOT NULL DEFAULT '', --回调地址
--   UNIQUE(app_id)
-- );