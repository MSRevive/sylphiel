package cmd

import (
	"flag"

	"github.com/saintwish/auralog"
)

var (
	logCore *auralog.Logger // Logs for core/server
	logAPI *auralog.Logger // Logs for endpoints/middleware
)

type flags struct {
	configFile string
	debug bool
}

func doFlags(args []string) *flags {
	flgs := &flags{}

	flagSet := flag.NewFlagSet(args[0], flag.ContinueOnError)
	flagSet.StringVar(&flgs.configFile, "cfile", "./runtime/config.yaml", "Location of via config file")
	flagSet.BoolVar(&flgs.debug, "d", false, "Run with debug mode.")
	flagSet.Parse(args[1:])

	return flgs
}

func initLoggers(filename string, dir string, level string, expire string) {
	ex, _ := time.ParseDuration(expire)
	flags := auralog.Ldate | auralog.Ltime | auralog.Lmicroseconds
	flagsWarn := auralog.Ldate | auralog.Ltime | auralog.Lmicroseconds
	flagsError := auralog.Ldate | auralog.Ltime | auralog.Lmicroseconds | auralog.Lshortfile
	flagsDebug := auralog.Ltime | auralog.Lmicroseconds | auralog.Lshortfile

	file := &auralog.RotateWriter{
		Dir: dir,
		Filename: filename,
		ExTime: ex,
		MaxSize: 5 * auralog.Megabyte,
	}

	logCore = auralog.New(auralog.Config{
		Output: io.MultiWriter(os.Stdout, file),
		Prefix: "[CORE] ",
		Level: auralog.ToLogLevel(level),
		Flag: flags,
		WarnFlag: flagsWarn,
		ErrorFlag: flagsError,
		DebugFlag: flagsDebug,
	})

	logAPI = auralog.New(auralog.Config{
		Output: io.MultiWriter(os.Stdout, file),
		Prefix: "[API] ",
		Level: auralog.ToLogLevel(level),
		Flag: flags,
		WarnFlag: flagsWarn,
		ErrorFlag: flagsError,
		DebugFlag: flagsDebug,
	})
}

func Run(args []string) error {
	flgs := doFlags(args)

	if flgs.debug {
		fmt.Println("Running in Debug mode, do not use in production!")
	}

	fmt.Println("Initiating Loggers...")
	initLoggers("server.log", config.Log.Dir, config.Log.Level, config.Log.ExpireTime)

	//Max threads allowed.
	if config.Core.MaxThreads != 0 {
		runtime.GOMAXPROCS(config.Core.MaxThreads)
	}
}