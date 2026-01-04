package structured

import (
	"github.com/sirupsen/logrus"
)

type Hook struct {
	id string
}

func (hook *Hook) Fire(entry *logrus.Entry) error {
	entry.Data["id"] = hook.id
	return nil
}

func (hook *Hook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func Logrus() {
	logrus.SetFormatter(&logrus.TextFormatter{})
	logrus.SetLevel(logrus.InfoLevel)
	logrus.AddHook(&Hook{id: "123"})

	fields := logrus.Fields{}
	fields["success"] = true
	fields["complex_struct"] = struct {
		Event string
		When  string
	}{
		Event: "file_upload",
		When:  "2024-04-27",
	}
	x := logrus.WithFields(fields)
	x.Warn("warning message")
	x.Error("Error!")
}
