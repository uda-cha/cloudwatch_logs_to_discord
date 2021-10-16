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

func SendToDiscord(msgSlice []string) (err error) {
	config := DiscordWebHookConfig{
		webhookID:    os.Getenv("WEBHOOK_ID"),
		webhookToken: os.Getenv("WEBHOOK_TOKEN"),
	}

	if err := config.Validate(); err != nil {
		return err
	}

	msg := strings.Join(msgSlice, "\r")

	client := webhook.NewClient(discord.Snowflake(config.webhookID), config.webhookToken)
	_, err = client.CreateContent(msg)

	return err
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
