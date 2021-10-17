package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/DisgoOrg/disgo/discord"
	"github.com/DisgoOrg/disgo/webhook"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/lambdacontext"
)

type DiscordWebHookConfig struct {
	webhookID    string
	webhookToken string
}

func (c DiscordWebHookConfig) Validate() (err error) {
	if len(c.webhookID) == 0 || len(c.webhookToken) == 0 {
		return fmt.Errorf("environment variables must be set")
	}

	return
}

func CreateDiscordWebHookConfig() (config DiscordWebHookConfig, err error) {
	config = DiscordWebHookConfig{
		webhookID:    os.Getenv("WEBHOOK_ID"),
		webhookToken: os.Getenv("WEBHOOK_TOKEN"),
	}

	err = config.Validate()

	return config, err
}

func CreateDiscordWebHookClient() (client *webhook.Client, err error) {
	config, err := CreateDiscordWebHookConfig()

	if err != nil {
		return nil, err
	}

	client = webhook.NewClient(discord.Snowflake(config.webhookID), config.webhookToken)

	return client, nil
}

// Discord Webhookで一度に送ることができるサイズである2000バイト以内にスライスの各要素を再構築する
func ReconstructSlicesforDiscordLimit(msgSlice []string) (newMsgSlice []string) {
	var currentSlice []string

	for i, msg := range msgSlice {
		if i == 0 {
			currentSlice = append(currentSlice, msg)
		} else {
			if len(strings.Join(currentSlice, "\r")+msg) >= 2000 {
				str := strings.Join(currentSlice, "\r")
				newMsgSlice = append(newMsgSlice, str)
				currentSlice = []string{msg}
			} else {
				currentSlice = append(currentSlice, msg)
			}
		}
	}

	if len(currentSlice) > 0 {
		newMsgSlice = append(newMsgSlice, strings.Join(currentSlice, "\r"))
	}

	return newMsgSlice
}

func SendToDiscord(msgSlice []string) (err error) {
	msgUnits := ReconstructSlicesforDiscordLimit(msgSlice)

	client, err := CreateDiscordWebHookClient()

	if err != nil {
		return err
	}

	var errs []error

	for _, msg := range msgUnits {
		if _, err := client.CreateContent(msg); err != nil {
			errs = append(errs, err)
		}
	}

	if len(errs) > 0 {
		return fmt.Errorf("errors: %v", errs)
	}

	return
}

func ParseAWSLogsToStringSlice(awsLogs events.CloudwatchLogsRawData) (msgSlice []string, err error) {
	data, err := awsLogs.Parse()
	if err != nil {
		return nil, err
	}

	for _, logEvent := range data.LogEvents {
		msgSlice = append(msgSlice, logEvent.Message)
	}

	return msgSlice, nil
}

func HandleRequest(ctx context.Context, event events.CloudwatchLogsEvent) (string, error) {
	lc, _ := lambdacontext.FromContext(ctx)

	msgSlice, err := ParseAWSLogsToStringSlice(event.AWSLogs)

	if err != nil {
		return fmt.Sprintf("Failed to parse log. request id: %s.", lc.AwsRequestID), err
	}

	if err := SendToDiscord(msgSlice); err != nil {
		return fmt.Sprintf("Failed to send log. request id: %s.", lc.AwsRequestID), err
	}

	return fmt.Sprintf("Success to send %d logs.", len(msgSlice)), nil
}

func main() {
	lambda.Start(HandleRequest)
}
