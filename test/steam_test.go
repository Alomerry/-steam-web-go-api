package test

import (
	"context"
	"fmt"
	steam "github.com/Alomerry/steam-web-go-api"
	"testing"
)

func TestGetUserStatsForGame(t *testing.T) {
	ctx := context.Background()
	client := steam.NewClient("F4329C307EF6193F7B04D7F8F6E054B0")
	resp, err := client.GetIUserStats().GetUserStatsForGame(ctx, &steam.GetUserStatsForGameRequest{
		SteamId: 76561198977254095,
		AppId:   381210,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
