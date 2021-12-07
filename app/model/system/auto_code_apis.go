package system

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"

	"github.com/flipped-aurora/gf-vue-admin/library/common"
	"github.com/pkg/errors"
)

type AutoCodeApis []Api

// Scan 扫描
// Author [SliverHorn](https://github.com/SliverHorn)
func (a *AutoCodeApis) Scan(value interface{}) error {
	switch v := value.(type) {
	case []byte:
		if err := json.Unmarshal(v, a); err != nil {
			return err
		}
	case string:
		if err := json.Unmarshal([]byte(v), a); err != nil {
			return err
		}
	default:
		return errors.New(fmt.Sprint("Failed to unmarshal AutoCodeApis value:", value))
	}
	return nil
}

// Value 值
// Author [SliverHorn](https://github.com/SliverHorn)
func (a AutoCodeApis) Value() (driver.Value, error) {
	bytes, err := json.Marshal(&a)
	if err != nil {
		return nil, err
	}
	return driver.Value(bytes), err
}

func (a *AutoCodeApis) ToCommonGetByID() *common.GetByIDs {
	if a == nil {
		return nil
	}
	apis := *a
	length := len(apis)
	ids := make([]uint, 0, length)
	for i := 0; i < length; i++ {
		ids = append(ids, apis[i].ID)
	}
	return &common.GetByIDs{Ids: ids}
}
