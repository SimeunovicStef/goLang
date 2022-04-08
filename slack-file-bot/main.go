package main

import (
	"fmt"
	"os"

	"github.com/slack-go/slack"
)

func main() {

	os.Setenv("SLACK_BOT_TOKEN", "xoxb-3361259751493-3362808057205-u6xLBycgSZGgeadjibVF8C9E")
	os.Setenv("CHANNEL_ID", "C03B2N4DPS5")
	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	channelArr := []string{os.Getenv("CHANNEL_ID")}
	fileArr := []string{"/home/pr0xi/Desktop/go/src/go-server/slack-file-bot/cenovnik-fiat-500.pdf"}

	for i := 0; i < len(fileArr); i++ {
		params := slack.FileUploadParameters{
			Channels: channelArr,
			file:     fileArr[i],
		}
		file, err := api.UploadFile(params)
		if err != nil {
			fmt.Printf("%s\n", err)
			return
		}
		fmt.Printf("Name: %s, URL: %s", file.Name, file.URL)
	}
}
