package web

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
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
					err = fmt.Errorf("cannot find %s or %s in %s", cmdDir, internalDir, wdir)
					panic(err)
				}
			}
			rootDir := wdir[:index]
			path = filepath.Join(rootDir, pathOfDomain)
		})
		return path
	}
}
