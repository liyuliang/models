package database

import (
	"time"
	"strings"
	"gitee.com/liyuliang/utils/format"
)

func init() {
	Register(func() model {
		return new(Job)
	})
}

type Job struct {
	ID          uint      `gorm:"primary_key"`
	JobId       string    `sql:"type:VARCHAR(255);index:jid"`
	ParentJobId string    `sql:"type:VARCHAR(255);index:pjid"`
	UserId      string    `sql:"type:VARCHAR(255);index:uid"`
	CreateTime  time.Time `sql:"DEFAULT:current_timestamp"`
	DelayTime   time.Time `sql:"DEFAULT:current_timestamp"`
	TTL         time.Time `sql:"DEFAULT:current_timestamp"`
	FinishTime  *time.Time
	Type        string    `sql:"type:VARCHAR(50)"`
	Data        string    `sql:"type:text"`
}

func (table *Job) Name() string {
	// custom table name, this is default
	return "JOB"
}

func (table *Job) BeforeSave() (err error) {
	//err = errors.New("Permission denied.")
	return err
}

func (table *Job) SetData(data map[string]string) {

	jobData := format.MapData()

	for key, value := range data {
		jobData[strings.ToLower(key)] = value
	}
	table.Data = jobData.ToString()
}
