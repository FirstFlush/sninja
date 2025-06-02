package cmd

import (
	"log/slog"
	"os"
	"sninja/core/shell"

	"github.com/spf13/cobra"
)

var shellCmd = &cobra.Command{
	Use:   "shell",
	Short: "Open Django shell_plus in the Docker container",
	Long: `Starts an interactive Django shell_plus session inside the running container.

	This wraps 'docker exec -it street_ninja_web python manage.py shell_plus' so you don't have to type it manually`,
	Run: func(cmd *cobra.Command, args []string) {
		err := shell.StartShell()
		if err != nil {
			slog.Error("Failed to run shell command", "error", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(shellCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// shellCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// shellCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
