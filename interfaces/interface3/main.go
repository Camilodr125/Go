package main

func (e email) cost() int {
	if e.isSubscribed {
		cost:= 2 * len(e.body)
		return cost
	} else {
		cost:= 5 * len(e.body)
		return cost
	}
}

func (e email) format() string {
	if e.isSubscribed {
		message:= "'"+e.body+"'" + " | " + "Subscribed"
		return message
	} else {
		message:= "'"+e.body+"'" + " | " + "Not Subscribed"
		return message
	}
}
type expense interface {
	cost() int
}

type formatter interface {
	format() string
}

type email struct {
	isSubscribed bool
	body         string
}
