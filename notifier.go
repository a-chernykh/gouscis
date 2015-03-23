package main

type notifier interface {
	Notify(text string)
}
