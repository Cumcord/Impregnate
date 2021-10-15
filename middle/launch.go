package middle

import (
	"fmt"
	"os/exec"
)

func OpenURL(url string) error {
	cmd := exec.Command("xdg-open", url)
	if cmd.Run() == nil {
		return nil
	}
	cmd = exec.Command("cmd", "/c", "start", url)
	if cmd.Run() == nil {
		return nil
	}
	cmd = exec.Command("open", url)
	if cmd.Run() == nil {
		return nil
	}
	return fmt.Errorf("all methods failed")
}
