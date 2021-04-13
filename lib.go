package DeepSortGoParser

import (
	"strconv"
	"strings"
)

type DeepSort struct {
	FrameNo  int
	PosX     float64
	PosY     float64
	PosW     float64
	PosH     float64
	Category float64
	//Uoverdot float64
	//Voverdot float64
	//Soverdot float64
	Identity float64
}

func CreateByLine(line string) (DeepSort, error) {
	parts := strings.Split(line, ",")
	var deepSort DeepSort

	if frameNo, error := strconv.Atoi(parts[0]); error == nil {
		deepSort.FrameNo = frameNo
	} else {
		return DeepSort{}, error
	}
	if PosX, error := strconv.ParseFloat(parts[1], 64); error == nil {
		deepSort.PosX = PosX
	} else {
		return DeepSort{}, error
	}
	if PosY, error := strconv.ParseFloat(parts[2], 64); error == nil {
		deepSort.PosY = PosY
	} else {
		return DeepSort{}, error
	}
	if PosW, error := strconv.ParseFloat(parts[3], 64); error == nil {
		deepSort.PosW = PosW
	} else {
		return DeepSort{}, error
	}
	if PosH, error := strconv.ParseFloat(parts[4], 64); error == nil {
		deepSort.PosH = PosH
	} else {
		return DeepSort{}, error
	}

	if Category, error := strconv.ParseFloat(parts[5], 64); error == nil {
		deepSort.Category = Category
	} else {
		return DeepSort{}, error
	}

	if Identity, error := strconv.ParseFloat(parts[9], 64); error == nil {
		deepSort.Identity = Identity
	} else {
		return DeepSort{}, error
	}

	return deepSort, nil
}
