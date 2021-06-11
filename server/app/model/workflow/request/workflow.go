package request

import model "flipped-aurora/gf-vue-admin/server/app/model/workflow"

type SearchWorkflowProcess struct {
	model.WorkflowProcess
	PageInfo
}
