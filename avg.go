package avg

import (
	"fmt"
)

type Kind int

const (
	Dot = Kind(iota)
	Line
	Curve
	Circle
	Square
)

type Point struct {
	x, y int
}

type Shape struct {
	kind Kind
	pts []Point
}

func (k Kind) String() string {
	switch k {
	case Dot: return "Dot"
	case Line: return "Line"
	case Curve: return "Curve"
	case Circle: return "Circle"
	case Square: return "Square"
	default: return "Unknown"
	}
}

func (p Point) String() string {
	return fmt.Sprintf("%d:%d", p.x, p.y)
}

func (s Shape) String()  string {
	return fmt.Sprintf("%v%v", s.kind, s.pts)
}

var (
	index [256]int
	cmap = "0123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghjklmnpqrstuvwxyz"
)

func mk(kind Kind, pts []Point) *Shape {
	return &Shape{kind:kind, pts:pts}
}

func init() {
	for i := 0; i < 256; i++ {
		index[i] = -1
	}
	for i,c:=range cmap {
		index[c] = i
	}
}

func ToPoints(lines []string)[][]Point{
	pts := make([][]Point,0,60)
	for i, line := range lines {
		for j, c := range []byte(line) {
			k := index[c]
			if k == -1 {
				continue
			}
			if len(pts) <= int(k) {
				pts = pts[0:k+1]
			}
			pts[k] = append(pts[k], Point{i,j})
		}
	}
	return pts
}

func ToFig(lines []string) (ss []*Shape) {
	points := ToPoints(lines) 
	points = append(points, nil)
	k := 0
	var pts []Point
	for i, p := range points {
		if len(p) == 1 {
			pts = append(pts, p[0])
			continue
		}
		switch {
		case i == k+1: ss = append(ss, mk(Dot, pts))
		case i > k+1: ss =  append(ss, mk(Curve, pts))
		}
		pts = nil
		//fmt.Printf("%d: %s", i, ss)
		switch len(p) {
		case 0: break
		case 2: ss = append(ss,mk(Line, p))
		case 4:
			if  p[0].x == p[1].x && p[2].x == p[3].x &&
			    p[0].y == p[2].y && p[1].y == p[3].y {
				ss = append(ss,mk(Square, p))
			} else {
				ss = append(ss,mk(Circle, p))
			}
		default: ss = append(ss,mk(Circle, p))
		}
		k = i+1
	}
	return ss
}
