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
)

func FormatLedger(currency string, locale string, entries []Entry) (string, error) {
	entriesCopy := make([]Entry, len(entries))
	copy(entriesCopy, entries)

	if len(entries) == 0 {
		if _, err := FormatLedger(currency, "en-US", []Entry{{Date: "2014-01-01", Description: "", Change: 0}}); err != nil {
			return "", err
		}
	}

	sort.Slice(entriesCopy, func(i, j int) bool {
		// sort by Date, then Description, then Change
		return entriesCopy[i].Date < entriesCopy[j].Date || (entriesCopy[i].Date == entriesCopy[j].Date && entriesCopy[i].Description < entriesCopy[j].Description) ||
			(entriesCopy[i].Description == entriesCopy[j].Description && entriesCopy[i].Change < entriesCopy[j].Change)
	})

	var s string
	switch locale {
	case nlLocaleString:
		s = fmt.Sprintf(headerFormatter, "Datum", "Omschrijving", "Verandering")
	case usLocaleString:
		s = fmt.Sprintf(headerFormatter, "Date", "Description", "Change")
	default:
		return "", errInvalidLocale
	}

	// Parallelism, always a great idea
	co := make(chan message)
	for i, et := range entriesCopy {
		go func(i int, entry Entry) {
			if len(entry.Date) != 10 {
				co <- message{err: errInvalidDate}
				return
			}
			d1, d2, d3, d4, d5 := entry.Date[0:4], entry.Date[4], entry.Date[5:7], entry.Date[7], entry.Date[8:10]
			if d2 != '-' {
				co <- message{err: errInvalidDate}
				return
			}
			if d4 != '-' {
				co <- message{err: errInvalidDate}
				return
			}
			de := entry.Description
			if len(de) > 25 {
				de = de[:22] + "..."
			} else {
				de = de + strings.Repeat(" ", 25-len(de))
			}
			var d string
			if locale == nlLocaleString {
				d = d5 + "-" + d3 + "-" + d1
			} else if locale == usLocaleString {
				d = d3 + "/" + d5 + "/" + d1
			}
			negative := false
			cents := entry.Change
			if cents < 0 {
				cents = cents * -1
				negative = true
			}
			var a string
			if locale == nlLocaleString {
				if currency == "EUR" {
					a += "€"
				} else if currency == "USD" {
					a += "$"
				} else {
					co <- message{err: errInvalidCurrency}
					return
				}
				a += " "
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
					a += parts[i] + "."
				}
				a = a[:len(a)-1]
				a += ","
				a += centsStr[len(centsStr)-2:]
				if negative {
					a += "-"
				} else {
					a += " "
				}
			} else if locale == usLocaleString {
				if negative {
					a += "("
				}
				if currency == "EUR" {
					a += "€"
				} else if currency == "USD" {
					a += "$"
				} else {
					co <- message{err: errInvalidCurrency}
					return
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
					a += parts[i] + ","
				}
				a = a[:len(a)-1]
				a += "."
				a += centsStr[len(centsStr)-2:]
				if negative {
					a += ")"
				} else {
					a += " "
				}
			} else {
				co <- message{err: errInvalidLocale}
				return
			}
			var al int
			for range a {
				al++
			}
			co <- message{index: i, formattedEntry: d + strings.Repeat(" ", 10-len(d)) + " | " + de + " | " +
				strings.Repeat(" ", 13-al) + a + "\n"}
		}(i, et)
	}
	ss := make([]string, len(entriesCopy))
	for range entriesCopy {
		v := <-co
		if v.err != nil {
			return "", v.err
		}
		ss[v.index] = v.formattedEntry
	}
	for i := 0; i < len(entriesCopy); i++ {
		s += ss[i]
	}
	return s, nil
}
