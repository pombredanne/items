

-- 活动基本信息
create table campaign_campaign(
id serial primary key,
name varchar(300) NOT NULL, --活动名称
user_id INTEGER NOT NULL, --用户ID
bundle_id varchar(300) NOT NULL, --应用包名
app_platform VARCHAR(10) NOT NULL,
category VARCHAR(10) NOT NULL, --应用分类  100100 100
sub_category VARCHAR(10) NOT NULL, --应用分类  100100 100
-- budget NUMERIC(12,4) NOT NULL DEFAULT 100, --预算
budget_day INTEGER NOT NULL DEFAULT 100, --每日预算
bidding_max FLOAT NOT NULL DEFAULT 0,
bidding_min FLOAT NOT NULL DEFAULT 0,
-- bidding numrange NOT NULL default '[0,0]', --出价范围
bidding_type VARCHAR(10) NOT NULL DEFAULT 'CPM', --出价类型
freq jsonb NOT NULL DEFAULT '{"open":true,"type":"day","num":100}', --投放频次  type hour,day
targeting jsonb NOT NULL,
url jsonb NOT NULL,
-- freq_hour INTEGER NOT NULL DEFAULT 100,--投放频次 时
-- freq_day INTEGER NOT NULL DEFAULT 100,--投放频次 天
speed SMALLINT NOT NULL DEFAULT 1, --投放速率 ((1,加速),(2,匀速),)
-- start_date DATE NOT NULL DEFAULT CURRENT_DATE, --投放开始时间 默认当天
-- end_date DATE NOT NULL DEFAULT CURRENT_DATE, --投放结束时间
create_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP, --创建时间
status SMALLINT NOT NULL DEFAULT 0-- 活动状态 ((0,暂停),(1,活跃),)
);

create table ad_creative (
id serial PRIMARY key,
name VARCHAR(100) NOT NULL,
width INTEGER  NOT NULL ,
height INTEGER  NOT NULL ,
ad_type VARCHAR(100) NOT NULL,
ad_size VARCHAR(100) NOT NULL,
material VARCHAR(100) NOT NULL,
ol SMALLINT  NOT  NULL ,
creative_id NOT NULL ,
description VARCHAR(300) NOT NULL,
create_date DATE NOT NULL DEFAULT CURRENT_DATE
);



-- -- 活动目标(广告组)
-- create table campaign_targeting(
-- id serial primary key,
-- info json,--{geo_code:[],app_os:,
-- --app_category_lv1:,app_category_lv2:,device_type:[],os_version:[],hour:[],
-- --carrier:[],network:[],brand[]}
--
--
--
--
-- -- info json, --province:[]
-- -- {
-- --   "3300": {"allcity":1, cities:[]},
-- --   "3300": {"allcity":0, cities:[3301, 33002]}
--
-- -- }
--
-- -- {
-- --   "geo_code": ["33001", "33002"]
-- -- }
--
-- -- geo_code, coutry, province, city
-- -- 33000, *, 江西, -
-- -- 33001, *, 江西, 南昌,
-- -- ..
--
--
-- -- province json, --选择区域 省份
-- -- city json, --城市
-- -- app_os SMALLINT  NOT NULL, --应用OS
-- -- app_category_lv1 SMALLINT NOT NULL , --选择类型
-- -- app_category_lv2  SMALLINT NOT NULL , --选择类型
-- -- device_type json, --选择设备类型 格式[Tablet,phone]
-- -- os_version json, --选择系统版本 格式[1.1,1.2,1.3]
-- -- hour json, --时段
-- -- carrier json, --运营商 []
-- -- network json, --连接类型 [] wifi *G
-- -- brand json, --设备品牌筛选 []
-- create_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP, --创建时间
-- campaign_id INTEGER NOT NULL,
--  UNIQUE(campaign_id)
-- );
-- 人群包
create table campaign_segment(
id serial primary key,
name varchar(30) NOT NULL, --名称
user_id INTEGER NOT NULL ,
uv INTEGER NOT NULL DEFAULT 0,
create_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ,--创建时间
update_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ,--修改时间
UNIQUE(name,user_id)
);
--活动广告
create table campaign_ad(
id serial primary key,
name VARCHAR (300) NOT NULL , --广告名称
-- segment_id INTEGER  , -- 人群包
ad_type varchar(30) not null,--分类(video,graphic,playable)
ad_size varchar(30) not null,--创意name(banner,inter,video)
creative_set_id INTEGER not null,--创意集合
ol SMALLINT NOT NULL ,
duration FLOAT NOT NULL DEFAULT 0,
url jsonb,-- {'tracking_imp_url':,'tracking_clk_url':,'jump_url':,'deep_link_url':}
-- tracking_imp_url VARCHAR (300) NOT NULL DEFAULT '', --展示监控地址
-- tracking_clk_url VARCHAR (300) NOT NULL DEFAULT '', --点击监控地址
-- landing_page_url VARCHAR (300) NOT NULL DEFAULT '', --落地页地址
-- deep_link_url VARCHAR (300) NOT NULL DEFAULT '', --跳转应用地址
-- download_url VARCHAR (300) NOT NULL DEFAULT '', --下载地址

creative jsonb NOT NULL ,--  {'inter_l':url,'inter_p':url}{'video':url,'inter_l':url,'inter_p':url}{'img':url}
create_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ,--创建时间
campaign_id INTEGER NOT NULL,
status SMALLINT DEFAULT 0, --状态 (-1 审核失败(0,待审核),(1,已审核),
 UNIQUE (ad_size,creative_set_id,campaign_id,ol)
);






-- 记录素材变更log
--活动素材库  素材
-- create table campaign_material(
-- id serial primary key,
-- name VARCHAR (30) NOT NULL, --素材名称
-- content VARCHAR (30) NOT NULL, --素材内容
-- url VARCHAR (300) NOT NULL, --素材地址
-- format VARCHAR (30) NOT NULL, --素材文件格式
-- size INTEGER NOT NULL DEFAULT 0,--素材大小
-- time SMALLINT NOT NULL DEFAULT 0, --素材时长
-- user_id INTEGER NOT NULL ,--用户ID
-- create_date DATE NOT NULL DEFAULT CURRENT_DATE ,--创建时间
-- create_time INTEGER NOT NULL ,--创建时间戳
-- update_time INTEGER NOT NULL,--修改时间戳
-- status SMALLINT DEFAULT 0 --状态 ((0,待审核),(1,已审核))
-- );
-- 上传时附加 img size  服务验证是否满足[JPG，PNG，GIF] file_size   ['mp4','mov'] time size
create table ad_creative(
    id serial primary key,
    category SMALLINT NOT NULL,--(1,Graphic)(2,video)
    type SMALLINT NOT NULL,--(1,'banner'),(2,'sdk_video')
    info jsonb
)
{'graphic':{'banner':{
'msg':['Portrait Interstitial Image (768x1024, 320x512, or 320x480)',Landscape Interstitial Image (1024x768, 512x320, or 480x320)]
,'creative':['inter_l','inter_p','video']
}}}
