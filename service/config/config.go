package config

import "path/filepath"

// SourceRootPath 存放目录的文件 相对路径
var SourceRootPath, _ = filepath.Abs("data")

// DirectoryfilePath 存放目录的文件 相对路径
var DirectoryfilePath = filepath.Join(SourceRootPath, "index.json")
