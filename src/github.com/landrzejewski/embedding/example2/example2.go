package main

import "fmt"

type user struct {
	name  string
	email string
}

func (u *user) notify() {
	fmt.Printf("Sending user email To %s<%s>\n",
		u.name,
		u.email)
}

type admin struct {
	user  // Embedded Type
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

	ad.user.notify()

	ad.notify()
}
