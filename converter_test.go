package currency

import (
	"testing"
)

func TestConvertEndpoint(t *testing.T) {
	// Use a real API key for testing
	converter := &Converter{API_KEY: "YOUR_API_KEY"}

	result, err := converter.Convert("USD", 10, "EUR")
	if err != nil {
		t.Fatalf("Convert() returned an error: %v", err)
	}

	// Check some keys exist
	if _, ok := result["to_currency"]; !ok {
		t.Errorf("Expected to_currency in result")
	}
	if _, ok := result["to_amount"]; !ok {
		t.Errorf("Expected to_amount in result")
	}
	if _, ok := result["amount_in_word"]; !ok {
		t.Errorf("Expected amount_in_word in result")
	}

	t.Logf("Conversion result: %+v", result)
}

func TestAmountToWords(t *testing.T) {
	tests := []struct {
		amount   float64
		currency string
		want     string
	}{
		{123.45, "USD", "One Hundred And Twenty Three Dollars And Forty Five Cents"},
		{1000, "EUR", "One Thousand Euro"},
	}

	for _, tt := range tests {
		got := AmountToWords(tt.amount, tt.currency)["amount_in_word"].(string)
		if got != tt.want {
			t.Errorf("AmountToWords(%v, %v) = %v; want %v", tt.amount, tt.currency, got, tt.want)
		}
	}
}
