package service

import (
	"errors"
	"github.com/Twitter/src/domain"
)

var twt *domain.Tweet

func PublishTweet(s *domain.Tweet) error {
	twt = s

	if s.Text == "" {
		return errors.New("text is required");
	}

	if len(s.Text) > 140 {
		return errors.New("text maximium es 140");
	}

	return nil
}

func GetTweet() *domain.Tweet {

	return twt;
}

