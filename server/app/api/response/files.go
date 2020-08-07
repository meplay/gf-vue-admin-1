package response

import (
	"server/app/model/breakpoint_files"
	"server/app/model/files"
)

type Files struct {
	File *files.Entity `json:"file"`
}

type ExaFile struct {
	File *breakpoint_files.Entity `json:"file"`
}
