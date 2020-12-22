package credit_test

import (
	"fmt"

	"github.com/apkraft/bgo_homework_1_3/pkg/credit"
)

func ExampleCalculate() { // имя функции - Example + имя проверяемой функции
	fmt.Println(credit.Calculate(20, 3, 1_000_000))
	fmt.Println(credit.Calculate(15, 2, 10_000_000))
	// Output:
	// 3716358 33788888 133788888
	// 48486648 163679552 1163679552
}
