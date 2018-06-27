--活动报表
create table report_base(
id serial primary key,
user_id INTEGER NOT NULL , --用户ID
campaign_id INTEGER NOT NULL,--活动ID
vendor_id INTEGER NOT NULL,--投放平台
ad_id INTEGER NOT NULL , --广告id
-- ad_direction varchar(30) NOT NULL , --横竖屏
report_date DATE NOT NULL DEFAULT CURRENT_DATE, --创建时间
hour SMALLINT NOT NULL,--时
win INTEGER NOT NULL DEFAULT 0, -- 胜出
imp INTEGER NOT NULL DEFAULT 0, -- 展示
eimp INTEGER NOT NULL DEFAULT 0, -- 图片展示
clk INTEGER NOT NULL DEFAULT 0, -- 点击
-- valid_clk INTEGER NOT NULL DEFAULT 0, -- 有效点击
-- valid_show INTEGER NOT NULL DEFAULT 0, -- 有效展示
cost NUMERIC (12,4),--消耗金额
UNIQUE (user_id,vendor_id,campaign_id,ad_id,hour,report_date)
);
--活动区域报表
create table report_geo(
id serial primary key,
user_id INTEGER NOT NULL , --用户ID
campaign_id INTEGER NOT NULL,--活动ID
vendor_id INTEGER NOT NULL,--投放平台
ad_id INTEGER NOT NULL , --广告id
report_date DATE NOT NULL DEFAULT CURRENT_DATE,--创建时间
country_code VARCHAR(30) NOT NULL DEFAULT '',
province_code VARCHAR(30) NOT NULL DEFAULT '',
city_code VARCHAR(30) NOT NULL DEFAULT '',
win INTEGER NOT NULL DEFAULT 0, -- 胜出
imp INTEGER NOT NULL DEFAULT 0, -- 展示
eimp INTEGER NOT NULL DEFAULT 0, -- 图片展示
clk INTEGER NOT NULL DEFAULT 0, -- 点击
-- valid_clk INTEGER NOT NULL DEFAULT 0, -- 有效点击
-- valid_show INTEGER NOT NULL DEFAULT 0, -- 有效展示
cost NUMERIC (12,4),--消耗金额
UNIQUE (user_id,vendor_id,campaign_id,ad_id,country_code,province_code,city_code,report_date)

);
--活动app报表
create table report_app(
id serial primary key,
user_id INTEGER NOT NULL , --用户ID
campaign_id INTEGER NOT NULL,--活动ID
ad_id INTEGER NOT NULL , --广告id
vendor_id INTEGER NOT NULL,--投放平台
-- hour SMALLINT NOT NULL,--时
bundle_id varchar(300) NOT NULL , --应用ID
os varchar(100) NOT NULL,
report_date DATE NOT NULL DEFAULT CURRENT_DATE,--创建时间
win INTEGER NOT NULL DEFAULT 0, -- 胜出
imp INTEGER NOT NULL DEFAULT 0, -- 展示
eimp INTEGER NOT NULL DEFAULT 0, -- 图片展示
clk INTEGER NOT NULL DEFAULT 0, -- 点击
-- valid_clk INTEGER NOT NULL DEFAULT 0, -- 有效点击
-- valid_imp INTEGER NOT NULL DEFAULT 0, -- 有效展示
cost NUMERIC (12,4),--消耗金额
UNIQUE (user_id,vendor_id,campaign_id,os,ad_id,bundle_id,report_date)
);

create table report_dev_profit(
  id serial primary key,
  app_key text NOT NULL, --用户ID
  bidding NUMERIC (12,4) NOT NULL DEFAULT 0,
  share NUMERIC (12,4) NOT NULL DEFAULT 0,
  report_date date not NULL,
  UNIQUE (app_key,report_date)
);


--活动app报表
create table report_app_slot(
id serial primary key,
app_key text NOT NULL , --用户ID
bundle_id varchar(300) NOT NULL, --应用ID
-- campaign_id INTEGER NOT NULL,--活动ID
slot_id INTEGER NOT NULL , --广告ID
os varchar(100) NOT NULL,
-- hour SMALLINT NOT NULL,--时
report_date DATE NOT NULL DEFAULT CURRENT_DATE,--创建时间
imp INTEGER NOT NULL DEFAULT 0, -- 展示
clk INTEGER NOT NULL DEFAULT 0, -- 点击
-- valid_clk INTEGER NOT NULL DEFAULT 0, -- 有效点击
-- valid_imp INTEGER NOT NULL DEFAULT 0, -- 有效展示
cost NUMERIC (12,4),--消耗金额
UNIQUE (os,app_key,bundle_id,report_date,slot_id)
);
INSERT INTO  report_app_slot(app_key,bundle_id,slot_id,os,imp,clk,cost) VALUES ('zonst','com.xqw.ncmj','1000','android',1000,100,10)
create table report_app_reward(
id serial primary key,
app_key text NOT NULL , --用户ID
-- campaign_id INTEGER NOT NULL,
-- clk INTEGER NOT NULL DEFAULT 0, -- 点击
imp INTEGER NOT NULL DEFAULT 0, -- 展示
os varchar(100) NOT NULL,
-- valid_clk INTEGER NOT NULL DEFAULT 0, -- 有效点击
-- valid_show INTEGER NOT NULL DEFAULT 0, -- 有效展示
bundle_id varchar(300) NOT NULL,--app_id
amount INTEGER NOT NULL DEFAULT 0, -- 每次奖励总单位数*奖励次数
reward  INTEGER NOT NULL DEFAULT 0, -- 奖励次数
uv INTEGER NOT NULL DEFAULT 0, -- uv
-- reward_people_num INTEGER NOT NULL DEFAULT 0, -- 人均奖励次数
report_date DATE NOT  NULL DEFAULT CURRENT_DATE, --上报时间
UNIQUE (os,app_key,bundle_id,report_date)
);

