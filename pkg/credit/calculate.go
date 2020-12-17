package credit

import (
	"math"
)

func Calculate(monthInterestRate, periodAmount, creditAmount int64) {

	multiplier := monthInterestRate * (1 + monthInterestRate) periodAmount / ((1 + monthInterestRate) periodAmount - 1),
	monthlyAnnuityPayment := multiplier * creditAmount

	return monthPayment, creditOverpayment, totalPayment
}
