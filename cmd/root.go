// Package cmd is implemented to represent the command line tool to execute the program logic
package cmd

import (
	"fmt"
	"os"

	"github.com/igorlopushko/ignite.homework/api/config"
	"github.com/igorlopushko/ignite.homework/api/service"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	easy "github.com/t-tomalak/logrus-easy-formatter"
)

var aliensCount int
var fenvfile string

var rootCmd = &cobra.Command{
	Use:   "go run main.go --aliens-count [number] --env [path]",
	Short: "Alien invasion game",
	Long:  "Alien invasion game just to demonstrate golang skills",
	RunE:  run,
	PersistentPreRunE: func(_ *cobra.Command, _ []string) error {
		if fenvfile != "" {
			if err := godotenv.Load(fenvfile); err != nil {
				logrus.Error(fmt.Printf("failed to load envfile [%s]", fenvfile))
				return err
			}
		}
		return nil
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().IntVarP(
		&aliensCount,
		"aliens-count",
		"ac",
		0,
		"Number of aliens to invade the planet X")

	rootCmd.PersistentFlags().StringVarP(
		&fenvfile,
		"env",
		"e",
		"",
		"Path to env file")

	cobra.OnInitialize(initConfig)
}

func initConfig() {
	err := godotenv.Load(fenvfile)
	if err != nil {
		panic("could not load config file")
	}
	config.App.Load()

	logrus.SetFormatter(&easy.Formatter{
		TimestampFormat: "2006-01-02 15:04:05",
		LogFormat:       "[%lvl%]: %time% - %msg%\n",
	})

	logLvl, err := logrus.ParseLevel(config.App.LogLevel)
	if err != nil {
		logrus.Warn("could not parse log level, using debug default")
		logLvl = logrus.DebugLevel
	}
	logrus.SetLevel(logLvl)
}

func run(cmd *cobra.Command, _ []string) error {
	// load map from file
	fs := &service.FileLoaderSrv{}
	m, err := fs.Load(config.App.MapFilePath)
	if err != nil {
		logrus.Fatal(err)
		return err
	}

	// create game object
	g := &service.Game{
		Cities:        m,
		AlienSvc:      service.AlienSvc{MaxStepsCount: config.App.AlienMaxStepsNumber},
		NavigationSvc: service.NavigationSvc{},
	}

	// run the game
	err = g.Run(aliensCount)
	if err != nil {
		logrus.Error(err)
	}

	return nil
}
