# twitchkit-irc
"twitchkit-irc" is a lightweight Go Library that establishes a connection with Twitch IRC over TCP and sends structured data to a Go channel. It simplifies reading real-time messages from Twitch Chat.

For detailed field info, see the [pkg.go.dev documentation](https://pkg.go.dev/github.com/brandonwhited-dev/twitchkit-irc).
# Features
- Connects to Twitch IRC over TCP
- Parses chat messages into structured Go types
- Supports all metadata including emotes, badges, etc.
# Installation 
```bash
go get github.com/brandonwhited-dev/twitchkit-irc
```
# Example
See [`twitchkit-examples`](https://github.com/brandonwhited-dev/twitchkit-examples) for more examples of twitckhit.
```go
package main

import (
    "fmt"
    "log"

    "github.com/brandonwhited-dev/twitchkit-irc"
)

func main() {
    oauth := "YOUR_TWITCH_OAUTH"       // must have chat:read perm (do not include OAuth2:)
    username := "YOUR_TWITCH_USERNAME" // name of the account linked to your OAuth
    channel := "CHANNEL_NAME"           // name of the channel you want to read messages from

	// ===Connect to twitch IRC===
	// Create a message channel
	chatMessages := make(chan twitchkitirc.ChatMessage)

	// Start recieving messages
	err := twitchkitirc.Start(chatMessages, oauth, username, channel)
	if err != nil {
		log.Fatal("IRC Error: ", err)
	}

	// Handle messages
	go func() {
		for message := range chatMessages {
			fmt.Printf("%s: %s \n", message.Username, message.Message)
		}
	}()
}
```
