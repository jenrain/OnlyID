package entity

type Segments struct {
	BizTag     string `json:"biz_tag" gorm:"not null pk VARCHAR(128) 'biz_tag'"`
	MaxId      int64  `json:"max_id" gorm:"BIGINT(20) 'max_id'"`
	Step       int64  `json:"step" gorm:"INT(11) 'step'"`
	CreateTime int64  `json:"create_time" gorm:"BIGINT(20) 'create_time'"`
	UpdateTime int64  `json:"update_time" gorm:"BIGINT(20) 'update_time'"`
}

func (Segments) TableName() string {
	return "segments"
}
