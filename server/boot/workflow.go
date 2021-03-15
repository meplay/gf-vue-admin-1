package boot

import model "gf-vue-admin/app/model/workflow"

var Workflow = new(_workflow)

type _workflow struct{}

func (w *_workflow) model() {
	model.WorkflowBusinessStruct = make(map[string]func() model.Workflow)
	model.WorkflowBusinessStruct["leave"] = func() model.Workflow {
		return new(model.LeaveWorkflow)
	}
}

func (w *_workflow) table() {
	model.WorkflowBusinessTable = make(map[string]func() interface{})
	model.WorkflowBusinessTable["leave"] = func() interface{} {
		return new(model.WorkflowLeave)
	}
}

func (w *_workflow) Initialize() {
	w.model()
	w.table()
}
