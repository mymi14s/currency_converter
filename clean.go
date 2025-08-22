package currency

// CleanData removes unnecessary fields and renames keys.
func CleanData(data map[string]interface{}) map[string]interface{} {
	data["from_currency"] = data["base_code"]
	data["to_currency"] = data["target_code"]
	data["to_amount"] = data["conversion_result"]

	for _, value := range []string{
		"documentation", "terms_of_use", "base_code", "target_code", "conversion_result"} {
		delete(data, value)
	}

	return data
}
