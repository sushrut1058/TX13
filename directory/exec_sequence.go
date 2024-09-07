package directory

func ExecuteSequence(dir string, length int) ([]string, []string) {
	GetDirectoryStructure(dir)
	return dirs, Traverse(length)
}
