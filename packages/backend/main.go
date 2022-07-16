package main

import (
	_ "embed"
	"io"
	"io/fs"
	"strings"

	"taskism/controllers"
	"taskism/utils"

	"github.com/mholt/archiver/v4"
)

var staticFS fs.FS

//go:embed assets.zip
var staticZip string

func main() {
	staticFS = archiver.ArchiveFS{
		Stream: io.NewSectionReader(strings.NewReader(staticZip), 0, int64(len(staticZip))),
		Format: archiver.Zip{},
	}
	utils.InitStatic(staticFS)
	r := controllers.GinEngine()

	r.Run("localhost:3001")
}
