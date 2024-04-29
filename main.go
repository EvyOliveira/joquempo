package main

import (
	"fmt"
	"math/rand"
)

var (
	pedra              = "PEDRA"
	papel              = "PAPEL"
	tesoura            = "TESOURA"
	firstParticipant   string
	secoundParticipant string
)

type participant struct {
	ID     uint32
	Name   string
	Result string
}

func main() {

	fmt.Println("Jogador 1, escolha entre PEDRA, PAPEL ou TESOURA: ")
	fmt.Scanln(&firstParticipant)

	fmt.Println("Jogador 2, escolha entre PEDRA, PAPEL ou TESOURA: ")
	fmt.Scanln(&secoundParticipant)

	winner, err := gameDecision(firstParticipant, secoundParticipant)
	if err != nil {
		fmt.Println("Algo está errado:", err)
	} else {
		fmt.Printf("Você é o vencedor: %s\n", winner)
	}

}

func (p participant) decision() string {
	decision := []string{pedra, papel, tesoura}
	return decision[rand.Intn(len(decision))]
}

func gameDecision(firstPlayer, secoundPlayer string) (string, error) {
	switch {
	case firstPlayer == pedra && secoundPlayer == papel:
		return secoundPlayer, nil
	case firstPlayer == papel && secoundPlayer == tesoura:
		return secoundPlayer, nil
	case firstPlayer == tesoura && secoundPlayer == pedra:
		return secoundPlayer, nil
	case firstPlayer == papel && secoundPlayer == pedra:
		return firstPlayer, nil
	case firstPlayer == tesoura && secoundPlayer == papel:
		return firstPlayer, nil
	case firstPlayer == pedra && secoundPlayer == tesoura:
		return firstPlayer, nil
	case firstPlayer == secoundPlayer:
		return "Os jogadores empataram!", nil
	default:
		return "", fmt.Errorf("Essa opção não é válida!")
	}
}
