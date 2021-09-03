package types

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"go.uber.org/zap"
	"time"
)

const DateFormat = "2006-01-02"

type Date struct {
	time.Time
}

// NewDate 构建 Date 方法
// Author SliverHorn
func NewDate(value string) *Date {
	_time, _ := time.Parse(DateFormat, value)
	return &Date{Time: _time}
}

// NewDateByLayout 构建 Date by layout
// Author SliverHorn
func NewDateByLayout(layout, value string) *Date {
	_time, _ := time.Parse(layout, value)
	return &Date{Time: _time}
}

// ToDatetime Date to Datetime
// Author SliverHorn
func (d *Date) ToDatetime() *Datetime {
	return &Datetime{Time: d.Time}
}

// MarshalJSON 序列化
// Author SliverHorn
func (d *Date) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, d.Time.Format("2006-01-02"))), nil
}

// UnmarshalJSON 反序列化
// Author SliverHorn
func (d *Date) UnmarshalJSON(b []byte) error {
	if string(b) > `""` {
		if _time, err := time.Parse(`"2006-01-02"`, string(b)); err != nil {
			zap.L().Error("Date 时间转换失败!", zap.String("date", string(b)))
			return err
		} else {
			d.Time = _time
			return nil
		}
	}
	return nil
}

// Scan 扫描
// Author SliverHorn
func (d *Date) Scan(value interface{}) error {
	nullTime := &sql.NullTime{}
	if err := nullTime.Scan(value); err != nil {
		zap.L().Error("Date To Database 时间转换失败!", zap.Any("Date", value))
		return err
	} else {
		d.Time = nullTime.Time
		return nil
	}
}

// Value 值
// Author SliverHorn
func (d Date) Value() (driver.Value, error) {
	return driver.Value(d.Time.Format(DateFormat)), nil
}

// GormDataType gorm 定义数据库字段类型
// Author SliverHorn
func (d Date) GormDataType() string {
	return "date"
}

// CostMonth 获取时间所属月份 2021-05-13 17:48:00 => May-21
// Author SliverHorn
func (d *Date) CostMonth() string {
	if d == nil {
		return ""
	}
	a := fmt.Sprintf("%v", d.Time.Month())
	if len(a) > 3 { // 取月份简称
		return fmt.Sprintf("%v-%v", a[:3], d.Time.Year()-2000)
	}
	return fmt.Sprintf("%v-%v", a, d.Time.Year()-2000)
}
