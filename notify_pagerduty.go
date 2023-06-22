package main

import (
	"log"
	"time"

	pagerduty "github.com/PagerDuty/go-pagerduty"
)

var _ NotifyInterface = new(PagerDuty)

type PagerDuty struct {
	AuthKey string
	Details string
}

func NewPagerDutyNotify(authKey string, details string) *PagerDuty {
	return &PagerDuty{AuthKey: authKey, Details: details}
}

// Notify send notify message to pagerduty
func (p *PagerDuty) Notify(summary, detail string) error {
	log.Printf("sending notify: %s to pagerduty\n", summary)
	pdPayload := pagerduty.V2Payload{
		Summary:   summary,
		Source:    "DeadMansSwitch",
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
