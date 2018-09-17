package domain

import "time"

type Tweet struct{
	User , Text string
	Date *time.Time
}

func NewTweet(user string, text string) *Tweet {
	var now = time.Now();
	return &Tweet{user,text, &now}
}