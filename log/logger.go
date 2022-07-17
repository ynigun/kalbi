package log

import "os"
import "github.com/sirupsen/logrus"

//Log global log object
var Log = logrus.New()

func init() {
	Log.Out = os.Stdout
	logrus.SetLevel(logrus.PanicLevel)

}

func SetLevel(Level logrus.Level) {
	logrus.SetLevel(Level)
}
