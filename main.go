//		application.NewRouter(backend),
//	}
//
//	srv.InitRouter(routers...)
//
//	go func() {
//		srv.ListenAndServe()
//	}()
//
//	<-sched.Start()
//}

package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/boltdb/bolt"
	. "github.com/memlis/boat/store/local"

	"github.com/Sirupsen/logrus"
	"github.com/urfave/cli"
	"golang.org/x/net/context"
)

func setupLogger(debug bool) {
	if debug {
		logrus.SetLevel(logrus.DebugLevel)
	}

	logrus.SetOutput(os.Stderr)

	logrus.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		FullTimestamp:   true,
	})
}

// waitForSignals wait for signals and do some clean up job.
func waitForSignals(unixSock string) {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	for sig := range signals {
		logrus.Debugf("Received signal %s , clean up...", sig)
		if _, err := os.Stat(unixSock); err == nil {
			logrus.Debugf("Remove %s", unixSock)
			if err := os.Remove(unixSock); err != nil {
				logrus.Errorf("Remove %s failed: %s", unixSock, err.Error())
			}
		}

		os.Exit(0)
	}
}

func main() {
	boat := cli.NewApp()
	boat.Name = "boat"
	boat.Usage = "A small mesos scheduler framework"
	boat.Version = version.Version

	boat.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "addr",
			Value: "0.0.0.0:9999",
			Usage: "API Server address <ip:port>",
		},
		cli.StringFlag{
			Name:  "sock",
			Value: "/var/run/swan.sock",
			Usage: "Unix socket for listening",
		},
		cli.StringFlag{
			Name:  "master",
			Value: "127.0.0.1:5050",
			Usage: "master address host1:port1,host2:port2,... or zk://host1:port1,host2:port2,.../path",
		},
		cli.StringFlag{
			Name:  "user",
			Value: "root",
			Usage: "mesos framework user",
		},
		cli.BoolFlag{
			Name:  "debug",
			Usage: "enable debug log",
		},
	}
	app.Action = func(c *cli.Context) error {
		setupLogger(c.Bool("debug"))

		store, err := NewBoltStore("bolt.db")
		if err != nil {
			logrus.Errorf("Init store engine failed:%s", err)
			return
		}

		go func() {
			waitForSignals(c.String("sock"))
		}()

		scheduler := NewMesosScheduler(config)

		return scheduler.Run()
	}

	if err := boat.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
