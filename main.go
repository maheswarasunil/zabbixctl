package main

import (
	"time"

	"github.com/kovetskiy/godocs"
	"github.com/kovetskiy/lorg"
	"github.com/kovetskiy/spinner-go"
	"github.com/zazab/hierr"
)

var (
	debugMode bool
	traceMode bool

	logger = getLogger()
)

func init() {
	spinner.SetInterval(time.Millisecond * 100)
}

func main() {
	args, err := godocs.Parse(
		docs, version,
		godocs.UsePager, godocs.Usage(usage), godocs.Options(options),
	)
	if err != nil {
		fatalln(err)
	}

	switch args["--verbosity"].(int) {
	case 1:
		debugMode = true
		logger.SetLevel(lorg.LevelDebug)
	case 2:
		debugMode = true
		traceMode = true
		logger.SetLevel(lorg.LevelTrace)
	}

	config, err := NewConfig(args["--config"].(string))
	if err != nil {
		fatalln(
			hierr.Errorf(
				err,
				"problem with configuration",
			),
		)
	}

	zabbix, err := NewZabbix(
		config.Server.Address,
		config.Server.Username,
		config.Server.Password,
		config.Server.IgnoreServerCert,
		config.Session.Path,
	)
	if err != nil {
		fatalln(err)
	}

	switch {
	case args["--triggers"].(bool):
		err = handleTriggers(zabbix, config, args)
	case args["--latest-data"].(bool):
		err = handleLatestData(zabbix, config, args)
	case args["--groups"].(bool):
		err = handleUsersGroups(zabbix, config, args)

	}

	if err != nil {
		fatalln(err)
	}
}
