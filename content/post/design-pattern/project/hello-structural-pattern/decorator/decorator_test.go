package decorator

import (
	"fmt"
	"testing"
)

func TestAfter(t *testing.T) {
	americano := &americano{}
	whipAmericano := &whip{
		beverage: americano,
	}

	latte := &latte{}
	shotLatte := &shot{
		beverage: latte,
	}
	whipShotLatte := &whip{
		beverage: shotLatte,
	}
	chipWhipShotLatte := &chip{
		beverage: whipShotLatte,
	}

	fmt.Printf("Price of americano with whip: %d\n", whipAmericano.getPrice())
	fmt.Printf("Price of latte with shot, whip, and chip: %d\n", chipWhipShotLatte.getPrice())
}
