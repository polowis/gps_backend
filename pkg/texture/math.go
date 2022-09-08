package texture


/*
Return reflect
*/
func reflect(target int) string {
	if target < 0 {
		return "-"
	}
	return "+"
}

func findYIntercept(slope int, target Point) int {
	slopeFlag := reflect(slope)
	var result int
	if slopeFlag == "-" {
		result = target.y - slope * target.x
	} else {
		result = target.y + slope * target.x
	}
	return result
}

/*
Use linear to map pixel
*/
func useLinearEquation(source Point, dest Point) (int, int) {
	// y = mx + b
	var slope int
	if ( dest.x - source.x ) == 0 {
		slope = 0
	} else {
		slope = (dest.y - source.y) / (dest.x - source.x)
	}

	interceptY := findYIntercept(slope, dest)
	return slope, interceptY
}