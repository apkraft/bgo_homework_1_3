package main

import "github.com/apkraft/bgo_homework_1_3/pkg/credit"

func main() {
	interestRate := 20
	creditPeriod := 3
	creditAmount := 1_000_000

	monthlyAnnuityPaymentInKopecks, creditOverpaymentInKopecks, totalPaymentInKopecks := credit.Calculate(interestRate, creditPeriod, creditAmount)

	monthlyAnnuityPayment := monthlyAnnuityPaymentInKopecks / 100
	creditOverpayment := creditOverpaymentInKopecks / 100
	totalPayment := totalPaymentInKopecks / 100

	println(monthlyAnnuityPayment, creditOverpayment, totalPayment)
}
