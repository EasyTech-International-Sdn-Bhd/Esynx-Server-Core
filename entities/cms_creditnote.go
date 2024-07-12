package entities

import (
	"time"
)

type CmsCreditnote struct {
	CnId             uint64    `xorm:"pk autoincr unique UNSIGNED BIGINT"`
	CnCode           string    `xorm:"index unique VARCHAR(20)"`
	CustCode         string    `xorm:"index VARCHAR(20)"`
	CnDate           time.Time `xorm:"TIMESTAMP"`
	CnUdf            string    `xorm:"not null JSON"`
	CnAmount         float64   `xorm:"DOUBLE"`
	SalespersonId    int       `xorm:"INT"`
	Cancelled        string    `xorm:"CHAR(1)"`
	UpdatedAt        time.Time `xorm:"not null default CURRENT_TIMESTAMP TIMESTAMP"`
	CnKnockoffAmount float64   `xorm:"DOUBLE"`
	Approved         int       `xorm:"default 0 INT"`
	Approver         string    `xorm:"VARCHAR(200)"`
	ApprovedAt       time.Time `xorm:"DATETIME"`
	FromDoc          string    `xorm:"default 'SL' ENUM('AR','SL')"`
	RefNo            string    `xorm:"VARCHAR(20)"`
}

func (m *CmsCreditnote) TableName() string {
	return "cms_creditnote"
}

func (m *CmsCreditnote) Validate() {

}

func (m *CmsCreditnote) ToUpdate() {
	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = time.Now()
	}
}
