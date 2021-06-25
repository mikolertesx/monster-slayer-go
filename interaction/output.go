package interaction

import (
	"fmt"
)

type RoundData struct {
	Action             string
	PlayerAttackDamage int
	PlayerHealValue    int
	MonsterAttackDmg   int
	PlayerHealth       int
	MonsterHealth      int
}

func PrintGreeting() {
	fmt.Println("Monster Slayer")
	fmt.Println("Starting a new game...")
	fmt.Println("Good luck!")
}

func PrintHealth(player int, monster int) {
	fmt.Printf("Player: %d\n", player)
	fmt.Printf("Monster: %d\n", monster)
}

func ShowAvailableActions(specialAttackIsAvailable bool) {
	fmt.Println("Please choose your action")
	fmt.Println("-----------------------")
	fmt.Println("(1) Attack Monster")
	fmt.Println("(2) Heal")
	if specialAttackIsAvailable {
		fmt.Println("(3) Special Attack")
	}
}

func PrintRoundStatistics(roundData *RoundData) {
	if roundData.Action == "ATTACK" {
		fmt.Printf("Player attacked monster for %v damage.\n", roundData.PlayerAttackDamage)
	} else if roundData.Action == "SPECIAL_ATTACK" {
		fmt.Printf("Player performed a strong attack against monster for %v damage.\n", roundData.PlayerAttackDamage)
	} else {
		fmt.Printf("Player healed for %v.\n", roundData.PlayerHealValue)
	}

	fmt.Printf("Monster attacked player for %v damage.\n", roundData.MonsterAttackDmg)
	fmt.Printf("Player Health: %v.\n", roundData.PlayerHealth)
	fmt.Printf("Monster Health: %v.\n", roundData.MonsterHealth)
}

func DeclareWinner(winner string) {
	fmt.Println("----------------------------")
	fmt.Println("GAME OVER!")
	fmt.Println("----------------------------")
	fmt.Printf("%v won!\n", winner)
}
