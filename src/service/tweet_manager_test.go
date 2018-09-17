package service_test

import (
	"github.com/Twitter/src/domain"
	"github.com/Twitter/src/service"
	"testing"
)

//Los tests siempre arrancan con la palabra Test
func TestPublishedTweetIsSaved(t *testing.T) {

	// Initialization
	tweetManager := service.NewTweetManager()
	usuarioManager := service.NewUsuarioManager()

	var tweet *domain.Tweet

	//user := "grupoesfera"
	text := "This is my first tweet"
	var usuario *domain.Usuario

	nombre := "Gabriel"
	mail := "gabriel@hotmail.com"
	nick := "Gaby"
	clave := "1234"

	usuario = domain.NewUsuario(nombre, mail, nick, clave)
	usuarioManager.AddUser(usuario)

	tweet = domain.NewTweet(usuario, text)

	// Operation
	id, _ := tweetManager.PublishTweet(tweet, usuarioManager)

	// Validation
	publishedTweet := tweetManager.GetTweet()

	isValidTweet(t, publishedTweet, id, usuario.Nombre, text)
}

func TestTweetWithoutUserIsNotPublished(t *testing.T) {

	// Initialization
	tweetManager := service.NewTweetManager()
	usuarioManager := service.NewUsuarioManager()
	usuarioManager.InitializeService()

	var usuario *domain.Usuario

	nombre := ""
	mail := "gabriel@hotmail.com"
	nick := "Gaby"
	clave := "1234"

	usuario = domain.NewUsuario(nombre, mail, nick, clave)
	usuarioManager.AddUser(usuario)

	var tweet *domain.Tweet

	text := "This is my first tweet"

	tweet = domain.NewTweet(usuario, text)

	// Operation
	var err error
	_, err = tweetManager.PublishTweet(tweet,usuarioManager)

	// Validation
	if err == nil {
		t.Error("Expected error")
		return
	}

	if err.Error() != "user is required" {
		t.Error("Expected error is user is required")
	}
}

func TestTweetWithoutTextIsNotPublished(t *testing.T) {

	// Initialization
	tweetManager := service.NewTweetManager()
	usuarioManager := service.NewUsuarioManager()
	usuarioManager.InitializeService()

	var usuario *domain.Usuario

	nombre := "Gabriel"
	mail := "gabriel@hotmail.com"
	nick := "Gaby"
	clave := "1234"

	usuario = domain.NewUsuario(nombre, mail, nick, clave)
	usuarioManager.AddUser(usuario)

	var tweet *domain.Tweet

	var text string

	tweet = domain.NewTweet(usuario, text)

	// Operation
	var err error
	_, err = tweetManager.PublishTweet(tweet,usuarioManager)

	// Validation
	if err == nil {
		t.Error("Expected error")
		return
	}

	if err.Error() != "text is required" {
		t.Error("Expected error is text is required")
	}
}

func TestTweetWhichExceeding140CharactersIsNotPublished(t *testing.T) {

	// Initialization
	tweetManager := service.NewTweetManager()
	usuarioManager := service.NewUsuarioManager()
	usuarioManager.InitializeService()

	var usuario *domain.Usuario

	nombre := "Gabriel"
	mail := "gabriel@hotmail.com"
	nick := "Gaby"
	clave := "1234"

	usuario = domain.NewUsuario(nombre, mail, nick, clave)
	usuarioManager.AddUser(usuario)

	var tweet *domain.Tweet

	text := `The Go project has grown considerably with over half a million users and community members 
	all over the world. To date all community oriented activities have been organized by the community
	with minimal involvement from the Go project. We greatly appreciate these efforts`

	tweet = domain.NewTweet(usuario, text)

	// Operation
	var err error
	_, err = tweetManager.PublishTweet(tweet,usuarioManager)

	// Validation
	if err == nil {
		t.Error("Expected error")
		return
	}

	if err.Error() != "text exceeds 140 characters" {
		t.Error("Expected error is text exceeds 140 characters")
	}
}

func TestCanPublishAndRetrieveMoreThanOneTweet(t *testing.T) {

	// Initialization
	tweetManager := service.NewTweetManager()
	usuarioManager := service.NewUsuarioManager()
	usuarioManager.InitializeService()

	var usuario *domain.Usuario

	nombre := "Gabriel"
	mail := "gabriel@hotmail.com"
	nick := "Gaby"
	clave := "1234"

	usuario = domain.NewUsuario(nombre, mail, nick, clave)
	usuarioManager.AddUser(usuario)

	var tweet, secondTweet *domain.Tweet

	text := "This is my first tweet"
	secondText := "This is my second tweet"

	tweet = domain.NewTweet(usuario, text)
	secondTweet = domain.NewTweet(usuario, secondText)

	// Operation
	firstId, _ := tweetManager.PublishTweet(tweet,usuarioManager)
	secondId, _ := tweetManager.PublishTweet(secondTweet,usuarioManager)

	// Validation
	publishedTweets := tweetManager.GetTweets()

	if len(publishedTweets) != 2 {

		t.Errorf("Expected size is 2 but was %d", len(publishedTweets))
		return
	}

	firstPublishedTweet := publishedTweets[0]
	secondPublishedTweet := publishedTweets[1]

	if !isValidTweet(t, firstPublishedTweet, firstId, usuario.Nombre, text) {
		return
	}

	if !isValidTweet(t, secondPublishedTweet, secondId, usuario.Nombre, secondText) {
		return
	}

}

func TestCanRetrieveTweetById(t *testing.T) {

	// Initialization
	tweetManager := service.NewTweetManager()
	usuarioManager := service.NewUsuarioManager()
	usuarioManager.InitializeService()

	var usuario *domain.Usuario

	nombre := "Gabriel"
	mail := "gabriel@hotmail.com"
	nick := "Gaby"
	clave := "1234"

	usuario = domain.NewUsuario(nombre, mail, nick, clave)
	usuarioManager.AddUser(usuario)

	var tweet *domain.Tweet
	var id int

	text := "This is my first tweet"

	tweet = domain.NewTweet(usuario, text)

	// Operation
	id, _ = tweetManager.PublishTweet(tweet,usuarioManager)

	// Validation
	publishedTweet := tweetManager.GetTweetById(id)

	isValidTweet(t, publishedTweet, id, usuario.Nombre, text)
}

func TestCanCountTheTweetsSentByAnUser(t *testing.T) {

	// Initialization
	tweetManager := service.NewTweetManager()
	usuarioManager := service.NewUsuarioManager()
	usuarioManager.InitializeService()

	var usuario , otroUsuario *domain.Usuario

	nombre := "Gabriel"
	mail := "gabriel@hotmail.com"
	nick := "Gaby"
	clave := "1234"

	usuario = domain.NewUsuario(nombre, mail, nick, clave)
	usuarioManager.AddUser(usuario)

	nombre = "Lucas"
	otroUsuario = domain.NewUsuario(nombre, mail, nick, clave)
	usuarioManager.AddUser(otroUsuario)

	var tweet, secondTweet, thirdTweet *domain.Tweet

	text := "This is my first tweet"
	secondText := "This is my second tweet"

	tweet = domain.NewTweet(usuario, text)
	secondTweet = domain.NewTweet(usuario, secondText)
	thirdTweet = domain.NewTweet(otroUsuario, text)

	tweetManager.PublishTweet(tweet,usuarioManager)
	tweetManager.PublishTweet(secondTweet,usuarioManager)
	tweetManager.PublishTweet(thirdTweet,usuarioManager)

	// Operation
	count := tweetManager.CountTweetsByUser(usuario.Nombre)

	// Validation
	if count != 2 {
		t.Errorf("Expected count is 2 but was %d", count)
	}

}

func TestCanRetrieveTheTweetsSentByAnUser(t *testing.T) {

	// Initialization
	tweetManager := service.NewTweetManager()
	usuarioManager := service.NewUsuarioManager()
	usuarioManager.InitializeService()

	var usuario , otroUsuario *domain.Usuario

	nombre := "Gabriel"
	mail := "gabriel@hotmail.com"
	nick := "Gaby"
	clave := "1234"

	usuario = domain.NewUsuario(nombre, mail, nick, clave)
	usuarioManager.AddUser(usuario)

	nombre = "Lucas"
	otroUsuario = domain.NewUsuario(nombre, mail, nick, clave)
	usuarioManager.AddUser(otroUsuario)

	var tweet, secondTweet, thirdTweet *domain.Tweet

	text := "This is my first tweet"
	secondText := "This is my second tweet"

	tweet = domain.NewTweet(usuario, text)
	secondTweet = domain.NewTweet(usuario, secondText)
	thirdTweet = domain.NewTweet(otroUsuario, text)

	firstId, _ := tweetManager.PublishTweet(tweet,usuarioManager)
	secondId, _ := tweetManager.PublishTweet(secondTweet,usuarioManager)
	tweetManager.PublishTweet(thirdTweet,usuarioManager)

	// Operation
	tweets := tweetManager.GetTweetsByUser(usuario.Nombre)

	// Validation
	if len(tweets) != 2 {

		t.Errorf("Expected size is 2 but was %d", len(tweets))
		return
	}

	firstPublishedTweet := tweets[0]
	secondPublishedTweet := tweets[1]

	if !isValidTweet(t, firstPublishedTweet, firstId, usuario.Nombre, text) {
		return
	}

	if !isValidTweet(t, secondPublishedTweet, secondId, usuario.Nombre, secondText) {
		return
	}

}

func TestTweetWithoutUserRegistered(t *testing.T) {

	// Initialization
	tweetManager := service.NewTweetManager()
	usuarioManager := service.NewUsuarioManager()
	usuarioManager.InitializeService()

	var usuario *domain.Usuario

	nombre := "Gabriel"
	mail := "gabriel@hotmail.com"
	nick := "Gaby"
	clave := "1234"

	usuario = domain.NewUsuario(nombre, mail, nick, clave)
	//usuarioManager.AddUser(usuario)  --Se comenta para que pase el test

	var tweet *domain.Tweet

	text := "This is my first tweet"

	tweet = domain.NewTweet(usuario, text)

	// Operation
	var err error
	_, err = tweetManager.PublishTweet(tweet,usuarioManager)

	if err == nil {
		t.Error("Expected error")
		return
	}

	if err.Error() != "user does not exist" {
		t.Error("Expected error is user does not exist")
	}
}

func isValidTweet(t *testing.T, tweet *domain.Tweet, id int, user, text string) bool {

	if tweet.Id != id {
		t.Errorf("Expected id is %v but was %v", id, tweet.Id)
	}

	if tweet.User.Nombre != user && tweet.Text != text {
		t.Errorf("Expected tweet is %s: %s \nbut is %s: %s",
			user, text, tweet.User, tweet.Text)
		return false
	}

	if tweet.Date == nil {
		t.Error("Expected date can't be nil")
		return false
	}

	return true

}
