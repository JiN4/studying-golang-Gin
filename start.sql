CREATE TABLE user (
id bigint(20) unsigned NOT NULL AUTO_INCREMENT,
nickname varchar(190) NOT NULL COMMENT 'ニックネーム',
PRIMARY KEY (id)
) ENGINE=InnoDB AUTO_INCREMENT=1002 DEFAULT CHARSET=utf8mb4 COMMENT='テスト';

insert into user (nickname) values ('itiro'),('ziro'),('saburo'),('siro'),('goro'),('rokuro'),('sitiro'),('hatiro'),('kuro'),('toro');

select * from user;