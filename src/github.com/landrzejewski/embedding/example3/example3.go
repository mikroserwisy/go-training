package main

import "fmt"

type notifier interface {
	notify()
}

type user struct {
	name  string
	email string
}

// notify implements a method notifies users
// of different events.
func (u *user) notify() {
	fmt.Printf("Sending user email To %s<%s>\n",
		u.name,
		u.email)
}

type admin struct {
	user
	level string
}

func main() {
	ad := admin{
		user: user{
			name:  "john smith",
			email: "john@yahoo.com",
		},
		level: "super",
	}

	// Send the admin user a notification.
	// The embedded inner type's implementation of the
	// interface is "promoted" to the outer type.
	sendNotification(&ad)
}

func sendNotification(n notifier) {
	n.notify()
}
