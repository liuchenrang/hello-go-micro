package models

type SqbApp struct {
	Id             int    `xorm:"not null pk autoincr INT(11)  'id'"`
	UserId         int64  `xorm:"not null default 0 index BIGINT(20)  'userId'"`
	Name           string `xorm:"not null default '' VARCHAR(20)  'name'"`
	Logo           string `xorm:"not null default '' VARCHAR(100)  'logo'"`
	SplashScreen   string `xorm:"default '' comment('启动页') VARCHAR(100)  'splashScreen'"`
	Pid            string `xorm:"not null default '' index VARCHAR(50)  'pid'"`
	PidName        string `xorm:"not null default '' VARCHAR(50)  'pidName'"`
	MemberId       int64  `xorm:"not null default 0 index BIGINT(20)  'memberId'"`
	Qq             string `xorm:"default '' VARCHAR(32)  'qq'"`
	Wechat         string `xorm:"default '' VARCHAR(32)  'wechat'"`
	Code           string `xorm:"not null default '' comment('省钱口令') unique VARCHAR(32)  'code'"`
	Status         string `xorm:"not null default '1' comment('1-启用 0-停用') VARCHAR(2)  'status'"`
	NewActivateNum int    `xorm:"not null default 0 comment('新总激活数') INT(11)  'newActivateNum'"`
	ActivateNum    int    `xorm:"not null default 0 comment('总激活数') INT(11)  'activateNum'"`
	Created        int    `xorm:"not null default 0 created INT(11)  'created'"`
	Updated        int    `xorm:"not null default 0 updated INT(11)  'updated'"`
	CustomizedId   int    `xorm:"not null default 0 comment('定制化ID') INT(11)  'customizedId'"`
	IsDeleted      int    `xorm:"not null default 0 deleted TINYINT(1)  'isDeleted'"`
	EnCode         string `xorm:"not null default '' comment('新省钱口令') index VARCHAR(20)  'enCode'"`
	OutDomain      string `xorm:"not null default '' comment('自定义域名') VARCHAR(128)  'outDomain'"`
}
