package ledger

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
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
	_, validCurrency := currencySymbols[currency]
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
		go parseEntry(i, entry, messages, locale, currency)
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

func parseEntry(i int, entry Entry, messages chan<- message, locale string, currency string) {
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

	change, err := formatChange(locale, currency, entry.Change)
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

func formatChange(locale, currency string, cents int) (string, error) {
	negative := false
	if cents < 0 {
		cents = cents * -1
		negative = true
	}
	var change string
	switch locale {
	case nlLocaleString:
		if currency == "EUR" {
			change += "€"
		} else if currency == "USD" {
			change += "$"
		} else {
			return "", errInvalidCurrency
		}
		change += " "
		centsStr := strconv.Itoa(cents)
		switch len(centsStr) {
		case 1:
			centsStr = "00" + centsStr
		case 2:
			centsStr = "0" + centsStr
		}
		rest := centsStr[:len(centsStr)-2]
		var parts []string
		for len(rest) > 3 {
			parts = append(parts, rest[len(rest)-3:])
			rest = rest[:len(rest)-3]
		}
		if len(rest) > 0 {
			parts = append(parts, rest)
		}
		for i := len(parts) - 1; i >= 0; i-- {
			change += parts[i] + "."
		}
		change = change[:len(change)-1]
		change += ","
		change += centsStr[len(centsStr)-2:]
		if negative {
			change += "-"
		} else {
			change += " "
		}
	case usLocaleString:
		if negative {
			change += "("
		}
		if currency == "EUR" {
			change += "€"
		} else if currency == "USD" {
			change += "$"
		} else {
			return "", errInvalidCurrency
		}
		centsStr := strconv.Itoa(cents)
		switch len(centsStr) {
		case 1:
			centsStr = "00" + centsStr
		case 2:
			centsStr = "0" + centsStr
		}
		rest := centsStr[:len(centsStr)-2]
		var parts []string
		for len(rest) > 3 {
			parts = append(parts, rest[len(rest)-3:])
			rest = rest[:len(rest)-3]
		}
		if len(rest) > 0 {
			parts = append(parts, rest)
		}
		for i := len(parts) - 1; i >= 0; i-- {
			change += parts[i] + ","
		}
		change = change[:len(change)-1]
		change += "."
		change += centsStr[len(centsStr)-2:]
		if negative {
			change += ")"
		} else {
			change += " "
		}
	default:
		return "", errInvalidLocale
	}
	return change, nil
}
