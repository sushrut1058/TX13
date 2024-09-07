package classification

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
)

var djangoFiles = []string{
	"manage.py",
	"settings.py",
	"urls.py",
	"wsgi.py",
	"asgi.py",
	"models.py",
}

func containsDjangoImports(filePath string) (bool, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return false, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(line, "from django") || strings.HasPrefix(line, "import django") {
			return true, nil
		}
	}

	return false, scanner.Err()
}

func isDjangoProject(root string) bool {
	var foundManagePy, foundSettingsPy, foundDjangoImport bool

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip directories like .git, node_modules, etc.
		if info.IsDir() && (info.Name() == ".git" || info.Name() == "node_modules" || info.Name() == "__pycache__") {
			return filepath.SkipDir
		}

		// Check for the manage.py file
		if info.Name() == "manage.py" {
			foundManagePy = true
		}

		// Check for settings.py file
		if info.Name() == "settings.py" {
			foundSettingsPy = true
		}

		// Check for Django-specific imports in Python files
		if strings.HasSuffix(path, ".py") {
			if found, _ := containsDjangoImports(path); found {
				foundDjangoImport = true
			}
		}

		// Early exit if we have found enough indicators
		if foundManagePy && foundSettingsPy && foundDjangoImport {
			return filepath.SkipDir // no need to walk further
		}

		return nil
	})

	if err != nil {
		return false
	}

	// If we found manage.py, settings.py, and a Django import, we can definitively say it's a Django project
	if foundManagePy && foundSettingsPy && foundDjangoImport {
		return true
	}

	return false
}
