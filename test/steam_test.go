package test

import (
	"context"
	"fmt"
	"testing"

	steam "github.com/alomerry/steam-web-go-api"
)

func TestGetUserStatsForGame(t *testing.T) {
	ctx := context.Background()
	client := steam.NewClient("")
	resp, err := client.GetIUserStats().GetUserStatsForGame(ctx, &steam.GetUserStatsForGameRequest{
		SteamId: 76561198977254095,
		AppId:   381210,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

func TestGetFriendList(t *testing.T) {
	ctx := context.Background()
	client := steam.NewClient("")
	resp, err := client.GetIUser().GetFriendList(ctx, &steam.GetFriendListRequest{
		SteamId: 76561198977254095,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
