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

// File ...
type File struct {
	Root   string
	Data   []byte
	Hash   string
	Suffix string
}

//CheckHash 计算hash值
func (p *File) CheckHash() string {
	if p.Hash != "" {
		return p.Hash
	}
	hash := md5.Sum(p.Data)
	result := hex.EncodeToString(hash[:])
	p.Hash = result
	return result
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
	return fmt.Sprintf("/static/%s/%s", p.CheckDate(), p.FileName())

}

// SaveFile ...
func (p *File) SaveFile() error {
	os.Mkdir(p.SavePath(), os.ModePerm)

	file, err := os.Create(p.FullFileName())
	if err != nil {
		return errors.New("create file fail")
	}
	defer file.Close()

	_, err = file.Write(p.Data)
	if err != nil {
		return errors.New("write file fail")
	}

	return nil
}
