package system

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"github.com/flipped-aurora/gf-vue-admin/library/global"
	"github.com/pkg/errors"
	"path/filepath"
	"time"
)

type AutoCodePaths struct {
	Filepath string `json:"filepath"`
}

func (a *AutoCodePaths) RmFilePath() string {
	return filepath.Join(global.Config.AutoCode.Root, "rm_file", time.Now().Format("20060102"), filepath.Base(filepath.Dir(filepath.Dir(a.Filepath))), filepath.Base(filepath.Dir(a.Filepath)), filepath.Base(a.Filepath))
}

// Scan 扫描
// Author [SliverHorn](https://github.com/SliverHorn)
func (a *AutoCodePaths) Scan(value interface{}) error {
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
		return errors.New(fmt.Sprint("Failed to unmarshal AutoCodeInjection value:", value))
	}
	return nil
}

// Value 值
// Author [SliverHorn](https://github.com/SliverHorn)
func (a AutoCodePaths) Value() (driver.Value, error) {
	bytes, err := json.Marshal(&a)
	if err != nil {
		return nil, err
	}
	return driver.Value(bytes), err
}