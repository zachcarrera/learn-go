package poker

import (
	"errors"
	"sort"
	"strings"
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
	bestHand := Hand{"", nil, -1}
	var bestHands []string
	for _, rawHand := range hands {
		hand, err := parseHand(rawHand)
		if err != nil {
			return nil, err
		}
		if hand.rank > bestHand.rank {
			bestHands = []string{rawHand}
			bestHand = hand
		}
	}
	return bestHands, nil
}

func parseHand(hand string) (Hand, error) {
	cards := strings.Split(hand, " ")
	if len(cards) != 5 {
		return Hand{}, errors.New("invalid hand")
	}
	var parsedCards []Card
	for _, c := range cards {
		if len(c) < 4 || len(c) > 5 {
			return Hand{}, errors.New("invalid rank-suit combination")
		}
		rank, err := parseCardRank(c)
		if err != nil {
			return Hand{}, err
		}
		suit, err := parseSuit(c)
		if err != nil {
			return Hand{}, err
		}
		parsedCards = append(parsedCards, Card{suit, rank})
	}
	sort.SliceStable(parsedCards, func(i, j int) bool { return parsedCards[i].rank < parsedCards[j].rank })
	return Hand{hand, parsedCards, computeHandRank(parsedCards)}, nil
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

func compareHighCard(h1, h2 Hand) int {
	for i := len(h1.cards) - 1; i >= 0; i-- {
		rank1, rank2 := h1.cards[i].rank, h2.cards[i].rank
		switch {
		case rank1 == rank2:
			continue
		case rank1 < rank2:
			return -1
		case rank1 > rank2:
			return 1
		}
	}
	return 0
}

func compareStraight(h1, h2 Hand) int {
	h1HighCard, h2HighCard := h1.cards[len(h1.cards)-1], h2.cards[len(h2.cards)-1]

	if h1HighCard.rank == Ace && h1.cards[0].rank == 0 {
		h1HighCard = h1.cards[len(h1.cards)-2]
	}
	if h2HighCard.rank == Ace && h2.cards[0].rank == 0 {
		h2HighCard = h2.cards[len(h2.cards)-2]
	}

	switch {
	case h1HighCard.rank < h2HighCard.rank:
		return -1
	case h1HighCard.rank > h2HighCard.rank:
		return 1
	default:
		return 0
	}
}
