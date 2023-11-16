package models

import "time"

type Kamar struct {
	Id          string    `gorm:"type:varchar(20); primaryKey" json:"id"`
	Markusnicu  string    `gorm:"type:varchar(20)" json:"markusnicu"`
	Markusvvip  string    `gorm:"type:varchar(20)" json:"markusvvip"`
	Markusvip   string    `gorm:"type:varchar(20)" json:"markusvip"`
	Lukas       string    `gorm:"type:varchar(20)" json:"lukas"`
	Maria       string    `gorm:"type:varchar(20)" json:"maria"`
	Fransiskus  string    `gorm:"type:varchar(20)" json:"fransiskus"`
	Matius      string    `gorm:"type:varchar(20)" json:"matius"`
	Teresia     string    `gorm:"type:varchar(20)" json:"teresia"`
	Teresiatiga string    `gorm:"type:varchar(20)" json:"teresiatiga"`
	Yosef       string    `gorm:"type:varchar(20)" json:"yosef"`
	Klara       string    `gorm:"type:varchar(20)" json:"klara"`
	Egidio      string    `gorm:"type:varchar(20)" json:"egidio"`
	Yohanes     string    `gorm:"type:varchar(20)" json:"yohanes"`
	UpdatedAt   time.Time `json:"updated_at"`
}