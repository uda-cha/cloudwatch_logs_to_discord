package main

import (
	"context"
	"fmt"
	"os"

	"github.com/DisgoOrg/disgo/discord"
	"github.com/DisgoOrg/disgo/webhook"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/lambdacontext"
)

func SendToDiscord(msg string) (err error) {
	webhookID := os.Getenv("WEBHOOK_ID")
	webhookToken := os.Getenv("WEBHOOK_TOKEN")

	if len(webhookID) == 0 || len(webhookToken) == 0 {
		return fmt.Errorf("environment variables must be set")
	}

	client := webhook.NewClient(discord.Snowflake(webhookID), webhookToken)
	_, err = client.CreateContent(msg)

	return err
}

func HandleRequest(ctx context.Context, event events.CloudwatchLogsEvent) (string, error) {
	lc, _ := lambdacontext.FromContext(ctx)

	data, err := event.AWSLogs.Parse()
	if err != nil {
		return fmt.Sprintf("Failed to parse AWSLogs. request id: %s", lc.AwsRequestID), err
	}

	for _, logEvent := range data.LogEvents {
		if err := SendToDiscord(logEvent.Message); err != nil {
			return fmt.Sprintf("Failed to send log event id: %s.", logEvent.ID), err
		}
	}

	return fmt.Sprintf("Success to send %d logs.", len(data.LogEvents)), nil
}

func main() {
	lambda.Start(HandleRequest)
}
