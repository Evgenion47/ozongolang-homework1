package chess

import "errors"

type Knight struct {
	x int
	y int
}

func CoordValidation(coord string) bool {
	if len(coord) == 2 {
		coordX := coord[0] - 96
		coordY := coord[1] - 48
		if coordX <= 8 && coordX >= 1 && coordY <= 8 && coordY >= 1 {
			return true
		}
	}
	return false
}

func CoordParse(str string) Knight {
	return Knight{int(str[0]) - 96, int(str[1]) - 48}
}

func IntAbs(a int) int {
	if a < 0 {
		return a * -1
	}
	return a
}

func CanKnightAttack(white, black string) (bool, error) {
	if !CoordValidation(white) || !CoordValidation(black) || black == white {
		return false, errors.New("Wrong coords")
	}
	wCoords := CoordParse(white)
	bCoords := CoordParse(black)
	dX := IntAbs(wCoords.x - bCoords.x)
	dY := IntAbs(wCoords.y - bCoords.y)
	if (dY == 2 && dX == 1) || (dY == 1 && dX == 2) {
		return true, nil
	}
	return false, nil
}
