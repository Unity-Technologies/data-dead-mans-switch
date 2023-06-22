package main

type NotifyInterface interface {
	Notify(summary string) error
}
