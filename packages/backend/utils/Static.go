package utils

import (
	"io/fs"
	"net/http"

	"github.com/gin-gonic/contrib/static"
)

const StaticFolder = "static"

var StaticFS static.ServeFileSystem

type GinFS struct {
	FS http.FileSystem
}

// Open 打开文件
func (b *GinFS) Open(name string) (http.File, error) {
	return b.FS.Open(name)
}

// Exists 文件是否存在
func (b *GinFS) Exists(prefix string, filepath string) bool {
	if _, err := b.FS.Open(filepath); err != nil {
		return false
	}
	return true
}

// InitStatic 初始化静态资源文件
func InitStatic(statics fs.FS) {
	if Exists(RelativePath(StaticFolder)) {
		StaticFS = static.LocalFile(RelativePath("static"), false)
	} else {
		StaticFS = &GinFS{
			FS: http.FS(statics),
		}
	}
}
