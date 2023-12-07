package sayajim

import "fmt"

func main() {
	fmt.Println(YouIsSayajim(701, "cabo"))
}

func YouIsSayajim(pointPower int32, planet string) bool {
	if pointPower > 700 && (planet == "terra" || planet == "vegeta") {
		return true
	}
	return false
}
