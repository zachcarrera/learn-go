package poker

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

func BestHand(hands []string) ([]string, error) {
	panic("Please implement the BestHand function")
}
