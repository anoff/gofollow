package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/coreos/pkg/flagutil"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

func main() {
	flags := flag.NewFlagSet("user-auth", flag.ExitOnError)
	consumerKey := flags.String("consumer-key", "", "Twitter Consumer Key")
	consumerSecret := flags.String("consumer-secret", "", "Twitter Consumer Secret")
	accessToken := flags.String("access-token", "", "Twitter Access Token")
	accessSecret := flags.String("access-secret", "", "Twitter Access Secret")
	flags.Parse(os.Args[1:])
	flagutil.SetFlagsFromEnvFile(flags, "TWITTER", ".env")

	if *consumerKey == "" || *consumerSecret == "" || *accessToken == "" || *accessSecret == "" {
		log.Fatal("Consumer key/secret and Access token/secret required")
	}

	config := oauth1.NewConfig(*consumerKey, *consumerSecret)
	token := oauth1.NewToken(*accessToken, *accessSecret)
	// OAuth1 http.Client will automatically authorize Requests
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter client
	client := twitter.NewClient(httpClient)

	// Home Timeline
	// homeTimelineParams := &twitter.HomeTimelineParams{
	// 	Count:     20,
	// 	TweetMode: "extended",
	// }
	// tweets, _, _ := client.Timelines.HomeTimeline(homeTimelineParams)
	// fmt.Println("User's HOME TIMELINE")
	// for _, t := range tweets {
	// 	user := t.User
	// 	if t.RetweetedStatus != nil {
	// 		user = t.RetweetedStatus.User
	// 		fmt.Print("RT: ")
	// 	}
	// 	followRatio := float64(user.FriendsCount) / float64(user.FollowersCount)
	// 	fmt.Println(user.ScreenName, followRatio)
	// }

	// People I follow
	// friendListParams := &twitter.FriendListParams{
	// 	ScreenName: "an0xff",
	// 	Count:      200,
	// 	Cursor:     0,
	// }
	// fmt.Println("Already followed:")
	// for {
	// 	friends, _, _ := client.Friends.List(friendListParams)
	// 	for _, u := range friends.Users {
	// 		followRatio := float64(u.FriendsCount) / float64(u.FollowersCount)
	// 		fmt.Println(u.ScreenName, followRatio)
	// 	}
	// 	cursor := friends.NextCursor
	// 	if cursor == 0 {
	// 		break
	// 	}
	// 	friendListParams.Cursor = cursor
	// }

	// People that already follow me
	followerListParams := &twitter.FollowerListParams{
		ScreenName: "an0xff",
		Count:      200,
		Cursor:     0,
	}
	fmt.Println("Already following:")

	for {
		followers, _, _ := client.Followers.List(followerListParams)
		for _, u := range followers.Users {
			followRatio := float64(u.FriendsCount) / float64(u.FollowersCount)
			fmt.Println(u.ScreenName, followRatio)
		}
		cursor := followers.NextCursor
		if cursor == 0 {
			break
		}
		followerListParams.Cursor = cursor
	}

}
