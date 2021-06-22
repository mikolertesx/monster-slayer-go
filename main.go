package main

import (
	"github.com/mikolertesx/monster-slayer/interaction"
)

func main() {
	startGame()
	winner := ""

	for winner == "" {
		winner = executeRound()
	}

	endGame()
}

func startGame() {
	interaction.PrintGreeting()
}

func executeRound() string {
	return ""
}

func endGame() {}
