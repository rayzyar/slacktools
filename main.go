package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/nlopes/slack"
)

const (
	ENV_SLACK_API_TOKEN = "ENV_SLACK_API_TOKEN"
)

type flagVar struct {
	prefix *bool
}

func main() {
	fVar := &flagVar{}
	fVar.prefix = flag.Bool("prefix", false, "for prefix match")
	flag.Parse()

	cmd := flag.Arg(0)
	switch cmd {
	case "subteam":
		subteamCMD(fVar)
	default:
		panic("command not supported")
	}

}

func subteamCMD(fVar *flagVar) {
	userGroupNamesStr := flag.Arg(1)
	userGroupNames := strings.Split(userGroupNamesStr, ",")
	result := map[string][]slack.UserGroup{}
	for _, userGroup := range userGroupNames {
		result[userGroup] = nil
	}
	apiToken := os.Getenv(ENV_SLACK_API_TOKEN)
	cli := slack.New(apiToken)
	userGroups, err := cli.GetUserGroups()
	if err != nil {
		panic(err)
	}
	for _, userGroup := range userGroups {
		for key := range result {
			if *fVar.prefix {
				if strings.HasPrefix(userGroup.Handle, key) {
					result[key] = append(result[key], userGroup)
				}
				continue
			}
			if userGroup.Handle == key {
				result[key] = append(result[key], userGroup)
			}
		}
	}
	for _, rSlice := range result {
		for _, r := range rSlice {
			id := ""
			id = r.ID
			fmt.Println(r.Handle, ":", id)
		}
	}
}
