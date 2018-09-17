package domain

import (
	"fmt"
	"time"
)

var twtSeq int

type Tweet struct{
	User *Usuario
	Text string
	Date *time.Time
	Id int
}

func NewTweet(user *Usuario, text string) *Tweet {
	var now = time.Now();
	twtSeq = twtSeq +1
	return &Tweet{user,text, &now, twtSeq}
}

func (t *Tweet) PrintableTweet() string {
	return fmt.Sprintf("@%s: %s", t.User, t.Text)
}