package git

import (
	"os"
	"os/exec"
)

// Запуск команды git log и возвращение вывода
func RunGitLog(repoPath string) ([]byte, error) {
	err := os.Chdir(repoPath)
	if err != nil {
		return nil, err
	}

	cmd := exec.Command("git", "log", "--pretty=format:%H|%an|%ad|%s", "--date=iso")
	return cmd.Output()
}

// Запуск команды для получения статистики по авторам
func RunGitAuthorStats(repoPath string) ([]byte, error) {
	cmd := exec.Command("git", "-C", repoPath, "log", "--pretty=format:%an")
	return cmd.Output()
}
func RunGitReadmeOnly(repoPath string) ([]byte, error) {
	err := os.Chdir(repoPath)
	if err != nil {
		return nil, err
	}
	cmd := exec.Command("git", "log", "--pretty=format:%H")
	return cmd.Output()
}
func RunGitSilent(repoPath string) ([]byte, error) {
	err := os.Chdir(repoPath)
	if err != nil {
		return nil, err
	}
	cmd := exec.Command("git", "log", "--pretty=format:%H")
	return cmd.Output()
}
