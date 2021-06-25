package actions

import (
	"math/rand"
	"time"
)

var randSource = rand.NewSource(time.Now().UnixNano())
var randGenerator = rand.New(randSource)
var currentMonsterHealth = MONSTER_HEALTH
var currentPlayerHealth = PLAYER_HEALTH

func AttackMonster(isSpecialAttack bool) {
	minAttackValue := PLAYER_ATTACK_MIN_DMG
	maxAttackValue := PLAYER_ATTACK_MAX_DMG

	if isSpecialAttack {
		minAttackValue = PLAYER_SPECIAL_ATTACK_MIN_DMG
		maxAttackValue = PLAYER_SPECIAL_ATTACK_MAX_DMG
	}

	dmgValue := generateRandBetween(minAttackValue, maxAttackValue)
	currentMonsterHealth -= dmgValue
}

func HealPlayer() {
	minHealValue := PLAYER_HEAL_MIN_VALUE
	maxHealValue := PLAYER_HEAL_MAX_VALUE

	healValue := generateRandBetween(minHealValue, maxHealValue)
	healthDiff := 100 - currentPlayerHealth

	if healthDiff >= healValue {
		currentPlayerHealth += healValue
	} else {
		currentPlayerHealth = PLAYER_HEALTH
	}
}

func AttackPlayer() {
	minAttackValue := MONSTER_ATTACK_MIN_DAMAGE
	maxAttackValue := MONSTER_ATTACK_MAX_DAMAGE

	dmgValue := generateRandBetween(minAttackValue, maxAttackValue)
	currentPlayerHealth -= dmgValue
}

func GetHealthAmounts() (int, int) {
	return currentPlayerHealth, currentMonsterHealth
}

func generateRandBetween(min int, max int) int {
	return randGenerator.Intn(max-min) + min
}
