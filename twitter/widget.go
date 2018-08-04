package twitter

import (
	"fmt"
	"os"

	"github.com/senorprogrammer/wtf/cfg"

	"github.com/senorprogrammer/cfg"
	"github.com/senorprogrammer/wtf/twitter/bearertoken"
	"github.com/senorprogrammer/wtf/wtf"
)

type Widget struct {
	wtf.TextWidget
}

func NewWidget() *Widget {
	widget := Widget{
		TextWidget: wtf.NewTextWidget("Twitter", "twitter", false),
	}

	widget.View.SetWrap(true)
	widget.View.SetWordWrap(true)

	setAPICredentials()

	return &widget
}

/* -------------------- Exported Functions -------------------- */

func (widget *Widget) Refresh() {
	client := NewClient("https://api.twitter.com/1.1/")
	userTweets := client.Tweets()

	widget.UpdateRefreshedAt()
	widget.View.SetTitle(widget.ContextualTitle(widget.Name))

	widget.View.SetText(widget.contentFrom(userTweets))
}

/* -------------------- Unexported Functions -------------------- */

func setAPICredentials() {
	token := wtf.Config.UString(
		"wtf.mods.twitter.bearerToken",
		os.Getenv("WTF_TWITTER_BEARER_TOKEN"),
	)
	if len(token) > 0 {
		return
	} else {
		token := bearertoken.NewClient().BearerToken()
		filePath := cfg.ConfigDir()
		config := cfg.LoadConfigFile(filePath)
	}
}

func (widget *Widget) contentFrom(tweets []Tweet) string {
	if len(tweets) == 0 {
		return fmt.Sprintf("\n\n\n%s", wtf.CenterText("[blue]No Tweets[white]", 50))
	}

	str := ""
	for _, tweet := range tweets {
		str = str + widget.format(tweet)
	}

	return str
}

func (widget *Widget) format(tweet Tweet) string {
	return fmt.Sprintf(" [white]%s[grey]\n [grey]%s - %s[white]\n\n", tweet.Text, tweet.Username(), tweet.CreatedAt)
}
