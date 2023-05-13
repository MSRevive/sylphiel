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
	
	"github.com/msrevive/sylphiel/cmd/dbot"
	"github.com/msrevive/sylphiel/internal/events"
	"github.com/msrevive/sylphiel/internal/commands"

	"github.com/disgoorg/disgo/handler"
	"github.com/saintwish/auralog"
	"github.com/saintwish/auralog/rw"
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

	file := &rw.RotateWriter{
		Dir: dir,
		Filename: filename,
		ExpireTime: ex,
		MaxSize: 5 * rw.Megabyte,
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
	config, err := dbot.LoadConfig(flgs.configFile, flgs.debug)
	if err != nil {
		fmt.Printf("Failed to load config: %s", flgs.configFile)
		return err
	}
	
	fmt.Println("Initiating Loggers...")
	initLoggers("server.log", config.Log.Dir, config.Log.Level, config.Log.ExpireTime)

	//Max threads allowed.
	if config.Core.MaxThreads != 0 {
		runtime.GOMAXPROCS(config.Core.MaxThreads)
	}

	logCore.Println("Initiating Bot...")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	b := dbot.New(ctx, logDisc, config)

	logCore.Println("Configuring bot...")
	if err := b.Setup(
		b.Handler,
		events.OnReady(b),
	); err != nil {
		return err
	}

	logCore.Println("Registering command handlers")
	b.Handler.Command("/ping", commands.HandlePing)
	b.Handler.Command("/restore", commands.HandleRestore)
	b.Handler.Route("/setup", func(cr handler.Router) {
		cr.Command("/roles", commands.HandleRolesSetup(b))
		cr.Command("/serverlist", commands.HandleServerListSetup)
	})

	logCore.Println("Connecting to Discord gateway...")
	if err := b.Start(); err != nil {
		return err
	}

	logCore.Printf("Syncing commands with guild %s", b.Config.Disc.GuildID)
	if _, err := b.Client.Rest().SetGuildCommands(b.Client.ApplicationID(), b.Config.Disc.GuildID, commands.Commands); err != nil {
		logCore.Errorf("Failed to sync commands: %s", err)
	}
	
	defer func() {
		b.Close()
		cancel()
	}()

	fmt.Println("\nBot is now running. Press CTRL-C to exit.\n")
	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGINT, syscall.SIGTERM)
	<-s

	return nil
}