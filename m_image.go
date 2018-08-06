package models

import (
	"time"
)

func init() {
	register(func() model {
		return new(Image)
	})
}

type Image struct {
	ID         uint      `gorm:"primary_key"`
	UUid       string    `sql:"type:VARCHAR(255);index:uuid"`
	OriHost    string    `sql:"type:VARCHAR(255);index:OriHost"`
	OriUrl     string    `sql:"type:VARCHAR(255);index:OriUrl"`
	ImgName    string    `sql:"type:VARCHAR(255)"`
	Host       string    `sql:"type:VARCHAR(50);index:host"`
	Port       string    `sql:"type:VARCHAR(25);index:port"`
	Fid        string    `sql:"type:VARCHAR(50)"`
	FSUrl      string    `sql:"type:VARCHAR(255)"`
	Size       int64     `sql:"type:int unsigned;index:size"`
	CreateTime time.Time `sql:"DEFAULT:current_timestamp"`
}

func (table *Image) Name() string {
	// custom table name, this is default
	return "IMAGE"
}
