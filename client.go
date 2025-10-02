package twitchkitirc

// twitchIRC is meant to make it simple to conntect to twitch's IRC and read chat messages

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"net"
	"strings"
)

// ChatMessage is a struct of useful message information
//
//	Username: string - The chatters username
//	Message: string - The message sent
type ChatMessage struct {
	Username string
	Message  string
}
