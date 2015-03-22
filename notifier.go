package main

type Notifier interface {
	Notify(text string)
}
