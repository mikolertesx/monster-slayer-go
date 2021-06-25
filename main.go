package main

import (
	"github.com/mikolertesx/monster-slayer/actions"
	"github.com/mikolertesx/monster-slayer/interaction"
)

var currentRound = 0

func main() {
	startGame()
	winner := ""

	for winner == "" {
		winner = executeRound()
	}

	endGame(winner)
}

func startGame() {
	interaction.PrintGreeting()
}

func executeRound() string {
	currentRound++
	isSpecialRound := currentRound%3 == 0

	interaction.ShowAvailableActions(isSpecialRound)
	userChoice := interaction.GetPlayerChoice(isSpecialRound)
	playerHealth, monsterHealth := actions.GetHealthAmounts()
	interaction.PrintHealth(playerHealth, monsterHealth)

	if userChoice == "ATTACK" {
		actions.AttackMonster(false)
	} else if userChoice == "HEAL" {
		actions.HealPlayer()
	} else {
		actions.AttackMonster(true)
	}

	_, monsterHealth = actions.GetHealthAmounts()

	if monsterHealth <= 0 {
		return "Player"
	}

	actions.AttackPlayer()

	playerHealth, _ = actions.GetHealthAmounts()

	if playerHealth <= 0 {
		return "MONSTER"
	}

	return ""
}

func endGame(winner string) {
	interaction.DeclareWinner(winner)
}
