package translations

import (
	"math/rand"

	address_mocks "github.com/brazzcore/mocai/pkg/mocai/entities/address/mocks/ptbr"
	company_mocks "github.com/brazzcore/mocai/pkg/mocai/entities/company/mocks/ptbr"
	gender_mocks "github.com/brazzcore/mocai/pkg/mocai/entities/gender/mocks/ptbr"
	person_mocks "github.com/brazzcore/mocai/pkg/mocai/entities/person/mocks/ptbr"
	phone_mocks "github.com/brazzcore/mocai/pkg/mocai/entities/phone/mocks/ptbr"
)

func init() {
	// Choose a random state
	state := address_mocks.States[rand.Intn(len(address_mocks.States))]

	// Gets the UF corresponding to the selected state
	uf := address_mocks.UFs[state]

	Register("ptbr", map[string]string{
		"person_first_name_male":   person_mocks.FirstNamesMale[rand.Intn(len(person_mocks.FirstNamesMale))],
		"person_first_name_female": person_mocks.FirstNamesFemale[rand.Intn(len(person_mocks.FirstNamesFemale))],
		"person_last_name":         person_mocks.LastNames[rand.Intn(len(person_mocks.LastNames))],
		"gender":                   gender_mocks.Genders[rand.Intn(len(gender_mocks.Genders))],
		"address_street":           address_mocks.Streets[rand.Intn(len(address_mocks.Streets))],
		"address_city":             address_mocks.Cities[rand.Intn(len(address_mocks.Cities))],
		"address_state":            state,
		"address_uf":               uf,
		"address_zip":              address_mocks.ZIPCodes[rand.Intn(len(address_mocks.ZIPCodes))],
		"phone_area_code":          phone_mocks.AreaCodes[rand.Intn(len(phone_mocks.AreaCodes))],
		"company_name":             company_mocks.CompanyNames[rand.Intn(len(company_mocks.CompanyNames))],

		// errors
		"invalid_certificate":                  "certidão inválida",
		"invalid_vital_records_office_number":  "número do cartório de registros civis inválido",
		"invalid_archive_number":               "número do arquivo inválido",
		"invalid_vital_records_service number": "número de serviço de registros vitais inválido",
		"invalid_birth_year":                   "ano de nascimento inválido",
		"invalid_book_number":                  "número do livro inválido",
		"invalid_page_number":                  "número da página inválida",
		"invalid_term_number":                  "número do termo inválido",
		"invalid_number_without_check_digits":  "número sem dígitos de verificação inválido",
		"error_generating_brazilian_company":   "erro ao gerar empresa brasileira",
		"invalid_cnpj":                         "cnpj inválido",
		"no_company_names_available":           "nenhum nome de empresa disponível",
		"invalid_cpf":                          "cpf inválido",
		"no_data_available_for_genders":        "não há dados disponíveis para este gênero",
		"error_converting_digit":               "erro ao converter dígito",
		"error_generating_person":              "erro ao gerar pessoa",
		"no_data_available_for_first_names":    "não há dados disponíveis para primeiro nome",
		"no_data_available_for_last_names":     "não há dados disponíveis para último nome",
		"error_generating_phone":               "erro ao gerar número de telefone",
		"no_data_available_for_area_codes":     "não há código de área disponível",
	})
}
