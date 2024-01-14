package configs

import (
	"path/filepath"
	"runtime"
	"strings"
	"sync"

	"github.com/joho/godotenv"
)

var onceDevelop = new(sync.Once)

func Develop() {
	onceDevelop.Do(func() {
		_, caller, _, _ := runtime.Caller(0)
		dir := filepath.Dir(caller)

		idx := strings.Index(dir, "estore")
		if idx < 0 {
			return
		}

		godotenv.Load(dir[:idx] + "\\estore\\.env")
	})
}
