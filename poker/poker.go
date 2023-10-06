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

func computeHandRank(cards []Card) HandRank {
	cardRankCount := make(map[CardRank]int)
	consecutiveRankCount := 1
	suit := cards[0].suit
	for i, c := range cards {
		cardRankCount[c.rank]++
		suit &= c.suit
		if i > 0 && c.rank-1 == cards[i-1].rank {
			consecutiveRankCount++
		}
	}
	var isFlush, isStraight bool
	if suit == cards[0].suit {
		isFlush = true
	}
	if consecutiveRankCount == 5 {
		isStraight = true
	}

	var pairs int
	var isThreeOfAKind, isFourOfAKind bool
	for _, count := range cardRankCount {
		switch count {
		case 2:
			pairs++
		case 3:
			isThreeOfAKind = true
		case 4:
			isFourOfAKind = true
		}
	}

	switch {
	case isStraight && isFlush:
		return StraightFlush
	case isFourOfAKind:
		return FourOfAKind
	case isThreeOfAKind && pairs == 1:
		return FullHouse
	case isFlush:
		return Flush
	case isStraight:
		return Straight
	case isThreeOfAKind:
		return ThreeOfAKind
	case pairs == 2:
		return TwoPair
	case pairs == 1:
		return OnePair
	default:
		return HighCard
	}
}
