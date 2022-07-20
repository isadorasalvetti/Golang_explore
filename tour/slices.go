package tour

import "golang.org/x/tour/pic"
import "fmt"

func Pic(dx, dy int) [][]uint8 {
	var ret = make([][]uint8, dy)
	for y := 0; y < dy; y++ {  // Fill array
		ret[y] = make([]uint8, dx)
		for x :=0; x < dx; x++{
			ret[y][x] = uint8(x*y)
		} 
	}

	return ret
}

func RunPic() {
	fmt.Println(Pic(6, 3))
	pic.Show(Pic)
}
