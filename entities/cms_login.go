package entities

type CmsLogin struct {
	LoginId     uint64 `xorm:"pk autoincr unique UNSIGNED BIGINT" json:"loginId,omitempty" xml:"loginId"`
	StaffCode   string `xorm:"unique VARCHAR(20)" json:"staffCode,omitempty" xml:"staffCode"`
	Login       string `xorm:"VARCHAR(30)" json:"login,omitempty" xml:"login"`
	Password    string `xorm:"VARCHAR(30)" json:"password,omitempty" xml:"password"`
	Name        string `xorm:"VARCHAR(200)" json:"name,omitempty" xml:"name"`
	Email       string `xorm:"VARCHAR(250)" json:"email,omitempty" xml:"email"`
	ContactNo   string `xorm:"VARCHAR(100)" json:"contactNo,omitempty" xml:"contactNo"`
	DeviceToken string `xorm:"VARCHAR(100)" json:"deviceToken,omitempty" xml:"deviceToken"`
	RoleId      int    `xorm:"default 2 comment('Officer, Salesperson, Admin') INT" json:"roleId,omitempty" xml:"roleId"`
	Remark      string `xorm:"BLOB" json:"remark,omitempty" xml:"remark"`
	LoginStatus int    `xorm:"default 1 comment('1=active, 0=inactive , please check the disable salesperson is not allowed to send in order.') INT" json:"loginStatus,omitempty" xml:"loginStatus"`
	DocSuffix   string `xorm:"default 'S' VARCHAR(10)" json:"docSuffix,omitempty" xml:"docSuffix"`
	ProjNo      string `xorm:"VARCHAR(100)" json:"projNo,omitempty" xml:"projNo"`
	SessionId   string `xorm:"default '' VARCHAR(100)" json:"sessionId,omitempty" xml:"sessionId"`
	Company     string `xorm:"default '' VARCHAR(50)" json:"company,omitempty" xml:"company"`
}

func (m *CmsLogin) TableName() string {
	return "cms_login"
}

func (m *CmsLogin) BeforeInsert() {
	m.BeforeUpdate()
}

func (m *CmsLogin) BeforeUpdate() {
	if m.RoleId == 0 {
		m.RoleId = 2
	}
}
