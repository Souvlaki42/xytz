//go:build !windows

package utils

import (
	"log"
	"syscall"

	"github.com/xdagiz/xytz/internal/types"

	tea "github.com/charmbracelet/bubbletea"
)

func PauseDownload() tea.Cmd {
	return tea.Cmd(func() tea.Msg {
		downloadMutex.Lock()
		defer downloadMutex.Unlock()

		if currentCmd != nil && currentCmd.Process != nil && !isPaused {
			isPaused = true
			if err := currentCmd.Process.Signal(syscall.SIGSTOP); err != nil {
				log.Printf("Failed to pause download: %v", err)
			}
		}

		return types.PauseDownloadMsg{}
	})
}

func ResumeDownload() tea.Cmd {
	return tea.Cmd(func() tea.Msg {
		downloadMutex.Lock()
		defer downloadMutex.Unlock()

		if currentCmd != nil && currentCmd.Process != nil && isPaused {
			isPaused = false
			if err := currentCmd.Process.Signal(syscall.SIGCONT); err != nil {
				log.Printf("Failed to resume download: %v", err)
			}
		}

		return types.ResumeDownloadMsg{}
	})
}
