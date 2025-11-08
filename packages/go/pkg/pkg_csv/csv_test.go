package pkgcsv_test

import (
	"testing"

	pkgcsv "github.com/vtuanjs/saas_management/packages/go/pkg/pkg_csv"
)

func TestParseCSVBool(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name          string
		input         string
		expectedValue bool
		expectedOk    bool
	}{
		{
			name:          "Valid TRUE value",
			input:         "TRUE",
			expectedValue: true,
			expectedOk:    true,
		},
		{
			name:          "Valid FALSE value",
			input:         "FALSE",
			expectedValue: false,
			expectedOk:    true,
		},
		{
			name:          "Empty string should parse as false",
			input:         "",
			expectedValue: false,
			expectedOk:    true,
		},
		{
			name:          "Invalid value - lowercase true",
			input:         "true",
			expectedValue: false,
			expectedOk:    false,
		},
		{
			name:          "Invalid value - lowercase false",
			input:         "false",
			expectedValue: false,
			expectedOk:    false,
		},
		{
			name:          "Invalid value - mixed case True",
			input:         "True",
			expectedValue: false,
			expectedOk:    false,
		},
		{
			name:          "Invalid value - mixed case False",
			input:         "False",
			expectedValue: false,
			expectedOk:    false,
		},
		{
			name:          "Invalid value - numeric 1",
			input:         "1",
			expectedValue: false,
			expectedOk:    false,
		},
		{
			name:          "Invalid value - numeric 0",
			input:         "0",
			expectedValue: false,
			expectedOk:    false,
		},
		{
			name:          "Invalid value - yes",
			input:         "yes",
			expectedValue: false,
			expectedOk:    false,
		},
		{
			name:          "Invalid value - no",
			input:         "no",
			expectedValue: false,
			expectedOk:    false,
		},
		{
			name:          "Invalid value - random string",
			input:         "random",
			expectedValue: false,
			expectedOk:    false,
		},
		{
			name:          "Valid value - whitespace only (becomes empty after trim)",
			input:         "   ",
			expectedValue: false,
			expectedOk:    true,
		},
		{
			name:          "Valid value - TRUE with spaces",
			input:         " TRUE ",
			expectedValue: true,
			expectedOk:    true,
		},
		{
			name:          "Valid value - FALSE with spaces",
			input:         " FALSE ",
			expectedValue: false,
			expectedOk:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			value, ok := pkgcsv.ParseCSVBool(tt.input)

			if value != tt.expectedValue {
				t.Errorf("ParseCSVBool(%q) value = %v, want %v", tt.input, value, tt.expectedValue)
			}

			if ok != tt.expectedOk {
				t.Errorf("ParseCSVBool(%q) ok = %v, want %v", tt.input, ok, tt.expectedOk)
			}
		})
	}
}

func TestParseCSVBool_Constants(t *testing.T) {
	t.Parallel()

	t.Run("Using CsvBoolTrue constant", func(t *testing.T) {
		value, ok := pkgcsv.ParseCSVBool(pkgcsv.CsvBoolTrue)
		if !value || !ok {
			t.Errorf("ParseCSVBool(CsvBoolTrue) = (%v, %v), want (true, true)", value, ok)
		}
	})

	t.Run("Using CsvBoolFalse constant", func(t *testing.T) {
		value, ok := pkgcsv.ParseCSVBool(pkgcsv.CsvBoolFalse)
		if value || !ok {
			t.Errorf("ParseCSVBool(CsvBoolFalse) = (%v, %v), want (false, true)", value, ok)
		}
	})
}
