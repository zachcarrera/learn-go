package poker

import (
	"errors"
	"sort"
	"strings"
)

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
		} else if hand.rank == bestHand.rank {
			switch evaluateTie(hand, bestHand) {
			case 1:
				// hand beats best
				bestHands = []string{rawHand}
				bestHand = hand
			case 0:
				// still tied so append to bestHands
				bestHands = append(bestHands, rawHand)
			}
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

	if cards[len(cards)-1].rank == Ace {
		cards = append([]Card{{cards[len(cards)-1].suit, -1}}, cards...)
	}

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

	if len(cards) == 6 {
		cards = cards[1:]
		delete(cardRankCount, -1)
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

func evaluateTie(h1, h2 Hand) int {
	switch h1.rank {
	case HighCard, Flush:
		return compareHighCard(h1, h2)
	case OnePair, TwoPair, ThreeOfAKind, FullHouse, FourOfAKind:
		return compareCardRankCount(h1, h2)
	case Straight, StraightFlush:
		return compareStraight(h1, h2)
	}
	return -1
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

func compareCardRankCount(h1, h2 Hand) int {
	h1CardRankCount := make(map[CardRank]int)
	for _, c := range h1.cards {
		h1CardRankCount[c.rank]++
	}
	var rc1 [][2]int
	for rank, count := range h1CardRankCount {
		rc1 = append(rc1, [2]int{int(rank), count})
	}
	sort.SliceStable(rc1, func(i, j int) bool {
		if rc1[i][1] == rc1[j][1] {
			return rc1[i][0] < rc1[j][0]
		}
		return rc1[i][1] < rc1[j][1]
	})
	h2CardRankCount := make(map[CardRank]int)
	for _, c := range h2.cards {
		h2CardRankCount[c.rank]++
	}
	var rc2 [][2]int
	for rank, count := range h2CardRankCount {
		rc2 = append(rc2, [2]int{int(rank), count})
	}
	sort.SliceStable(rc2, func(i, j int) bool {
		if rc2[i][1] == rc2[j][1] {
			return rc2[i][0] < rc2[j][0]
		}
		return rc2[i][1] < rc2[j][1]
	})
	for i := len(rc1) - 1; i >= 0; i-- {
		rank1, rank2 := rc1[i][0], rc2[i][0]
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
