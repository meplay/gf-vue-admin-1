package boot

import (
	model "gf-vue-admin/app/model/system"
	"gf-vue-admin/library/excel"
	"testing"
)

func TestName(t *testing.T) {
	excel.Export.Export(model.Admin{})
}
