package lib

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"strings"
	"time"
)

const ErrorMessageCommandDeadline = "Deadline exceeded"
const DefaultTimeDurationExecuteCommand time.Duration = time.Second * 10

const (
	SystemCtrl = "systemctl"
	SystemCtrlServiceDir = "/lib/systemd/system"
	SystemCtrlActionStop    = "stop"
	SystemCtrlActionStart   = "start"
	SystemCtrlActionRestart = "restart"
	SystemCtrlActionReload  = "reload"
	SystemCtrlActionEnable  = "enable"
	SystemCtrlActionDisable = "disable"
)

type (
	OptionSystemCtrlCreate struct {
		WorkDir string
		UnitName string
		BinaryName string
		Description string
		PathSave string
		Type string
	}
	OptionSystemCtrlAction struct {
		Action string
		WorkDir string
		UnitName string
		Duration time.Duration
	}
)
func (i *OptionSystemCtrlCreate) prepare () () {
	if len(i.BinaryName) == 0 {
		i.BinaryName = i.UnitName
	}

	if len(i.PathSave) == 0 {
		i.PathSave = path.Join(SystemCtrlServiceDir, i.UnitName + ".service")
	}

	if len(i.Type) == 0 {
		i.Type = "Web server"
	}
}

func (i *OptionSystemCtrlCreate) tplName () (name string) {
	return "TplCreateSystemUnit" + i.UnitName
}

func (i *OptionSystemCtrlCreate) UnitDrop () (err error) {
	i.prepare()

	optStop := OptionSystemCtrlAction{
		Action:   SystemCtrlActionStop,
		WorkDir:  i.WorkDir,
		UnitName: i.UnitName,
	}

	if _, err = SystemCtlServiceAction(&optStop); err != nil {
		return err
	}

	return os.Remove(i.PathSave)
}

func (i *OptionSystemCtrlCreate) UnitCreate () (err error) {
	i.prepare()
	name := i.tplName()

	buffer, err := SlowRenderContent(&name, ServiceTpl , i)

	if err != nil {
		return err
	}

	return ioutil.WriteFile(i.PathSave, buffer.Bytes(), 0644)
}

func (i *OptionSystemCtrlAction) GetDuration () (duration time.Duration) {
	if i.Duration == 0 {
		return DefaultTimeDurationExecuteCommand
	}

	return i.Duration
}

func ExecCommandWithTimeLimit(timeLimit time.Duration, workDir string, base string, args ...string) (out []byte, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeLimit)
	defer cancel()

	cmd := exec.CommandContext(ctx, base, args...)
	cmd.Dir = workDir

	if out, err = cmd.CombinedOutput(); ctx.Err() == context.DeadlineExceeded {
		return out, fmt.Errorf("%v for %v %v", ErrorMessageCommandDeadline, base, strings.Join(args, " "))
	}

	return out, err
}

func SystemCtlServiceAction(opt *OptionSystemCtrlAction) (out []byte, err error) {
	return ExecCommandWithTimeLimit(opt.GetDuration(), opt.WorkDir, SystemCtrl , opt.Action, opt.UnitName)
}

/**
systemctl daemon-reload
systemctl restart deployTools
systemctl status deployTools
*/
const ServiceTpl = `[Unit]
Description={{.Description}}
After=network-online.target
[Service]
Restart=always
StandardOutput=syslog
StandardError=syslog
SyslogIdentifier={{.UnitName}}
WorkingDirectory={{.WorkDir}}
ExecStart={{.WorkDir}}/{{.BinaryName}}
Type={{.Type}}
[Install]
WantedBy=multi-user.target
`
