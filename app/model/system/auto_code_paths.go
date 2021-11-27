package system

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"github.com/flipped-aurora/gf-vue-admin/library/global"
	"github.com/flipped-aurora/gf-vue-admin/library/types"
	"github.com/pkg/errors"
	"path/filepath"
	"time"
)

type AutoCodePath struct {
	Filepath string `json:"filepath"`
}

func (a *AutoCodePath) RmFilePath() string {
	path := global.Config.AutoCode.RubbishPath
	if path == "" {
		path = global.Config.AutoCode.Root
	}
	return filepath.Join(path, "files", time.Now().Format(types.DateTimeFormat), filepath.Base(filepath.Dir(filepath.Dir(a.Filepath))), filepath.Base(filepath.Dir(a.Filepath)), filepath.Base(a.Filepath))
}

type AutoCodePaths []AutoCodePath

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
