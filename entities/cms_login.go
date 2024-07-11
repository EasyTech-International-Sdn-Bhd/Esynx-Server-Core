package entities

type CmsLogin struct {
	LoginId     int64  `xorm:"not null pk autoincr index BIGINT"`
	StaffCode   string `xorm:"unique VARCHAR(20)"`
	Login       string `xorm:"VARCHAR(30)"`
	Password    string `xorm:"VARCHAR(30)"`
	Name        string `xorm:"VARCHAR(200)"`
	Email       string `xorm:"VARCHAR(250)"`
	ContactNo   string `xorm:"VARCHAR(100)"`
	DeviceToken string `xorm:"VARCHAR(100)"`
	RoleId      int    `xorm:"default 2 comment('Officer, Salesperson, Admin') INT"`
	Remark      string `xorm:"LONGTEXT(4294967295)"`
	LoginStatus int    `xorm:"default 1 comment('1=active, 0=inactive , please check the disable salesperson is not allowed to send in order.') INT"`
	DocSuffix   string `xorm:"default 'S' VARCHAR(10)"`
	ProjNo      string `xorm:"VARCHAR(100)"`
	SessionId   string `xorm:"default '' VARCHAR(100)"`
	Company     string `xorm:"default '' VARCHAR(50)"`
}

func (m *CmsLogin) TableName() string {
	return "cms_login"
}
