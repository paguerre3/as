package infrastructure

import (
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/labstack/gommon/log"
)

const (
	// no way to load from config file as it will produce a cycle:
	cmdDir      = "/cmd/"
	internalDir = "/internal/"
)

func NewPathResolver() func(string) string {
	var (
		path string
		once sync.Once
	)

	return func(pathOfDomain string) string {
		once.Do(func() {
			// Get the directory of the executable
			wdir, err := os.Getwd()
			if err != nil {
				panic(err)
			}
			index := strings.Index(wdir, cmdDir)
			if index == -1 {
				index = strings.Index(wdir, internalDir)
				if index == -1 {
					log.Warnf("cannot find %s or %s in %s using dafault wdir as root directory", cmdDir, internalDir, wdir)
				}
			}
			var rootDir string
			if index != -1 {
				rootDir = wdir[:index]
			} else {
				rootDir = wdir
			}
			path = filepath.Join(rootDir, pathOfDomain)
		})
		return path
	}
}
