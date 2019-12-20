# nightwatchjs_exporter

nightwatchjs_exporter runs [nightwatchjs](https://nightwatchjs.org/) tests periodically and exports the results for [prometheus.io](https://prometheus.io/).

... because end-to-end testing is important, too!

## Running this software

**You must copy the `nightwatch_json_reporter.js` file from this repo to the working directory nightwatchjs_exporter will run in.**  This file is necessary to format nightwatch.js results such that they are easily parsed by nightwatchjs_exporter.

### From Binaries
Download the most suitable binary from [the releases tab](../../releases).

Then:

    ./nightwatchjs_exporter <flags>

### Building From Source

    go build

You will need to resolve missing dependencies with `go get`.

### Checking the results

Visiting [http://localhost:9355/metrics](http://localhost:9355/metrics)
will return metrics for each your nightwatch.js tests.

## Configuration

**nightwatchjs_exporter requires nightwatch.js! First, please ensure you have a working [nightwatch.js](https://nightwatchjs.org/gettingstarted) installation that can successfully run tests.**


### Required command-line flags
```shell
  -n, --nightwatch=<path>  REQUIRED: Path to your nightwatch executable.
  -t, --testdir=<path>     REQUIRED: Directory containing your 'nightwatch.json' file and 'tests' directory.
```

### Optional command-line flags
The full nightwatchjs_exporter usage is:

```shell
Usage:
  nightwatchjs_exporter --nightwatch=<path> --testdir=<path> [options]
  nightwatchjs_exporter --help
  nightwatchjs_exporter --version

Options:
  -n, --nightwatch=<path>  REQUIRED: Path to your nightwatch executable.
  -t, --testdir=<path>     REQUIRED: Directory containing your 'nightwatch.json' file and 'tests' directory.
  --delay=<secs>           Delay in seconds between test executions [default: 30].
  --listen=<host:port>     HTTP listen address [default: :9355].

Example:
  nightwatchjs_exporter --nightwatch=/usr/bin/nightwatch --testdir=/home/my_test_dir
```
## Prometheus Configuration
nightwatchjs_exporter acts as a standard Prometheus target with no special configuration:
```yaml
  - job_name: 'nightwatchjs'
    scrape_interval: 30s
    static_configs:
    - targets: ['localhost:9355']
```

## Contributing

**Something bugging you?** Please open an Issue or Pull Request - we're here to help!

**New Feature Ideas?** Please open Pull Request, or consider one of these ideas:
* Support nightwatch `--env` flag.
* Support multiple nightwatch configs/test directories.
* Support parallel test execution.
 
**All Humans Are Equal In This Project And Will Be Treated With Respect.**
