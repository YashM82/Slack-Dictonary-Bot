package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/shomali11/slacker"
)

type Response struct {
	Shortdef []string `json:"shortdef"`
}

func FetchDefination(word string) ([]string, error) {
	apiKey := os.Getenv("API_KEY")

	apiURL := fmt.Sprintf("https://www.dictionaryapi.com/api/v3/references/sd3/json/%s?key=%s", word, apiKey)

	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, fmt.Errorf("error making API call: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch data: status code %d", resp.StatusCode)
	}
	// Decode JSON response
	var response []Response
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, fmt.Errorf("error decoding API response: %v", err)
	}
	// Handle empty or invalid responses
	if len(response) == 0 || len(response[0].Shortdef) == 0 {
		return nil, fmt.Errorf("no definitions found for the word '%s'", word)
	}
	// Return only the `shortdef` field
	return response[0].Shortdef, nil
}

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		fmt.Println("Command Evenets")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)

		fmt.Println()
	}
}

func main() {

	// Load environment variables from the `.env` file if it exists
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, relying on system environment variables")
	}

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	ctx, cancel := context.WithCancel(context.Background())

	go printCommandEvents(bot.CommandEvents())

	bot.Command("define <word>", &slacker.CommandDefinition{
		Description: "Dictonary Bot",
		Examples:    []string{"g"},
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			word := request.Param("word")
			if word == "" {
				response.Reply("Please provide a word to define.")
				return
			}
			definations, err := FetchDefination(word)
			if err != nil {
				response.Reply(fmt.Sprintf("Error: %s", err))
				return
			}
			formattedDefs := ""
			for i, def := range definations {
				formattedDefs += fmt.Sprintf("%d. %s\n", i+1, def)
			}

			response.Reply(fmt.Sprintf("Defination of word %s is \n %s", word, formattedDefs))
		},
	})

	bot.Command("end", &slacker.CommandDefinition{
		Description: "End the bot process",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			response.Reply("Recieved End")
			response.Reply("Bot closing!")
			cancel()

		},
	})

	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Bot closed")
}
