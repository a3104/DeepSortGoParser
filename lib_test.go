package DeepSortGoParser

import (
	"fmt"
	"testing"
)

func TestCreateDeepSort(t *testing.T) {
	a, error := CreateByLine("368,110.61227191652152,60.857036926171105,141.4988353230947,151.90456305208812,0.0,-2.3347693962743024,0.018604981375242913,-114.70112103062752,331.0")
	if error != nil {
		t.Errorf("error must be nil %v", error)
	} else {
		fmt.Println(error)
	}
	if a.Identity != 331.0 {
		t.Errorf("ID must be 331.0 but %f", a.Identity)
	}
	correct := DeepSort{FrameNo: 368, PosX: 110.61227191652152, PosY: 60.857036926171105,
		PosW: 141.4988353230947, PosH: 151.90456305208812,
		Category: 0, Identity: 331}

	if a != correct {
		t.Errorf("%v but %v", a, correct)
	}

	b, error := CreateByLine("368.0,110.61227191652152,60.857036926171105,141.4988353230947,151.90456305208812,0.0,-2.3347693962743024,0.018604981375242913,-114.70112103062752,331.0")
	if error == nil {
		t.Errorf("error must be not nil %v", error)
	} else {
		fmt.Printf("%v\n", error)
	}
	fmt.Println(b)

	c, error := CreateByLine("")
	if error == nil {
		t.Errorf("error must be not nil %v", error)
	} else {
		fmt.Printf("%v\n", error)
	}
	fmt.Println(c)

}
