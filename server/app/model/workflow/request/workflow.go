package request

import model "gf-vue-admin/app/model/workflow"

type SearchWorkflowProcess struct {
	model.WorkflowProcess
	PageInfo
}
