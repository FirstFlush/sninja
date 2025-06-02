package shell

import (
	"log/slog"
	"os/exec"
	"sninja/internal/types"
	"sninja/internal/utils"
)

func StartShell() error {
	cmd := exec.Command("docker", "exec", "-it", "python", "manage.py", "shell_plus")
	slog.Info("Starting Django shell inside Docker container", "container", types.Container.Web)
	return utils.RunCommand(cmd)
}
