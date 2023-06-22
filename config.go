package main

import (
	"os"
	"time"

	"github.com/prometheus/alertmanager/template"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Interval time.Duration
	Notify   *Notify
	Evaluate *Evaluate
}

type Notify struct {
	Pagerduty *Pagerduty
}

type Pagerduty struct {
	Key string
}

type EvaluateType string

const (
	EvaluateEqual   EvaluateType = "equal"
	EvaluateInclude EvaluateType = "include"
)

type Evaluate struct {
	Data template.Data
	Type EvaluateType
}

func ParseConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	config := &Config{}
	err = yaml.NewDecoder(file).Decode(config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
