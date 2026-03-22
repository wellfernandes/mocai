package translations

func init() {
	Register("en_us", map[string]string{
		// certificate errors
		"invalid_certificate":                  "invalid certificate",
		"invalid_vital_records_office_number":  "invalid vital records office number",
		"invalid_archive_number":               "invalid archive number",
		"invalid_vital_records_service number": "invalid vital records service number",
		"invalid_birth_year":                   "invalid birth year",
		"invalid_book_number":                  "invalid book number",
		"invalid_page_number":                  "invalid page number",
		"invalid_term_number":                  "invalid term number",
		"invalid_number_without_check_digits":  "invalid number without check digits",

		// company errors
		"error_generating_brazilian_company": "error generating brazilian company",
		"invalid_cnpj":                       "invalid cnpj",
		"no_company_names_available":         "no company names available",

		// cpf errors
		"invalid_cpf": "invalid cpf",

		// gender errors
		"no_data_available_for_genders": "no data available for this gender",

		// national ID errors
		"error_converting_digit": "error converting digit",

		// person errors
		"error_generating_person":           "error generating person",
		"no_data_available_for_first_names": "no data available for first name",
		"no_data_available_for_last_names":  "no data available for last name",

		// phone errors
		"error_generating_phone":           "error generating phone number",
		"no_data_available_for_area_codes": "no area code available",

		// voter Registration errors
		"invalid_vote_registration": "invalid vote registration",
		"invalid_check_digit_1":     "invalid check digit 1",
		"invalid_check_digit_2":     "invalid check digit 2",
	})
}
