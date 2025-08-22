# Currency Converter

**Currency Converter** Go library! This tool makes it easy to convert money from one currency to another and even tells you the amount in words. It is perfect for invoices, financial apps, and international transactions.

---

## What This Tool Does

This library helps you:
- Convert an amount from one currency (like USD) to another (like EUR).
- Get the current exchange rate.
- See the converted amount written out in words (e.g., *“one hundred twenty-three euros”*).
- Work with over 160 global currencies.

All powered by the reliable [ExchangeRate-API](https://www.exchangerate-api.com/).

---

## Getting Started

### Step 1: Get Your Free API Key

1. Go to [https://www.exchangerate-api.com](https://www.exchangerate-api.com)
2. Sign up for a free account.
3. Copy your API key from your dashboard.

> The free plan allows **1,500 requests per month** — great for testing and small apps!

---

### Step 2: Install the Library

Open your terminal and run:

```bash
go get github.com/mymi14s/currency_converter
```

---

### Step 3: Write Your First Conversion

Here’s a simple program to convert $100 USD to Euros (EUR):

```go
package main

import (
    "fmt"
    "log"
    currency "github.com/mymi14s/currency_converter"

func main() {
    converter := &currency.Converter{API_KEY: "your_api_key_here"}
	result, err := converter.Convert("USD", 100, "EUR")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result)
	}
    
    if err != nil {
        log.Fatalf("Conversion failed: %v", err)
    }

    // Print the result
    fmt.Printf("%v %s = %v %s\n",
        result["from_amount"],
        result["from_currency"],
        result["to_amount"],
        result["to_currency"])

    fmt.Printf("Rate: 1 %s = %v %s\n", 
        result["from_currency"], 
        result["conversion_rate"], 
        result["to_currency"])

    fmt.Printf("In words: %s\n", result["amount_in_word"])
}
```

### Sample Output:
```
map[amount_in_word:Eighty Six Euro And Five Cents conversion_rate:0.8605 from_amount:100 from_currency:USD result:success time_last_update_unix:1.755820801e+09 time_last_update_utc:Fri, 22 Aug 2025 00:00:01 +0000 time_next_update_unix:1.755907201e+09 time_next_update_utc:Sat, 23 Aug 2025 00:00:01 +0000 to_amount:86.05 to_currency:EUR]

100 USD = 92.34 EUR
Rate: 1 USD = 0.9234 EUR
In words: ninety-two euros and thirty-four cents
```

---

## How to Use `Convert()`

### Function: `Convert(from, amount, to)`

```go
result, err := converter.Convert("USD", 50.75, "JPY")
```

| Parameter     | Example     | Description |
|---------------|-------------|-------------|
| `from`        | `"USD"`     | The currency you're converting **from** (3-letter code) |
| `amount`      | `50.75`     | The amount of money (as a number) |
| `to`          | `"JPY"`     | The currency you're converting **to** |

> Supported codes: USD, EUR, GBP, JPY, CAD, AUD, CHF, and [160+ more](https://www.exchangerate-api.com/docs/supported-currencies)

---

## What You Get Back

After conversion, you get a result with useful details:

| Field               | Example Value        | Meaning |
|---------------------|----------------------|--------|
| `from_currency`     | `"USD"`              | Original currency |
| `to_currency`       | `"EUR"`              | Target currency |
| `from_amount`       | `100.0`              | Original amount |
| `to_amount`         | `92.34`              | Converted amount |
| `conversion_rate`   | `0.9234`             | Exchange rate used |
| `amount_in_word`    | `"ninety-two euros..."` | Amount in spoken/written form |

---

## Tips & Best Practices

### Use Valid Currency Codes
Always use **3-letter ISO codes**:
- Correct: `USD`, `EUR`, `GBP`, `INR`, `AUD`
- Wrong: `usd`, `dollars`, `US`

### Handle Errors Gracefully
Network issues or invalid currencies can cause errors. Always check:

```go
if err != nil {
    fmt.Println("Oops! Could not convert:", err)
    return
}
```

### Be Mindful of Rate Limits
- Free plan: 1,500 requests/month
- Avoid rapid loops or unnecessary calls
- Consider caching common conversions (e.g., USD → EUR)

### Want Better Performance?
Add a timeout to avoid hanging:

```go
client := &http.Client{Timeout: 10 * time.Second}
// (Advanced: modify the library to use this client)
```

---

## Real-World Use Cases

| Use Case              | How This Helps |
|-----------------------|----------------|
| Invoicing software    | Show totals in client’s currency + in words |
| E-commerce stores     | Display prices in local currency |
| Travel budget apps    | Convert spending across countries |
| Accounting tools      | Automate multi-currency reports |

---

## Frequently Asked Questions

**Q: Is this free to use?**  
A: Yes! The library is free. It uses the free tier of ExchangeRate-API (1,500 requests/month).

**Q: Does it work offline?**  
A: No. It needs internet to fetch live exchange rates.

**Q: Can I convert multiple currencies at once?**  
A: Not directly. You’ll need to call `Convert()` for each pair.

**Q: What if I enter an invalid currency?**  
A: You’ll get an error like: `failed to make request` or `invalid currency`.

**Q: Are the exchange rates real-time?**  
A: Yes! Rates are updated frequently by ExchangeRate-API.

---

## 🛠 Need Help?

If something isn’t working:
1. Double-check your API key
2. Verify internet connection
3. Confirm currency codes are correct
4. Check your usage at [https://manage.exchangerate-api.com](https://manage.exchangerate-api.com)

For bugs or suggestions, please open an issue on GitHub or contact the developer.

---

## License

This project is open-source and licensed under the MIT License. Feel free to use, modify, and share it.

---

## Thank You!