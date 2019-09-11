package models

type SqbUser struct {
	Id           int64  `xorm:"pk autoincr BIGINT(20)"`
	Name         string `xorm:"not null default '' index VARCHAR(100)"`
	Phone        string `xorm:"not null default '' index unique(UNI_PHONE_CUSTOMIZEDID) VARCHAR(20)"`
	Password     string `xorm:"not null default '' VARCHAR(255)"`
	Avatar       string `xorm:"not null default '' VARCHAR(100)"`
	Gender       int    `xorm:"not null default 0 comment('性别1男 2女 3宝妈') TINYINT(1)"`
	Role         int    `xorm:"not null default 0 comment('用户角色') index TINYINT(3)"`
	Created      int    `xorm:"not null default 0 INT(11)"`
	Updated      int    `xorm:"not null default 0 INT(11)"`
	Customizedid int    `xorm:"not null default 0 comment('定制化ID') unique(UNI_PHONE_CUSTOMIZEDID) INT(11)"`
	Isdeleted    int    `xorm:"not null default 0 TINYINT(3)"`
}
