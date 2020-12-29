package day22

import (
	"fmt"
	"strconv"
	"strings"
)

// Cards represents the cards of each player
type Cards []int

func FirstPart(lines []string) (int, error) {
	players, err := readLines(lines)
	if err != nil {
		return 0, err
	}
	result := playFirstPart(players[0], players[1])
	score := getScore(result)
	return score, nil
}

func SecondPart(lines []string) (int, error) {
	players, err := readLines(lines)
	if err != nil {
		return 0, err
	}
	_, result := playRecursiveCombat(players[0], players[1])
	score := getScore(result)
	return score, nil
}

func readLines(lines []string) ([]Cards, error) {
	player := 0
	cards1 := Cards{}
	cards2 := Cards{}
	for _, l := range lines[1:] {
		if strings.Contains(l, "Player") {
			player++
			continue
		} else if l == "" {
			continue
		}
		v, err := strconv.Atoi(l)
		if err != nil {
			return nil, err
		}
		if player == 0 {
			cards1 = append(cards1, v)
		} else {
			cards2 = append(cards2, v)
		}
	}
	return []Cards{cards1, cards2}, nil
}

func getScore(cards Cards) int {
	var score, carLength = 0, len(cards)
	for i, v := range cards {
		score += v * (carLength - i)
	}
	return score
}

func playFirstPart(player1, player2 Cards) Cards {
	for len(player1) > 0 && len(player2) > 0 {
		card1 := player1[0]
		card2 := player2[0]
		player1 = player1[1:]
		player2 = player2[1:]
		if card1 > card2 {
			player1 = append(player1, card1, card2)
		} else {
			player2 = append(player2, card2, card1)
		}
	}
	if len(player1) > 0 {
		return player1
	}
	return player2
}

func playRecursiveCombat(player1, player2 Cards) (int, Cards) {
	prevRounds := map[string]bool{}
	for len(player1) > 0 && len(player2) > 0 {
		// if the game is equal to some previous game, with same cards, etc.
		// player1 automatically wins
		strRepr := fmt.Sprintf("%#v|%#v", player1, player2)
		if _, ok := prevRounds[strRepr]; ok {
			return 1, player1
		}
		prevRounds[strRepr] = true
		card1 := player1[0]
		card2 := player2[0]
		player1 = player1[1:]
		player2 = player2[1:]
		if card1 <= len(player1) && card2 <= len(player2) {
			player1CopiedCards := copyCards(player1[:card1])
			player2CopiedCards := copyCards(player2[:card2])
			winner, _ := playRecursiveCombat(player1CopiedCards, player2CopiedCards)
			if winner == 1 {
				player1 = append(player1, card1, card2)
			} else {
				player2 = append(player2, card2, card1)
			}
		} else {
			if card1 > card2 {
				player1 = append(player1, card1, card2)
			} else {
				player2 = append(player2, card2, card1)
			}
		}
	}
	if len(player1) > 0 {
		return 1, player1
	}
	return 2, player2
}

func copyCards(cards Cards) Cards {
	return append([]int{}, cards...)
}
