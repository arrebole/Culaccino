package model

import "encoding/json"

// Builder ...
type Builder interface {
	Build([]byte) []byte
}

// DirectoryBuilder ...
type DirectoryBuilder struct{}

// Build ...
func (p DirectoryBuilder) Build(data []byte) []byte {
	var directoryItems []DirectoryItem
	json.Unmarshal(data, &directoryItems)

	buffer, _ := json.Marshal(&Response{0, "目录", directoryItems})
	return buffer
}

// SourceBuilder ...
type SourceBuilder struct{}

// Build ...
func (p SourceBuilder) Build(data []byte) []byte {
	buffer, _ := json.Marshal(&Response{0, "详细内容", string(data)})
	return buffer
}
