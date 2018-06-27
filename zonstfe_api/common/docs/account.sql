-- 账户信息
create table account_account(
 id serial primary key,
 email varchar(320) NOT NULL,
 user_id INTEGER NOT NULL ,--用户ID,
 user_type VARCHAR (20) NOT NULL DEFAULT '',
 user_role SMALLINT NOT NULL ,
 qq VARCHAR (18) NOT NULL DEFAULT '',
 phone VARCHAR (20) NOT NULL DEFAULT '',
 real_name VARCHAR (20) NOT NULL DEFAULT '',
 company_name VARCHAR (300) NOT NULL DEFAULT '',
 qualification jsonb  NOT NULL DEFAULT '{}',
--  business_license VARCHAR (300) NOT NULL DEFAULT '',--营业执照
--  organization_code VARCHAR (300) NOT NULL DEFAULT '',--组织机构代码
--  partner_name VARCHAR (20) NOT NULL DEFAULT '',
--  partner_idcard VARCHAR (100) NOT NULL DEFAULT '',
--  contact_name VARCHAR (20) NOT NULL DEFAULT '',
--  contact_address VARCHAR (300) NOT NULL DEFAULT '',
 app_key VARCHAR (300) NOT NULL DEFAULT '',
 deal_type VARCHAR (20) NOT NULL DEFAULT 'share', -- share 分成模式  bidding 出价模式
 deal_scale FLOAT NOT NULL DEFAULT 0.3,
--  company_addr varchar(100) NOT NULL DEFAULT '',--地址
--  tax_no  VARCHAR(30) NOT NULL DEFAULT '',--税号
 statement_type SMALLINT NOT NULL DEFAULT 0, -- 1对公 0对私
 need_invoice SMALLINT NOT NULL DEFAULT 0 ,--是否需要开票 1是 0否
 zonst_user_id INTEGER NOT NULL DEFAULT 0,--业务ID
 create_date TIMESTAMP  NOT NULL DEFAULT CURRENT_TIMESTAMP, --创建时间
 status SMALLINT not NULL DEFAULT 0, -- -1 审核失败 0 待审核 1已审核
  UNIQUE (email,user_role)
);
-- 支付信息
create table account_payment(
  id serial PRIMARY KEY,
  user_id integer NOT NULL,
  user_role SMALLINT NOT NULL,
  order_no VARCHAR (30) NOT NULL DEFAULT '',
  order_date timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP, -- 处理时间
  order_money numeric(9, 2) NOT NULL DEFAULT 0,
  balance numeric(9, 2) NOT NULL DEFAULT 0,
  apply_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP , -- 申请时间
  status smallint NOT NULL DEFAULT 0 -- -1 审核失败 0 待审核 1已审核
);
-- 消息
create table account_message(
id serial primary key,
user_id integer NOT NULL,
user_role SMALLINT NOT NULL,
title VARCHAR(100) NOT NULL DEFAULT '',
content TEXT NOT NULL DEFAULT '',
group_name VARCHAR(100) NOT NULL DEFAULT '',
status SMALLINT NOT NULL DEFAULT 0, -- 0
create_date date DEFAULT CURRENT_DATE,
create_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


-- 账户余额信息
create table account_balance(
   id serial primary key,
   user_id INTEGER NOT NULL ,--用户ID,
   app_key text NOT NULL , --用户ID
   user_role SMALLINT NOT NULL ,
   balance NUMERIC (12,4) NOT NULL DEFAULT 0, --账户余额
  UNIQUE (user_id,user_role)
);

-- 邮寄信息
create table account_deliver(
 id serial primary key,
 user_id INTEGER NOT NULL ,--用户ID,
 user_role SMALLINT NOT NULL,
 address text NOT NULL DEFAULT '',
 receiver VARCHAR (20) NOT NULL DEFAULT '',
 telephone VARCHAR (20) NOT NULL DEFAULT '',
 email varchar(320) NOT NULL DEFAULT '',
  UNIQUE (user_id,user_role)
);

-- 税务信息
create table account_tax(
  id serial primary key,
  user_id INTEGER NOT NULL ,--用户ID,
  user_role SMALLINT NOT NULL,
  tax_no VARCHAR (20) NOT NULL DEFAULT '', --15-20
  tax_voucher VARCHAR (300) NOT NULL DEFAULT '',
  company_name VARCHAR (300) NOT NULL DEFAULT '',
  address text NOT NULL DEFAULT '',
  telephone VARCHAR (20) NOT NULL DEFAULT '',
  bank_name VARCHAR (100) NOT NULL DEFAULT '',
  bank_no  VARCHAR (19) NOT NULL DEFAULT '',
  UNIQUE (user_id,user_role)
);

-- 财务信息
create table account_finance(
 id serial primary key,
 user_id INTEGER NOT NULL ,--用户ID,
 user_role SMALLINT NOT NULL,
 account_name VARCHAR (20) NOT NULL DEFAULT '',
 bank_name VARCHAR (100) NOT NULL DEFAULT '',
 bank_no  VARCHAR (19) NOT NULL DEFAULT '',
 bank_province VARCHAR (100) NOT NULL DEFAULT '',
 bank_city VARCHAR (100) NOT NULL DEFAULT '',
 sub_branch VARCHAR (300) NOT NULL DEFAULT '',
  UNIQUE (user_id,user_role)
);


-- 充值信息
create table account_recharge(
id serial primary key,
user_id INTEGER NOT NULL ,--用户ID
user_role SMALLINT NOT NULL,
zonst_user_id INTEGER NOT NULL DEFAULT 0,--业务ID
order_no VARCHAR (30) NOT  NULL DEFAULT '',--订单号
order_money NUMERIC (12,4),--金额
order_type SMALLINT NOT NULL , -- 1 银行, 2 支付宝, 3 微信
order_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,--付款时间
account_no VARCHAR(20) NOT NULL, --用户账号
account_name VARCHAR(20) NOT NULL ,--用户名
description text NOT NULL DEFAULT '',--注释
update_date  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ,--修改时间
create_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,--创建时间
status SMALLINT NOT NULL DEFAULT 0--状态 ((0,待处理),(1,已处理))
);