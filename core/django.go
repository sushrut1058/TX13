package core

import (
	"os"
	"path/filepath"
	"promptgen/directory"
	"promptgen/types"
)

var Files map[string]string
var UrlsPy []string

func FindUrlsPy(path string) []string {
	_, Files := directory.GetDirectoryStructure(path)
	// fmt.Println(Files)
	q := make([]string, 0)
	q = append(q, "/")

	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		// fmt.Println(node)
		for _, entry := range Files[node] {
			// fmt.Println(GetBaseName(entry))
			if filepath.Base(entry) == "urls.py" {
				UrlsPy = append(UrlsPy, entry)
			}
			q = append(q, entry)
		}
	}
	return UrlsPy
}

func readFile(path string)

func fetchAllFiles(urls []string) {
	content := ""
	for _, name := range urls {
		data, _ := os.ReadFile(name)
		content = content + "[START OF FILE:" + name + "]\n" + string(data) + "\n[END OF FILE:" + name + "]\n"
	}
	prompt := ""
	file := ""
	message := []types.MessageStruct{
		{
			Role:    "system",
			Content: []string{prompt},
		},
		{
			Role:    "user",
			Content: []string{file},
		},
	}
	jsonBytes, _ :=
		getResponse("")
}

func sendPrompt()
