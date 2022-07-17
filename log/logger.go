package log

import "os"
import "github.com/sirupsen/logrus"

//Log global log object
var Log = logrus.New()

func init() {
	Log.Out = os.Stdout
	Log.Level = logrus.PanicLevel

}

func SetLevel(Level logrus.Level) {
		Log.Level = logrus.PanicLevel
}
