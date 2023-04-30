package main

import (
	"fmt"
	"net/smtp"
	"os"
	"os/exec"
	"time"

	"github.com/slack-go/slack"
)

func fileBot() {
	// Download image
	cmd := exec.Command("./script/conf.sh")
	err := cmd.Run()
	if err != nil {
		fmt.Println("could not run command: ", err)
	}
	// Configure Slack Bot
	fmt.Println("Slack File Bot")
	err = os.Setenv("SLACK_BOT_TOKEN", "tocken")
	if err != nil {
		return
	}
	err = os.Setenv("CHANEL_ID", "channel id")
	if err != nil {
		return
	}
	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	channelArr := []string{os.Getenv("CHANEL_ID")}
	fileArr := []string{"./cool-menu.png", "./rosto-menu.jpg", "pilsner-menu.png"}
	for i := 0; i < len(fileArr); i++ {
		params := slack.FileUploadParameters{
			Channels: channelArr,
			File:     fileArr[i],
		}
		file, err := api.UploadFile(params)
		if err != nil {
			fmt.Println("")
			fmt.Println("Error:")
			fmt.Println("")
			now := time.Now()
			fmt.Println(now.Format("2 Jan 06 03:04PM"))
			fmt.Printf("%s\n", err)
			return
		}
		var fileUrl = "/Users/maksymmochaliuk/work/home/projects/slack-file-upload/file-bot"
		now := time.Now()
		fmt.Println("")
		fmt.Println(now.Format("2 Jan 06 03:04PM"))
		fmt.Printf("Name: %s, URL:%s\n", file.Name, fileUrl)
		fmt.Println("")
		fmt.Println("Success!")
	}
}
func sendMail() {
	mail := "maks.mochaliuk@gmail.com"
	token := "cbxggbxdqesbgeqz"
	subject := "🍕 Menu For Today"
	intro := "Hello! 🖐️"
	introEnter := "This is your personal menu assistant 👩‍🔬"
	body := "List of menu you can find by click this link: https://app.slack.com/client/T052GJDTYS0/C0527G4VDD4"
	auth := smtp.PlainAuth(
		"",
		mail,
		token,
		"smtp.gmail.com",
	)
	fmt.Println("Sending email to this adress:", mail)
	msg := "Subject:" + subject + "\n" + intro + "\n" + introEnter + "\n" + body
	err := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		mail,
		[]string{mail},
		[]byte(msg),
	)
	// If error
	if err != nil {
		fmt.Println("err")
	}
}
func main() {
	fileBot()
	sendMail()
}
