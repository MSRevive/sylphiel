package cmd

import (
	"os"
	"time"
	"runtime"
	"flag"
	"context"
	"fmt"
	"io"
	"syscall"
	"os/signal"

	"github.com/msrevive/sylphiel/internal/system"
	"github.com/msrevive/sylphiel/internal/events"
	"github.com/msrevive/sylphiel/internal/commands"

	"github.com/disgoorg/disgo"
	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/gateway"
	"github.com/disgoorg/disgo/cache"
	"github.com/disgoorg/disgo/handler"
	"github.com/saintwish/auralog"
)

var (
	logCore *auralog.Logger // Logs for core/server
	logDisc *auralog.Logger // Logs for Discord
)

type flags struct {
	configFile string
	debug bool
	syncCommands bool
}

func doFlags(args []string) *flags {
	flgs := &flags{}

	flagSet := flag.NewFlagSet(args[0], flag.ContinueOnError)
	flagSet.StringVar(&flgs.configFile, "cfile", "./runtime/config.yaml", "Location of via config file")
	flagSet.BoolVar(&flgs.debug, "d", false, "Run with debug mode.")
	flagSet.BoolVar(&flgs.syncCommands, "s", false, "Sync commands with all servers.")
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

	logDisc = auralog.New(auralog.Config{
		Output: io.MultiWriter(os.Stdout, file),
		Prefix: "[DISC] ",
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

	fmt.Println("Loading config file...")
	config, err := system.LoadConfig(flgs.configFile, flgs.debug)
	if err != nil {
		return err
	}

	fmt.Println("Initiating Loggers...")
	initLoggers("server.log", config.Log.Dir, config.Log.Level, config.Log.ExpireTime)

	//Max threads allowed.
	if config.Core.MaxThreads != 0 {
		runtime.GOMAXPROCS(config.Core.MaxThreads)
	}

	logCore.Printf("Initiating Disgo (%s)...", disgo.Version)
	client, err := disgo.New(config.Core.Token,
		bot.WithGatewayConfigOpts(
			gateway.WithIntents(
				gateway.IntentGuilds,
				gateway.IntentGuildMessages,
			),
		),
		bot.WithCacheConfigOpts(cache.WithCaches(cache.FlagGuilds)),
		bot.WithEventListenerFunc(events.ApplicationCommands),
		//bot.WithLogger(logDisc),
	)
	if err != nil {
		return err
	}

	logDisc.Print("Registering slash commands...")
	handle := handler.New()
	handle.HandleCommand("/ping", commands.PingHandler)

	if (flgs.syncCommands) {
		logDisc.Print("Syncing global application commads...")
		_,err := client.Rest().SetGlobalCommands(client.ApplicationID(), commands.Commands)

		if (err != nil) {
			logDisc.Errorf("Failed to sync commands: %s", err)
		}
	}

	defer client.Close(context.TODO())
	if err = client.OpenGateway(context.TODO()); err != nil {
		return err
	}

	fmt.Println("\nBot is now running. Press CTRL-C to exit.\n")
	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGINT, syscall.SIGTERM)
	<-s

	return nil
}