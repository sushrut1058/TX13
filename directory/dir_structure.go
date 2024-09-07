package directory

import (
	"os"
	"path/filepath"
)

// List of directories to ignore
var ignoreDirs = map[string]bool{
	".git":         true,
	"node_modules": true,
	"vendor":       true,
	"__pycache__":  true,
}

var dirs []string
var edges map[string][]string

func genDirTree(cur_path, abs_path string) {
	entries, _ := os.ReadDir(filepath.Join(abs_path, cur_path))

	for _, entry := range entries {
		if entry.IsDir() && ignoreDirs[entry.Name()] {
			continue
		}
		relPath := filepath.Join(cur_path, entry.Name())
		edges[cur_path] = append(edges[cur_path], relPath)
		dirs = append(dirs, relPath)
		if entry.IsDir() {
			genDirTree(relPath, abs_path)
		}
	}

}

// GetDirectoryStructure returns the directory structure of the given path
func GetDirectoryStructure(path string) ([]string, map[string][]string) {
	abs_path, _ := filepath.Abs(path)
	dirs = []string{}
	edges = make(map[string][]string)
	genDirTree("/", abs_path)
	return dirs, edges
}
