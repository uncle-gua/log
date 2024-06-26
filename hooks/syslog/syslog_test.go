// +build !windows,!nacl,!plan9

package syslog

import (
	"log/syslog"
	"testing"

	"github.com/uncle-gua/log"
)

func TestLocalhostAddAndPrint(t *testing.T) {
	log := log.New()
	hook, err := NewSyslogHook("udp", "localhost:514", syslog.LOG_INFO, "")

	if err != nil {
		t.Errorf("Unable to connect to local syslog.")
	}

	log.Hooks.Add(hook)

	for _, level := range hook.Levels() {
		if len(log.Hooks[level]) != 1 {
			t.Errorf("SyslogHook was not added. The length of log.Hooks[%v]: %v", level, len(log.Hooks[level]))
		}
	}

	log.Info("Congratulations!")
}
