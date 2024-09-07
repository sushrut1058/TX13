package directory

func bfs(start string, length int) []string {
	queue := make([]string, 0)
	final := make([]string, 0)

	// Start from the root node (assuming "0" is the root)
	root := start
	queue = append(queue, root)
	i := 0
	for len(queue) > 0 && i < length {
		// Dequeue a vertex from queue
		node := queue[0]
		queue = queue[1:]
		_, boo := edges[node]
		if !boo {
			i += 1
			final = append(final, node)
		}

		queue = append(queue, edges[node]...)

	}
	return final
}

func Traverse(length int) []string {
	return bfs("/", length)
}

func ReturnTopLevel() []string {
	return edges["/"]
}
