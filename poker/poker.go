package poker

import (
	"errors"
)

type Suit rune

const (
	Heart   Suit = '♡'
	Spade   Suit = '♤'
	Diamond Suit = '♢'
	Club    Suit = '♧'
)

type CardRank int

const (
	Two CardRank = iota
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
	Ace
)

var cardRanks = map[string]CardRank{
	"2":  Two,
	"3":  Three,
	"4":  Four,
	"5":  Five,
	"6":  Six,
	"7":  Seven,
	"8":  Eight,
	"9":  Nine,
	"10": Ten,
	"J":  Jack,
	"Q":  Queen,
	"K":  King,
	"A":  Ace,
}

type Card struct {
	suit Suit
	rank CardRank
}

type HandRank int

const (
	HighCard HandRank = iota
	OnePair
	TwoPair
	ThreeOfAKind
	Straight
	Flush
	FullHouse
	FourOfAKind
	StraightFlush
)

type Hand struct {
	raw   string
	cards []Card
	rank  HandRank
}

func BestHand(hands []string) ([]string, error) {
	panic("Please implement the BestHand function")
}

func parseCardRank(card string) (CardRank, error) {
	rank := card[0:1]
	if len(card) == 5 {
		rank = card[0:2]
	}
	validRank, ok := cardRanks[rank]
	if !ok {
		return -1, errors.New("invalid card rank")
	}
	return validRank, nil
}

func parseSuit(card string) (Suit, error) {
	suit := []rune(card[1:])
	if len(card) == 5 {
		suit = []rune(card[2:])
	}
	switch Suit(suit[0]) {
	case Heart, Club, Diamond, Spade:
		return Suit(suit[0]), nil
	default:
		return -1, errors.New("invalid suit")
	}
}
