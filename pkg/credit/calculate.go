package credit

import "math"

func Calculate(interestRate, creditPeriod, creditAmount int) (int, int, int) {

	monthlyInterestRate := (float64(interestRate) / 100) / 12
	periodsAmount := creditPeriod * 12
	creditAmountInKopecks := creditAmount * 100

	annuityCoeffDividend := monthlyInterestRate * math.Pow((1+monthlyInterestRate), float64(periodsAmount))

	annuityCoeffDivisor := math.Pow((1+monthlyInterestRate), float64(periodsAmount)) - 1

	annuityCoeff := annuityCoeffDividend / annuityCoeffDivisor

	monthlyAnnuityPaymentInKopecks := int(annuityCoeff * float64(creditAmountInKopecks))

	totalPaymentInKopecks := monthlyAnnuityPaymentInKopecks * periodsAmount

	creditOverpaymentInKopecks := totalPaymentInKopecks - creditAmountInKopecks

	return monthlyAnnuityPaymentInKopecks, creditOverpaymentInKopecks, totalPaymentInKopecks
}
