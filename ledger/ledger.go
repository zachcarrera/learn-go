package ledger

import (
	"errors"
	"fmt"
	"sort"
	"strings"
)

type Entry struct {
	Date        string // "Y-m-d"
	Description string
	Change      int // in cents
}

type message struct {
	index          int
	formattedEntry string
	err            error
}

const (
	nlLocaleString   = "nl-NL"
	usLocaleString   = "en-US"
	euCurrencyString = "EUR"
	usCurrencyString = "USD"
	headerFormatter  = "%-10s | %-25s | %s\n"
)

var (
	errInvalidLocale   = errors.New("Invalid locale specified")
	errInvalidCurrency = errors.New("Invalid currency specified")
	errInvalidDate     = errors.New("Invalide entry date provided")

	currencySymbols = map[string]rune{
		"EUR": '€',
		"USD": '$',
	}
)

func FormatLedger(currency string, locale string, entries []Entry) (string, error) {
	currencySymbol, validCurrency := currencySymbols[currency]
	if !validCurrency {
		return "", errInvalidCurrency
	}

	entriesCopy := make([]Entry, len(entries))
	copy(entriesCopy, entries)

	sort.Slice(entriesCopy, func(i, j int) bool {
		// sort by Date, then Description, then Change
		return entriesCopy[i].Date < entriesCopy[j].Date || (entriesCopy[i].Date == entriesCopy[j].Date && entriesCopy[i].Description < entriesCopy[j].Description) ||
			(entriesCopy[i].Description == entriesCopy[j].Description && entriesCopy[i].Change < entriesCopy[j].Change)
	})

	var ledger string
	switch locale {
	case nlLocaleString:
		ledger = fmt.Sprintf(headerFormatter, "Datum", "Omschrijving", "Verandering")
	case usLocaleString:
		ledger = fmt.Sprintf(headerFormatter, "Date", "Description", "Change")
	default:
		return "", errInvalidLocale
	}

	// Parallelism, always a great idea
	messages := make(chan message)
	for i, entry := range entriesCopy {
		go parseEntry(i, entry, messages, locale, currencySymbol)
	}
	entryStrings := make([]string, len(entriesCopy))
	for range entriesCopy {
		message := <-messages
		if message.err != nil {
			return "", message.err
		}
		entryStrings[message.index] = message.formattedEntry
	}
	ledger += strings.Join(entryStrings, "")
	return ledger, nil
}

func parseEntry(i int, entry Entry, messages chan<- message, locale string, currencySymbol rune) {
	year, month, day, err := parseDate(entry.Date)
	if err != nil {
		messages <- message{err: err}
		return
	}
	desc := entry.Description
	if len(desc) > 25 {
		desc = desc[:22] + "..."
	}
	date, err := formatDate(locale, year, month, day)
	if err != nil {
		messages <- message{err: err}
	}

	change, err := formatChange(locale, currencySymbol, entry.Change)
	if err != nil {
		messages <- message{err: err}
	}
	messages <- message{index: i, formattedEntry: fmt.Sprintf("%-10s | %-25s | %13s\n", date, desc, change)}
}

func parseDate(date string) (year, month, day string, err error) {
	if len(date) != 10 {
		err = errInvalidDate
		return
	}
	if date[4] != '-' || date[7] != '-' {
		err = errInvalidDate
		return
	}
	year, month, day = date[0:4], date[5:7], date[8:10]
	return
}

func formatDate(locale, year, month, day string) (string, error) {
	var date string
	switch locale {
	case nlLocaleString:
		date = fmt.Sprintf("%s-%s-%s", day, month, year)
	case usLocaleString:
		date = fmt.Sprintf("%s/%s/%s", month, day, year)
	default:
		return date, errInvalidLocale
	}
	return date, nil
}

func formatChange(locale string, currencySymbol rune, cents int) (string, error) {
	negative := false
	if cents < 0 {
		cents = cents * -1
		negative = true
	}
	var change string
	switch locale {
	case nlLocaleString:
		change = fmt.Sprintf("%c %s", currencySymbol, formatNumber(cents, ',', '.'))
		if negative {
			change = fmt.Sprintf("%s-", change)
		} else {
			change = fmt.Sprintf("%s ", change)
		}
	case usLocaleString:
		change = fmt.Sprintf("%c%s", currencySymbol, formatNumber(cents, '.', ','))
		if negative {
			change = fmt.Sprintf("(%s)", change)
		} else {
			change = fmt.Sprintf("%s ", change)
		}
	default:
		return "", errInvalidLocale
	}
	return change, nil
}

func formatNumber(cents int, decimalSeperator, thousandsSeperator rune) string {
	number := fmt.Sprintf("%c%02d", decimalSeperator, cents%100)
	cents /= 100
	number = fmt.Sprintf("%d%s", cents%1000, number)
	cents /= 1000

	for cents > 0 {
		number = fmt.Sprintf("%d%c%s", cents%1000, thousandsSeperator, number)
		cents /= 1000
	}
	return number
}
