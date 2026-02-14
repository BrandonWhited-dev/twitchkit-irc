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
See [`twitchkit-examples`](https://github.com/brandonwhited-dev/twitchkit-examples) for examples of twitckhit implementation.
