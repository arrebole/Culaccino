package module

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// FileStatus ...
type FileStatus struct {
	Hash string `json:"hash"`
	URL  string `json:"url"`
}

// File ...
type File struct {
	Root       string
	Data       []byte
	Suffix     string
	FileStatus *FileStatus
}

//CheckHash 计算hash值
func (p *File) CheckHash() string {
	if p.FileStatus.Hash != "" {
		return p.FileStatus.Hash
	}
	hash := md5.Sum(p.Data)
	result := hex.EncodeToString(hash[:])
	return p.SetHash(result)
}

// SetHash ...
func (p *File) SetHash(hash string) string {
	p.FileStatus.Hash = hash
	return p.FileStatus.Hash
}

// SetRoot 设置路径
func (p *File) SetRoot(path string) *File {
	rootPath, _ := filepath.Abs(path)
	p.Root = rootPath
	return p
}

//CheckDate 计算时间
func (p *File) CheckDate() string {
	return time.Now().Format("2006-01-02")
}

// FileName 文件名称
func (p *File) FileName() string {
	return fmt.Sprintf("%s.%s", p.CheckHash(), p.Suffix)
}

// SavePath 返回储存路径
func (p *File) SavePath() string {
	return filepath.Join(p.Root, p.CheckDate())
}

// FullFileName 返回带有文件名的储存路径
func (p *File) FullFileName() string {
	return filepath.Join(p.SavePath(), p.FileName())
}

// URL 文件的链接
func (p *File) URL() string {
	var url = fmt.Sprintf("/static/%s/%s", p.CheckDate(), p.FileName())
	p.FileStatus.URL = url
	return url
}

// SaveFile ...
func (p *File) SaveFile() (*FileStatus, error) {
	p.FileStatus = &FileStatus{}
	os.Mkdir(p.SavePath(), os.ModePerm)

	file, err := os.Create(p.FullFileName())
	if err != nil {
		return nil, errors.New("create file fail")
	}
	defer file.Close()

	_, err = file.Write(p.Data)
	if err != nil {
		return nil, errors.New("write file fail")
	}

	return p.FileStatus, nil
}
