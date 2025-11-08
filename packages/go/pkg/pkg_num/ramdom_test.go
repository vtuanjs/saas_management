package pkgnum_test

import (
	"testing"

	pkgnum "github.com/vtuanjs/saas_management/packages/go/pkg/pkg_num"
)

func TestGenerateSixDigitOtp(t *testing.T) {
	otp := pkgnum.GenerateSixDigitOtp()
	if otp < 100000 || otp > 999999 {
		t.Errorf("Generated OTP %d is out of range", otp)
	}
}
