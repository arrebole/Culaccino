package config

import "path/filepath"

// SourceRoot 存放数据文件的更目录
var SourceRoot, _ = filepath.Abs("data")

// StaticRoot ...
var StaticRoot, _ = filepath.Abs("themes/build")

// DirectoryfilePath 存放目录的文件 相对路径
var DirectoryfilePath = filepath.Join(SourceRoot, "index.json")
