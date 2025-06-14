package clone

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/google/go-github/v62/github"
)

func CloneTheRepo(repo *github.Repository, destPath string) error {

	cloneUrl := repo.GetCloneURL()

	destPath, err := filepath.Abs(destPath)
	if err != nil {
		return fmt.Errorf("invalid path:- %v", err)
	}

	// 3. Check if the path is a directory (if it exists)
	if fileInfo, err := os.Stat(destPath); err == nil {
		if !fileInfo.IsDir() {
			return fmt.Errorf("path exists but is not a directory: %s", destPath)
		}
	} else if os.IsNotExist(err) {
		// 4. Create directory (including parents if needed)
		if err := os.MkdirAll(destPath, 0755); err != nil {
			return fmt.Errorf("failed to create directory: %v", err)
		}
	} else {
		// Other errors (e.g., permission issues)
		return fmt.Errorf("failed to check path: %v", err)
	}

	cmd := exec.Command("git", "clone", cloneUrl, destPath)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("git clone failed: %v", err)
	}

	fmt.Printf("✅ Successfully cloned %s to %s\n", repo.GetName(), destPath)

	return nil

}
