package lib

import (
	"context"
	"errors"
	"os/exec"
	"strings"
	"time"
)

func ExecCommandWithTimeLimit(timeLimit time.Duration, base string, args ...string) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeLimit)
	defer cancel()

	out := []byte("")
	cmd := exec.CommandContext(ctx, base, args...)
	out, err := cmd.CombinedOutput()

	if ctx.Err() == context.DeadlineExceeded {
		return out, errors.New("Deadline exceeded for " + base + " " + strings.Join(args, " "))
	}

	return out, err
}
