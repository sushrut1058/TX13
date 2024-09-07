package classification

import (
	"promptgen/types"
)

func GetType(project string) int {
	switch true {
	case isDjangoProject(project):
		return types.Django
	default:
		return -1
	}
}
