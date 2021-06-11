package boot

import (
	model "flipped-aurora/gf-vue-admin/server/app/model/system"
	"flipped-aurora/gf-vue-admin/server/library/excel"
	"testing"
)

func TestName(t *testing.T) {
	excel.Export.Export(model.Admin{})
}
