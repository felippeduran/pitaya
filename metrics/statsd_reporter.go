// Copyright (c) TFG Co. All Rights Reserved.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package metrics

import (
	"fmt"
	"strconv"

	"github.com/DataDog/datadog-go/statsd"
	"github.com/felippeduran/pitaya/v2/config"
	"github.com/felippeduran/pitaya/v2/logger"
)

// Client is the interface to required dogstatsd functions
type Client interface {
	Count(name string, value int64, tags []string, rate float64) error
	Gauge(name string, value float64, tags []string, rate float64) error
	TimeInMilliseconds(name string, value float64, tags []string, rate float64) error
}

// StatsdReporter sends application metrics to statsd
type StatsdReporter struct {
	client      Client
	rate        float64
	serverType  string
	defaultTags []string
}

// StatsdConfig provides configuration for statsd
type StatsdConfig struct {
	Host        string
	Prefix      string
	Rate        float64
	ConstLabels map[string]string
}

// NewDefaultStatsdConfig provides default configuration for statsd
func NewDefaultStatsdConfig() StatsdConfig {
	return StatsdConfig{
		Host:        "localhost:9125",
		Prefix:      "pitaya.",
		Rate:        1,
		ConstLabels: map[string]string{},
	}
}

// NewStatsdConfig reads from config to build configuration for statsd
func NewStatsdConfig(config *config.Config) StatsdConfig {
	rate, err := strconv.ParseFloat(config.GetString("pitaya.metrics.statsd.rate"), 64)
	if err != nil {
		panic(err.Error())
	}

	statsdConfig := StatsdConfig{
		Host:        config.GetString("pitaya.metrics.statsd.host"),
		Prefix:      config.GetString("pitaya.metrics.statsd.prefix"),
		Rate:        rate,
		ConstLabels: config.GetStringMapString("pitaya.metrics.constTags"),
	}

	return statsdConfig
}

// NewStatsdReporter returns an instance of statsd reportar and an
// error if something fails
func NewStatsdReporter(
	config StatsdConfig,
	serverType string,
	clientOrNil ...Client,
) (*StatsdReporter, error) {
	return newStatsdReporter(config, serverType, clientOrNil...)
}

func newStatsdReporter(
	config StatsdConfig,
	serverType string,
	clientOrNil ...Client) (*StatsdReporter, error) {
	sr := &StatsdReporter{
		rate:       config.Rate,
		serverType: serverType,
	}

	sr.buildDefaultTags(config.ConstLabels)

	if len(clientOrNil) > 0 {
		sr.client = clientOrNil[0]
	} else {
		c, err := statsd.New(config.Host)
		if err != nil {
			return nil, err
		}
		c.Namespace = config.Prefix
		sr.client = c
	}
	return sr, nil
}

func (s *StatsdReporter) buildDefaultTags(tagsMap map[string]string) {
	defaultTags := make([]string, len(tagsMap)+1)

	defaultTags[0] = fmt.Sprintf("serverType:%s", s.serverType)

	idx := 1
	for k, v := range tagsMap {
		defaultTags[idx] = fmt.Sprintf("%s:%s", k, v)
		idx++
	}

	s.defaultTags = defaultTags
}

// ReportCount sends count reports to statsd
func (s *StatsdReporter) ReportCount(metric string, tagsMap map[string]string, count float64) error {
	fullTags := s.defaultTags

	for k, v := range tagsMap {
		fullTags = append(fullTags, fmt.Sprintf("%s:%s", k, v))
	}

	err := s.client.Count(metric, int64(count), fullTags, s.rate)
	if err != nil {
		logger.Log.Errorf("failed to report count: %q", err)
	}

	return err
}

// ReportGauge sents the gauge value and reports to statsd
func (s *StatsdReporter) ReportGauge(metric string, tagsMap map[string]string, value float64) error {
	fullTags := s.defaultTags

	for k, v := range tagsMap {
		fullTags = append(fullTags, fmt.Sprintf("%s:%s", k, v))
	}

	err := s.client.Gauge(metric, value, fullTags, s.rate)
	if err != nil {
		logger.Log.Errorf("failed to report gauge: %q", err)
	}

	return err
}

// ReportSummary observes the summary value and reports to statsd
func (s *StatsdReporter) ReportSummary(metric string, tagsMap map[string]string, value float64) error {
	fullTags := s.defaultTags

	for k, v := range tagsMap {
		fullTags = append(fullTags, fmt.Sprintf("%s:%s", k, v))
	}

	err := s.client.TimeInMilliseconds(metric, float64(value), fullTags, s.rate)
	if err != nil {
		logger.Log.Errorf("failed to report summary: %q", err)
	}

	return err
}
