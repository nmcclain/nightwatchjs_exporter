package main

import (
	"github.com/prometheus/client_golang/prometheus"
)

type Metrics struct {
	Passed             prometheus.Gauge
	Failed             prometheus.Gauge
	Errors             prometheus.Gauge
	Skipped            prometheus.Gauge
	Tests              prometheus.Gauge
	Assertions         prometheus.Gauge
	ModuleAssertions   *prometheus.GaugeVec
	ModuleTesttime     *prometheus.GaugeVec
	ModuleTestscount   *prometheus.GaugeVec
	ModuleSkippedcount *prometheus.GaugeVec
	ModuleFailedcount  *prometheus.GaugeVec
	ModuleErrorscount  *prometheus.GaugeVec
	ModulePassedcount  *prometheus.GaugeVec
	ModuleTests        *prometheus.GaugeVec
	ModuleFailures     *prometheus.GaugeVec
	ModuleErrors       *prometheus.GaugeVec
}

func setupProm() Metrics {
	m := Metrics{}
	m.Passed = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "nightwatchjs",
		Name:      "passed",
		Help:      "The number of tests that passed.",
	})
	m.Failed = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "nightwatchjs",
		Name:      "failed",
		Help:      "The number of tests that failed.",
	})
	m.Errors = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "nightwatchjs",
		Name:      "errors",
		Help:      "The number of test errors.",
	})
	m.Skipped = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "nightwatchjs",
		Name:      "skipped",
		Help:      "The number of tests that were skipped.",
	})
	m.Tests = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "nightwatchjs",
		Name:      "tests",
		Help:      "The total number of tests.",
	})
	m.Assertions = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "nightwatchjs",
		Name:      "assertions",
		Help:      "The total number of assertions.",
	})

	m.ModuleAssertions = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "nightwatchjs",
		Name:      "module_assertions",
		Help:      "The number of assertions in this module.",
	}, []string{"module"})
	m.ModuleTesttime = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "nightwatchjs",
		Name:      "module_testtime",
		Help:      "The test time for this module.",
	}, []string{"module"})
	m.ModuleTestscount = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "nightwatchjs",
		Name:      "module_testscount",
		Help:      "The count of tests in this module.",
	}, []string{"module"})
	m.ModuleSkippedcount = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "nightwatchjs",
		Name:      "module_skippedcount",
		Help:      "The number of skipped tests in this module.",
	}, []string{"module"})
	m.ModuleFailedcount = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "nightwatchjs",
		Name:      "module_failedcount",
		Help:      "The number of failed tests in this module.",
	}, []string{"module"})
	m.ModuleErrorscount = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "nightwatchjs",
		Name:      "module_errorscount",
		Help:      "The total of errors in this module.",
	}, []string{"module"})
	m.ModulePassedcount = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "nightwatchjs",
		Name:      "module_passedcount",
		Help:      "The number of passed tests in this module.",
	}, []string{"module"})
	m.ModuleTests = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "nightwatchjs",
		Name:      "module_tests",
		Help:      "The number of tests in this module.",
	}, []string{"module"})
	m.ModuleFailures = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "nightwatchjs",
		Name:      "module_failures",
		Help:      "The number of test failures in this module.",
	}, []string{"module"})
	m.ModuleErrors = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "nightwatchjs",
		Name:      "module_errors",
		Help:      "The number of errors in this module.",
	}, []string{"module"})

	prometheus.MustRegister(
		m.Passed,
		m.Failed,
		m.Errors,
		m.Skipped,
		m.Tests,
		m.Assertions,
		m.ModuleAssertions,
		m.ModuleTesttime,
		m.ModuleTestscount,
		m.ModuleSkippedcount,
		m.ModuleFailedcount,
		m.ModuleErrorscount,
		m.ModulePassedcount,
		m.ModuleTests,
		m.ModuleFailures,
		m.ModuleErrors,
	)

	return m
}
