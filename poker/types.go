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
