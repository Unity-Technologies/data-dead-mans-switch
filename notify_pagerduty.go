package main

import (
	"fmt"
	"log"
	"time"

	pagerduty "github.com/PagerDuty/go-pagerduty"
)

var _ NotifyInterface = new(PagerDuty)

type PagerDuty struct {
	AuthKey string
	Source  string
	Details map[string]interface{}
}

func NewPagerDutyNotify(authKey string, source string, details map[string]interface{}) *PagerDuty {
	return &PagerDuty{AuthKey: authKey, Source: source, Details: details}
}

// Notify send notify message to pagerduty
func (p *PagerDuty) Notify(summary string) error {
	log.Printf("sending notify: %s to pagerduty\n", summary)
	pdPayload := pagerduty.V2Payload{
		Summary:   fmt.Sprintf("%s %s", p.Source, summary),
		Source:    p.Source,
		Severity:  "critical",
		Timestamp: time.Now().Format(time.RFC3339),
		Details:   p.Details,
		Group:     "DeadMansSwitch",
		// used for group alerting event
		Class: summary,
	}
	event := pagerduty.V2Event{
		RoutingKey: p.AuthKey,
		Action:     "trigger",
		Client:     "DeadMansSwitch",
		Payload:    &pdPayload,
	}

	_, err := pagerduty.ManageEvent(event)
	return err
}
