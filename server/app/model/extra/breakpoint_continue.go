package model

import "flipped-aurora/gf-vue-admin/server/library/global"

// BreakpointContinue 文件结构体
type BreakpointContinue struct {
	global.Model
	FileMd5    string                    `json:"FileMd5"`
	FileName   string                    `json:"FileName"`
	FilePath   string                    `json:"FilePath"`
	IsFinish   bool                      `json:"IsFinish"`
	ChunkTotal int                       `json:"ChunkTotal"`
	FileChunk  []BreakpointContinueChunk `orm:"-" gorm:"foreignKey:FileID"`
}

func (b *BreakpointContinue) TableName() string {
	return "breakpoint_continue"
}

// BreakpointContinueChunk 切片结构体
type BreakpointContinueChunk struct {
	global.Model
	FileChunkPath   string
	FileID          uint
	FileChunkNumber int
}

func (b *BreakpointContinueChunk) TableName() string {
	return "breakpoint_continue_chunks"
}
