package shell

import (
	"log/slog"
	"os/exec"
	"sninja/internal/types/docker"
	"sninja/internal/utils"
)

func StartShell() error {
	cmd := exec.Command("docker", "exec", "-it", "street_ninja_web", "python", "manage.py", "shell_plus")
	slog.Info("Starting Django shell inside Docker container", "container", docker.Web)
	return utils.RunCommand(cmd)
}
