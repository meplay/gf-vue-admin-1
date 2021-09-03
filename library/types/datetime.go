package types

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"go.uber.org/zap"
	"time"
)

const DateTimeFormat = "2006-01-02 15:04:05"

type Datetime struct {
	time.Time
}

// NewDatetime 构建 Datetime 方法
// Author SliverHorn
func NewDatetime(value string) *Datetime {
	_time, _ := time.Parse(DateTimeFormat, value)
	return &Datetime{Time: _time}
}

// NewDatetimeByLayout 构建 Datetime by layout
// Author SliverHorn
func NewDatetimeByLayout(layout, value string) *Datetime {
	_time, _ := time.Parse(layout, value)
	_d := &Datetime{Time: _time}
	return _d
}

// ToDate Datetime 2 Date
// Author SliverHorn
func (d *Datetime) ToDate() *Date {
	if d == nil {
		return nil
	}
	return &Date{Time: d.Time}
}

// ToTime Datetime 2 *time.Time
// Author SliverHorn
func (d *Datetime) ToTime() *time.Time {
	if d == nil {
		return nil
	}
	return &d.Time
}

// Scan 扫描
// Author SliverHorn
func (d *Datetime) Scan(value interface{}) error {
	nullTime := &sql.NullTime{}
	if err := nullTime.Scan(value); err != nil {
		zap.L().Error("Date To Database 时间转换失败!", zap.Any("datetime", value))
		return err
	} else {
		*d = Datetime{Time: nullTime.Time}
		return nil
	}
}

// Value 值
// Author SliverHorn
func (d Datetime) Value() (driver.Value, error) {
	return driver.Value(d.Time.Format(DateTimeFormat)), nil
}

// MarshalJSON 序列化
// Author SliverHorn
func (d *Datetime) MarshalJSON() ([]byte, error) {
	return []byte(d.Time.Format(`"2006-01-02 15:04:05"`)), nil
}

// UnmarshalJSON 反序列化
// Author SliverHorn
func (d *Datetime) UnmarshalJSON(b []byte) error {
	if string(b) > `""` {
		if _time, err := time.Parse(`"2006-01-02 15:04:05"`, string(b)); err != nil {
			zap.L().Error("时间转换失败!", zap.String("datetime", string(b)))
			return err
		} else {
			*d = Datetime{Time: _time}
			return nil
		}
	}
	return nil
}

// GormDataType gorm 定义数据库字段类型
// Author SliverHorn
func (d *Datetime) GormDataType() string {
	return "datetime"
}

// CostMonth 获取时间所属月份 2021-05-13 17:48:00 => May-21
// Author SliverHorn
func (d *Datetime) CostMonth() string {
	if d == nil {
		return ""
	}
	_time := d.Time
	a := fmt.Sprintf("%v", _time.Month())
	if len(a) > 3 { // 取月份简称
		return fmt.Sprintf("%v-%v", a[:3], _time.Year()-2000)
	}
	return fmt.Sprintf("%v-%v", a, _time.Year()-2000)
}
