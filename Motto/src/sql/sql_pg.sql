


DROP TABLE IF EXISTS "flag";
CREATE TABLE "flag" (
  "thisTableSaveFlag" varchar(255) NULL DEFAULT NULL
);



INSERT INTO flag VALUES ('flag{sql-flag}');




DROP TABLE IF EXISTS "motto_infos";
CREATE TABLE "motto_infos" (
  "motto_id" bigserial NOT NULL,
  "nick_name" varchar(25) NULL DEFAULT NULL,
  "motto" varchar(500) NULL DEFAULT NULL,
  PRIMARY KEY ("motto_id")
);



INSERT INTO motto_infos VALUES (1, 'Yliken', '层楼终究误少年 自由早晚乱余生');
INSERT INTO motto_infos VALUES (3, 'Yliken', '他说你任何为人称道的美丽 不及他第一次遇到你');
INSERT INTO motto_infos VALUES (6, 'Yliken', '你是我患得患失的梦，我是你可有可无的人。毕竟这穿越山河的箭，刺的都是用情致疾的人');
INSERT INTO motto_infos VALUES (7, 'Zer0', '岂不闻天无绝人之路，只要我想走，路就在脚下');
INSERT INTO motto_infos VALUES (8, 'Ch13hh', 'The best way to predict the future is to create it.');
INSERT INTO motto_infos VALUES (9, 'NEWYM', '每天都是新的一天，新的你我');
INSERT INTO motto_infos VALUES (10, E'\177Err0r', '长大了，才能够体会到失去。');
INSERT INTO motto_infos VALUES (11, 'RedBean', '红豆生南国 春来发几枝 愿君多采撷 此物最相思');
INSERT INTO motto_infos VALUES (12, 'p0l1st', '学习是学习 生活是生活');
INSERT INTO motto_infos VALUES (13, '168', '人生如逆旅 我亦是行人');




DROP TABLE IF EXISTS "register_infos";
CREATE TABLE "register_infos" (
  "user_id" bigserial NOT NULL,
  "nickname" varchar(25) NOT NULL,
  "username" varchar(25) NULL DEFAULT NULL,
  "password" varchar(50) NOT NULL,
  PRIMARY KEY ("user_id"),
  UNIQUE ("username")
);



INSERT INTO register_infos VALUES (1, E'Yliken\' -- a', 'admin', '123456');
INSERT INTO register_infos VALUES (2, 'RedBean', 'RedBean', 'RedBean');
INSERT INTO register_infos VALUES (3, 'Zer0', 'Zer0', 'Zer0123456');


