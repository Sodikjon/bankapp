package card

import (
	"bank/pkg/bank/types"
	"fmt"
)

func ExampleTotal() {
	fmt.Println(Total([]types.Card{
		{
			Balance: 1_000_00,
			Active:  true,
		},
	}))
	fmt.Println(Total([]types.Card{
		{
			Balance: 1_000_00,
			Active:  true,
		},
		{
			Balance: 2_000_00,
			Active:  true,
		},
	}))
	fmt.Println(Total([]types.Card{
		{
			Balance: 1_000_00,
			Active:  false,
		},
	}))
	fmt.Println(Total([]types.Card{
		{
			Balance: -1_000_00,
			Active:  true,
		},
	}))

	//Output: 100000
	// 300000
	// 0
	// 0
}

func ExamplePaymentSources() {
	cards := []types.Card{
	  {
		PAN:     "5058 xxxx xxxx 8888",
		Balance: 10_000_00,
		Active:  true,
	  },
	  {
		PAN:     "5058 xxxx xxxx 0000",
		Balance: -10_000_00,
		Active:  true,
	  },
	  {
		PAN:     "5058 xxxx xxxx 1111",
		Balance: 15_000_00,
		Active:  false,
	  },
	}
	paymentSources := PaymentSources(cards)
  
	for _, v := range paymentSources {
	  fmt.Println(v.Number)
	}
  
	//Output:  5058 xxxx xxxx 8888
  
  }