package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/nlopes/slack"
)

const (
	ENV_SLACK_API_TOKEN = "ENV_SLACK_API_TOKEN"
)

func main() {
	userGroupNamesStr := os.Args[1]
	userGroupNames := strings.Split(userGroupNamesStr, ",")
	result := map[string]*slack.UserGroup{}
	for _, userGroup := range userGroupNames {
		result[userGroup] = nil
	}
	apiToken := os.Getenv(ENV_SLACK_API_TOKEN)
	cli := slack.New(apiToken)
	userGroups, err := cli.GetUserGroups()
	if err != nil {
		panic(err)
	}
	for i, userGroup := range userGroups {
		_, ok := result[userGroup.Handle]
		if ok {
			result[userGroup.Handle] = &userGroups[i]
		}
	}
	for key, r := range result {
		id := ""
		if r != nil {
			id = r.ID
		}
		fmt.Println(key, ":", id)
	}
}
