package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"time"

	"github.com/docopt/docopt-go"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var version = "1.0.0"

var usage = `nightwatchjs_exporter
Usage:
  nightwatchjs_exporter --nightwatch=<path> --testdir=<path> [options]
  nightwatchjs_exporter --help
  nightwatchjs_exporter --version

Options:
  -n, --nightwatch=<path>  REQUIRED: Path to your nightwatch executable.
  -t, --testdir=<path>     REQUIRED: Directory containing your 'nightwatch.json' file and 'tests' directory.
  --delay=<secs>           Delay in seconds between test executions [default: 30].
  --listen=<host:port>     HTTP listen address [default: :9116].
 
Example:
  nightwatchjs_exporter --nightwatch=/usr/bin/nightwatch --testdir=/home/my_test_dir
`

type Config struct {
	NightwatchjsDir string
	NightwatchjsCmd string
	ListenAddr      string
	DelayTime       time.Duration
}

func main() {
	cfg, err := getConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}
	go start_nightwatch_runner(cfg)
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(cfg.ListenAddr, nil))
}

func getConfig() (Config, error) {
	c := Config{}
	args, err := docopt.Parse(usage, nil, true, version, false)
	if err != nil {
		return c, err
	}
	c.NightwatchjsDir = args["--testdir"].(string)
	c.NightwatchjsCmd = args["--nightwatch"].(string)
	c.ListenAddr = args["--listen"].(string)

	val, err := strconv.Atoi(args["--delay"].(string))
	if err != nil {
		return c, fmt.Errorf("Invalid --delay, expecting a number of seconds")
	}
	c.DelayTime = time.Duration(val) * time.Second

	return c, nil
}

var validationRe = regexp.MustCompile(`[a-zA-Z_][a-zA-Z0-9_]*`)

/*
func getMetricName(name, metric string) string {
	metricName := fmt.Sprintf("module_%s_%s", name, metric)
	if model.IsValidMetricName(model.LabelValue(metricName)) {
		return metricName
	}
	validParts := validationRe.FindAllStringSubmatch(metricName, -1)
	validName := ""
	for _, part := range validParts {
		if len(part) > 0 {
			validName += part[0]
		}
	}
	return validName
}
*/
