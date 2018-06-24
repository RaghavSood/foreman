package supervisor

import (
// "github.com/kolo/xmlrpc"
)

type ProcessInfo struct {
	Name          string `xmlrpc:"name"`
	Description   string `xmlrpc:"description"`
	Group         string `xmlrpc:"group"`
	Start         int64  `xmlrpc:"start"`
	Stop          int64  `xmlrpc:"stop"`
	Now           int64  `xmlrpc:"now"`
	State         int64  `xmlrpc:"state"`
	StateName     string `xmlrpc:"statename"`
	SpawnErr      string `xmlrpc:"spawnerr"`
	ExitStatus    int64  `xmlrpc:"exitstatus"`
	Logfile       string `xmlrpc:"logfile"`
	StdoutLogfile string `xmlrpc:"stdout_logfile"`
	StderrLogfile string `xmlrpc:"stderr_logfile"`
	PID           int64  `xmlrpc:"pid"`
}
