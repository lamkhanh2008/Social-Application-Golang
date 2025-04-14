package common

import "time"

type SQLModel struct {
	Id        int        `json:"-" gorm:"column:id;"`
	FakeId    *UID       `json:"id" gorm:"-"`
	CreatedAt *time.Time `json:created_at, omitempty" gorm:"column:created_at;"`
	UpdatedAt *time.Time `json:"updated_at, omitempty" gorm:"column:updated_at;"`
}

func (sqlModel *SQLModel) Mask(dbType DBType) {
	uid := NewUID(uint32(sqlModel.Id), int(dbType), 1)
	sqlModel.FakeId = &uid
}

type SimpleUser struct {
}
