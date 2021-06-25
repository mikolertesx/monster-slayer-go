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

	var playerAttackDmg int
	var playerHealValue int
	var monsterAttackDmg int

	if userChoice == "ATTACK" {
		playerAttackDmg = actions.AttackMonster(false)
	} else if userChoice == "HEAL" {
		playerHealValue = actions.HealPlayer()
	} else {
		playerAttackDmg = actions.AttackMonster(true)
	}

	monsterAttackDmg = actions.AttackPlayer()

	playerHealth, monsterHealth := actions.GetHealthAmounts()

	roundData := interaction.RoundData{
		Action:             userChoice,
		PlayerHealth:       playerHealth,
		MonsterHealth:      monsterHealth,
		PlayerAttackDamage: playerAttackDmg,
		PlayerHealValue:    playerHealValue,
		MonsterAttackDmg:   monsterAttackDmg,
	}

	interaction.PrintRoundStatistics(&roundData)

	return ""
}

func endGame(winner string) {
	interaction.DeclareWinner(winner)
}
