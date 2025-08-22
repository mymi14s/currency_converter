package currency

import (
	"fmt"
	"math"
	"strings"
)

func AmountToWords(amount float64, currency string) map[string]interface{} {
	intPart := int64(amount)                                         // convert to int64
	fracPart := int64(math.Round((amount - float64(intPart)) * 100)) // convert to int64

	mainUnit, subUnit := getCurrencyUnits(currency)
	amountInWords := fmt.Sprintf("%s %s", numberToWords(intPart), mainUnit)

	if fracPart > 0 {
		amountInWords += fmt.Sprintf(" and %s %s", numberToWords(fracPart), subUnit)
	}

	return map[string]interface{}{
		"amount_in_word": strings.Title(amountInWords),
		"amount":         amount,
		"currency":       currency,
	}
}

var smallNumbers = []string{
	"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
	"ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen",
	"seventeen", "eighteen", "nineteen",
}

var tensNumbers = []string{
	"", "", "twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety",
}

var scaleNumbers = []string{
	"", "thousand", "million", "billion", "trillion", "quadrillion", "quintillion",
}

func numberToWords(n int64) string {
	if n == 0 {
		return "zero"
	}

	words := []string{}
	group := 0

	for n > 0 {
		rem := n % 1000
		if rem != 0 {
			part := threeDigitToWords(int(rem))
			if scaleNumbers[group] != "" {
				part += " " + scaleNumbers[group]
			}
			words = append([]string{part}, words...)
		}
		n /= 1000
		group++
	}

	return strings.Join(words, " ")
}

func threeDigitToWords(n int) string {
	words := []string{}

	if n >= 100 {
		words = append(words, smallNumbers[n/100]+" hundred")
		n %= 100
		if n > 0 {
			words = append(words, "and")
		}
	}

	if n >= 20 {
		words = append(words, tensNumbers[n/10])
		n %= 10
		if n > 0 {
			words[len(words)-1] += " " + smallNumbers[n]
		}
	} else if n > 0 {
		words = append(words, smallNumbers[n])
	}

	return strings.Join(words, " ")
}

func getCurrencyUnits(currency string) (string, string) {
	switch currency {
	case "AED":
		return "Dirham", "Fils"
	case "AFN":
		return "Afghani", "Pul"
	case "ALL":
		return "Lek", "Qindarka"
	case "AMD":
		return "Dram", "Luma"
	case "AOA":
		return "Kwanza", "Cêntimo"
	case "ARS":
		return "Peso", "Centavo"
	case "AUD":
		return "Dollars", "Cents"
	case "AWG":
		return "Florin", "Cents"
	case "AZN":
		return "Manat", "Qəpik"
	case "BAM":
		return "Convertible Mark", "Fening"
	case "BBD":
		return "Dollars", "Cents"
	case "BDT":
		return "Taka", "Poisha"
	case "BGN":
		return "Lev", "Stotinki"
	case "BHD":
		return "Dinar", "Fils"
	case "BIF":
		return "Franc", "N/A"
	case "BMD":
		return "Dollars", "Cents"
	case "BND":
		return "Dollars", "Sen"
	case "BOB":
		return "Boliviano", "Centavo"
	case "BRL":
		return "Real", "Centavos"
	case "BSD":
		return "Dollars", "Cents"
	case "BTN":
		return "Ngultrum", "Chetrum"
	case "BWP":
		return "Pula", "Thebe"
	case "BYN":
		return "Ruble", "Kopek"
	case "BZD":
		return "Dollars", "Cents"
	case "CAD":
		return "Dollars", "Cents"
	case "CDF":
		return "Franc", "Centime"
	case "CHF":
		return "Franc", "Rappen"
	case "CLP":
		return "Peso", "N/A"
	case "CNY":
		return "Yuan", "Jiao"
	case "COP":
		return "Peso", "Centavo"
	case "CRC":
		return "Colón", "Céntimo"
	case "CUP":
		return "Peso", "Centavo"
	case "CVE":
		return "Escudo", "Centavo"
	case "CZK":
		return "Koruna", "Heller"
	case "DJF":
		return "Franc", "N/A"
	case "DKK":
		return "Krone", "Øre"
	case "DOP":
		return "Peso", "Centavo"
	case "DZD":
		return "Dinar", "Centime"
	case "EGP":
		return "Pound", "Piastre"
	case "ERN":
		return "Nakfa", "Cent"
	case "ETB":
		return "Birr", "Santim"
	case "EUR":
		return "Euro", "Cents"
	case "FJD":
		return "Dollar", "Cents"
	case "FKP":
		return "Pound", "Penny"
	case "GBP":
		return "Pound", "Pence"
	case "GEL":
		return "Lari", "Tetri"
	case "GHS":
		return "Cedi", "Pesewa"
	case "GIP":
		return "Pound", "Penny"
	case "GMD":
		return "Dalasi", "Butut"
	case "GNF":
		return "Franc", "N/A"
	case "GTQ":
		return "Quetzal", "Centavo"
	case "GYD":
		return "Dollar", "Cents"
	case "HKD":
		return "Dollars", "Cents"
	case "HNL":
		return "Lempira", "Centavo"
	case "HRK":
		return "Kuna", "Lipa"
	case "HTG":
		return "Gourde", "Centime"
	case "HUF":
		return "Forint", "N/A"
	case "IDR":
		return "Rupiah", "Sen"
	case "ILS":
		return "New Shekel", "Agorot"
	case "INR":
		return "Rupee", "Paisa"
	case "IQD":
		return "Dinar", "Fils"
	case "IRR":
		return "Rial", "N/A"
	case "ISK":
		return "Krona", "N/A"
	case "JMD":
		return "Dollar", "Cents"
	case "JOD":
		return "Dinar", "Fils"
	case "JPY":
		return "Yen", "N/A"
	case "KES":
		return "Shilling", "Cent"
	case "KGS":
		return "Som", "Tyiyn"
	case "KHR":
		return "Riel", "Sen"
	case "KMF":
		return "Franc", "N/A"
	case "KPW":
		return "Won", "Chon"
	case "KRW":
		return "Won", "N/A"
	case "KWD":
		return "Dinar", "Fils"
	case "KYD":
		return "Dollar", "Cents"
	case "KZT":
		return "Tenge", "Tïın"
	case "LAK":
		return "Kip", "Att"
	case "LBP":
		return "Pound", "Piastre"
	case "LKR":
		return "Rupee", "Cent"
	case "LRD":
		return "Dollar", "Cents"
	case "LSL":
		return "Loti", "Sente"
	case "LYD":
		return "Dinar", "Dirham"
	case "MAD":
		return "Dirham", "Centime"
	case "MDL":
		return "Leu", "Ban"
	case "MGA":
		return "Ariary", "Iraimbilanja"
	case "MKD":
		return "Denar", "Deni"
	case "MMK":
		return "Kyat", "Pya"
	case "MNT":
		return "Tugrik", "Möngö"
	case "MOP":
		return "Pataca", "Avo"
	case "MRU":
		return "Ouguiya", "Khoums"
	case "MUR":
		return "Rupee", "Cent"
	case "MVR":
		return "Rufiyaa", "Laari"
	case "MWK":
		return "Kwacha", "Tambala"
	case "MXN":
		return "Peso", "Centavo"
	case "MYR":
		return "Ringgit", "Sen"
	case "MZN":
		return "Metical", "Centavo"
	case "NAD":
		return "Dollar", "Cents"
	case "NGN":
		return "Naira", "Kobo"
	case "NIO":
		return "Córdoba", "Centavo"
	case "NOK":
		return "Krone", "Øre"
	case "NPR":
		return "Rupee", "Paisa"
	case "NZD":
		return "Dollars", "Cents"
	case "OMR":
		return "Rial", "Baisa"
	case "PAB":
		return "Balboa", "Centésimo"
	case "PEN":
		return "Sol", "Céntimo"
	case "PGK":
		return "Kina", "Toea"
	case "PHP":
		return "Peso", "Sentimo"
	case "PKR":
		return "Rupee", "Paisa"
	case "PLN":
		return "Złoty", "Grosz"
	case "PYG":
		return "Guaraní", "Céntimo"
	case "QAR":
		return "Riyal", "Dirham"
	case "RON":
		return "Leu", "Ban"
	case "RSD":
		return "Dinar", "Para"
	case "RUB":
		return "Ruble", "Kopek"
	case "RWF":
		return "Franc", "N/A"
	case "SAR":
		return "Riyal", "Halala"
	case "SBD":
		return "Dollar", "Cents"
	case "SCR":
		return "Rupee", "Cent"
	case "SDG":
		return "Pound", "Piastre"
	case "SEK":
		return "Krona", "Öre"
	case "SGD":
		return "Dollar", "Cents"
	case "SHP":
		return "Pound", "Penny"
	case "SLL":
		return "Leone", "Cent"
	case "SOS":
		return "Shilling", "Cent"
	case "SRD":
		return "Dollar", "Cent"
	case "SSP":
		return "Pound", "Piastre"
	case "STN":
		return "Dobra", "Cêntimo"
	case "SVC":
		return "Colón", "Centavo"
	case "SYP":
		return "Pound", "Piastre"
	case "SZL":
		return "Lilangeni", "Cent"
	case "THB":
		return "Baht", "Satang"
	case "TJS":
		return "Somoni", "Diram"
	case "TMT":
		return "Manat", "Tenge"
	case "TND":
		return "Dinar", "Millime"
	case "TOP":
		return "Pa'anga", "Seniti"
	case "TRY":
		return "Lira", "Kuruş"
	case "TTD":
		return "Dollar", "Cent"
	case "TWD":
		return "New Dollar", "Cent"
	case "TZS":
		return "Shilling", "Cent"
	case "UAH":
		return "Hryvnia", "Kopiyka"
	case "UGX":
		return "Shilling", "N/A"
	case "USD":
		return "Dollars", "Cents"
	case "UYU":
		return "Peso", "Centésimo"
	case "UZS":
		return "Som", "Tiyin"
	case "VND":
		return "Đồng", "Hào"
	case "VUV":
		return "Vatu", "N/A"
	case "WST":
		return "Tala", "Sene"
	case "XAF":
		return "Franc", "Centime"
	case "XCD":
		return "Dollar", "Cent"
	case "XOF":
		return "Franc", "Centime"
	case "XPF":
		return "Franc", "N/A"
	case "YER":
		return "Rial", "Fils"
	case "ZAR":
		return "Rand", "Cents"
	case "ZMW":
		return "Kwacha", "Ngwee"
	case "ZWL":
		return "Dollar", "Cents"
	default:
		return currency, ""
	}
}
