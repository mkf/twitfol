package main

import (
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	//	"github.com/skratchdot/open-golang/open"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func getTrimmedStringFromDotDatFileFromDotTwitfol(name string) (trimmedString string) {
	path_home_twitfol := os.Getenv("HOME") + "/.twitfol/"
	path := path_home_twitfol + name + ".dat"
	bytesUntrimmed, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Got " + path)
	trimmedString = strings.TrimSpace(string(bytesUntrimmed))
	return
}

func main() {
	accessToken := getTrimmedStringFromDotDatFileFromDotTwitfol("twitteraccesstoken")
	accessSecret := getTrimmedStringFromDotDatFileFromDotTwitfol("twitteraccesssecret")
	consumerKey := getTrimmedStringFromDotDatFileFromDotTwitfol("twitterconsumerkey")
	consumerSecret := getTrimmedStringFromDotDatFileFromDotTwitfol("twitterconsumersecret")
	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessSecret)
	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)
	log.Println("Twitter client creation complete.")
	var cursor int64 = -1
	var pages uint = 1
	var sofar uint = 0
	var skipstatus = true
	var entities = false
	const desiredCount = 200
	for {
		if cursor == -1 {
			log.Println("Starting...")
		}
		followers, _, err := client.Followers.List(&twitter.FollowerListParams{
			Cursor:              cursor,
			Count:               desiredCount,
			SkipStatus:          &skipstatus,
			IncludeUserEntities: &entities,
		})
		if err != nil {
			log.Fatalln("Error while getting page number ", pages, " of cursor ", cursor, " : ", err)
		}
		var thispage uint = 0
		for _, follo := range followers.Users {
			fmt.Printf("%s\t%s\n", follo.ScreenName, follo.IDStr)
			thispage += 1
		}
		sofar += thispage
		log.Printf("Finished page number %d of cursor %d containing %d entries. So far %d entries.",
			pages, cursor, thispage, sofar)
		cursor = followers.NextCursor
		if cursor != 0 {
			pages += 1
			continue
		}
		log.Println("Reached end of paging.")
		break
	}
	//log.Println(followers)
	log.Println("Fin.")
}
