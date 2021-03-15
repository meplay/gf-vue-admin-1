package model

import "gf-vue-admin/library/global"

type WorkflowStartPoint struct {
	global.Model
	X              float64 `json:"x"`
	Y              float64 `json:"y"`
	Index          int     `json:"index"`
	WorkflowEdgeID string
}

type WorkflowEndPoint struct {
	global.Model
	X              float64 `json:"x"`
	Y              float64 `json:"y"`
	Index          int     `json:"index"`
	WorkflowEdgeID string
}
