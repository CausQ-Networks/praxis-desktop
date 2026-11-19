package main

import (
	"errors"
	"fmt"
	"os"
	"syscall"
	"time"

	"github.com/CausQ-Networks/praxis-newt/logger"
	"github.com/CausQ-Networks/praxis-desktop/config"
	"github.com/CausQ-Networks/praxis-desktop/managers"
	"github.com/CausQ-Networks/praxis-desktop/ui"
	"github.com/CausQ-Networks/praxis-desktop/version"
)

const usage = `Usage: praxis-client [--autostart]
       praxis-client --daemon
       praxis-client --tunnel <config.json>
`

// daemonUnavailableWarning is shown when the UI can't reach the daemon, with
// the command to start the service under this system's init: systemd for the
// packages, OpenRC or neither for the tar.gz's install.sh.
func daemonUnavailableWarning() string {
	start := "sudo systemctl enable --now praxis-client"
	if _, err := os.Stat("/run/systemd/system"); err != nil {
		if _, err := os.Stat("/etc/init.d/praxis-client"); err == nil {
			start = "sudo rc-update add praxis-client default\nsudo rc-service praxis-client start"
		} else {
			start = "sudo praxis-client --daemon"
		}
	}
	return "The Praxis background service isn't running, so Praxis can't connect " +
		"or remember your sign-in.\n\nStart it with:\n\n" + start
}

func main() {
	setupLogging()
	logger.Info("Praxis version %s starting", version.Number)

	args := os.Args[1:]
	switch {
	case len(args) == 1 && args[0] == "--daemon":
		// Run by the praxis-client systemd or OpenRC service.
		if err := managers.Run(); err != nil {
			logger.Fatal("Daemon failed: %v", err)
		}
		return

	case len(args) == 2 && args[0] == "--tunnel":
		// Run by the daemon, one process per tunnel.
		configJSON, err := os.ReadFile(args[1])
		if err != nil {
			logger.Fatal("Failed to read tunnel config: %v", err)
		}
		if err := managers.RunTunnelService(string(configJSON)); err != nil {
			logger.Fatal("Tunnel failed: %v", err)
		}
		return

	case len(args) == 1 && args[0] == "--autostart":
		// Run from the XDG autostart entry at sign-in: only open the UI if
		// the user asked for it, as the Windows service does at logon.
		if !config.LaunchUIAtLoginEnabledIn(config.UserDataDir()) {
			logger.Info("Open at login is off; exiting")
			return
		}

	case len(args) > 0:
		fmt.Fprint(os.Stderr, usage)
		os.Exit(2)
	}

	var warning string
	reader, writer, events, err := managers.AttachUI()
	switch {
	case errors.Is(err, managers.ErrUIAlreadyRunning):
		logger.Info("Praxis is already running; asking it to show itself")
		managers.RequestUIAction(managers.UIActionShow)
		return
	case errors.Is(err, managers.ErrUpdateInstalling):
		logger.Info("The daemon is installing an update; waiting for it")
		if err := ui.RunInstallingUpdate(waitForUpdateThenRestart); err != nil {
			logger.Error("Failed to run UI: %v", err)
			os.Exit(1)
		}
		return
	case err != nil:
		logger.Error("Failed to connect to the Praxis daemon: %v", err)
		warning = daemonUnavailableWarning()
	default:
		managers.InitializeIPCClient(reader, writer, events)
		logger.Info("Connected to the Praxis daemon")
	}

	if err := runUI(warning); err != nil {
		logger.Error("Failed to run UI: %v", err)
		os.Exit(1)
	}
}

// updateInstallTimeout bounds how long a UI waits for the daemon to install
// an update before starting anyway.
const updateInstallTimeout = 15 * time.Minute

// waitForUpdateThenRestart waits for the daemon's update install to finish,
// then starts this program again in place of this process, so the UI that
// opens is the new version.
func waitForUpdateThenRestart() {
	if !managers.WaitForUpdateInstall(updateInstallTimeout) {
		logger.Error("Timed out waiting for the update to install")
	}
	exe, err := os.Executable()
	if err != nil {
		logger.Error("Failed to restart after the update: %v", err)
		return
	}
	logger.Info("Update finished; starting %s again", exe)
	if err := syscall.Exec(exe, os.Args, os.Environ()); err != nil {
		logger.Error("Failed to restart after the update: %v", err)
	}
}
