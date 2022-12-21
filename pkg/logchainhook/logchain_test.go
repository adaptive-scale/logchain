package logchainhook_test

import (
	"github.com/debarshibasak/logchain/pkg/logchainhook"
	"github.com/sirupsen/logrus"
	"testing"
)

func Test(t *testing.T) {

	hook := logchainhook.NewLogChainHook("testapp", "localhost:9090",
		logrus.DebugLevel, logrus.ErrorLevel)
	logrus.AddHook(hook)
	logrus.WithField("test", "Test").Error("test123")
}
