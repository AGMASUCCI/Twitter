package service

import (
	"errors"
	"github.com/Twitter/src/domain"
)

var twt *domain.Tweet



type TweetManager struct {
	twts []*domain.Tweet

}

func (tm *TweetManager)PublishTweet(s *domain.Tweet, usrman *UsuarioManager) (int,error) {

	if s.Text == "" {
		return 0, errors.New("text is required")
	}

	if len(s.Text) > 140 {
		return 0, errors.New("text exceeds 140 characters")
	}

	if s.User.Nombre=="" {
		return 0, errors.New("user is required")
	}

	if usrman.VerifyUser(s.User.Nombre)==false {
		return 0, errors.New("user does not exist")
	}

	twt = s
	tm.twts = append(tm.twts, twt)

	return s.Id, nil
}


func (tm *TweetManager)GetTweet() *domain.Tweet {

	return twt;
}



func (tm *TweetManager)GetTweets() []*domain.Tweet{

	return tm.twts;
}

func (tm *TweetManager)InitializeService(){

	tm.twts = make([]*domain.Tweet, 0)

}

func (tm *TweetManager)GetTweetById(id int) *domain.Tweet {
	for _, value := range tm.twts {
		if value.Id == id {
			return value
		}
	}

	return nil
}

func (tm *TweetManager)CountTweetsByUser(usr string) int {
	var cant int

	for _, value := range tm.twts {
		if value.User.Nombre == usr {
			cant++
		}
	}

	return cant
}

func (tm *TweetManager)GetTweetsByUser(user string) []*domain.Tweet {
	var result map[string][]*domain.Tweet
	result = make(map[string][]*domain.Tweet)
	for _, value := range tm.twts {
		result[value.User.Nombre] = append(result[value.User.Nombre], value)
	}

	return result[user]
}

func NewTweetManager() *TweetManager{

	tweetManager := TweetManager{make([]*domain.Tweet, 0)}
	return &tweetManager

}

