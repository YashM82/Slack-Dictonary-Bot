**Slack Dictionary Bot**

A Slack bot that provides dictionary definitions for words using the Merriam-Webster API.

**## **Features****
- Fetches and displays definitions for given words.
- Handles errors gracefully if a word has no definitions.
- Includes a command to stop the bot.

## **Prerequisites**
- A [Slack App](https://api.slack.com/apps) with a bot token and app-level token.
- Merriam-Webster API key for accessing dictionary data.

## **Setup**

1. Clone the repository:
   ```bash
   git clone <repository-url>
   cd <project-directory>
   ```

2. Create a `.env` file in the root directory:
   ```plaintext
   SLACK_BOT_TOKEN=<your-slack-bot-token>
   SLACK_APP_TOKEN=<your-slack-app-token>
   API_KEY=<your-merriam-webster-api-key>
   ```

3. Install dependencies:
   ```bash
   go mod tidy
   ```

4. Run the bot:
   ```bash
   go run main.go
   ```

## **Commands**
- **`define <word>`**: Fetches the definition of the given word.
- **`end`**: Stops the bot.

## **Environment Variables**
- `SLACK_BOT_TOKEN`: Bot token from Slack.
- `SLACK_APP_TOKEN`: App-level token from Slack.
- `API_KEY`: API key for the Merriam-Webster Dictionary API.

## **Project Structure**
```
├── main.go       # Main application logic
├── .env          # Environment variables (excluded from Git)
├── .gitignore    # Ensures .env is not tracked
├── go.mod        # Module definition
├── go.sum        # Dependency checksum
```

## **License**
This project is open-source. Feel free to contribute!

---

Replace `<repository-url>` with your actual GitHub repository URL.
