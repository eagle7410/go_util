package lib

import (
	"context"
	"fmt"
	"os/exec"
	"strings"
	"time"
)

const ErrorMesssageCommandDeadline = "Deadline exceeded"

func ExecCommandWithTimeLimit(timeLimit time.Duration, workDir string, base string, args ...string) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeLimit)
	defer cancel()

	out := []byte("")
	cmd := exec.CommandContext(ctx, base, args...)
	cmd.Dir = workDir
	out, err := cmd.CombinedOutput()

	if ctx.Err() == context.DeadlineExceeded {
		return out, fmt.Errorf("%v for %v %v", ErrorMesssageCommandDeadline, base, strings.Join(args, " "))
	}

	return out, err
}
