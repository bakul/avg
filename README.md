# avg
Ascii vector graphics

Inspired by:
<http://cocoamine.net/blog/2015/03/20/replacing-photoshop-with-nsstring/>

Usage:

        import "github.com/bakul/avg/avg"

        pic := []string{
                "0..0...1....",
                ".....1...1..",
                "0..0...1....",
                "2..3....A..A",
                ".7....4....9",
                "6..5........",
        }
        fig := avg.ToFig(pic)

The returned value is a vector of Shapes, where each Shape
has a kind (Dot, Line, Curve, Circle, Square) and a list of
points. The example above will create a square, a circle, a
(closed) curve, a dot and a line.

You can have upto 58 shapes in a figure. The rules are as follows:

All letters except 0..9A..HJ..NP..Za..hj..np..z is ignored.
Letters O,o,I,i are not used as they are too close to 0 and 1
in some fonts.  The sequence is as shown (this A follows 9, J
follows H, a follows Z etc.). Note: this sequence is not the
same as in the referenced blog.

x direction does *down*, while the y direction goes *across*.
If one of the above letters is placed an x,y location, it
defines a point.

A dot is defined with a point.

If there are exactly 4 points defined with the same letter,
and they line up in x & y directions, it is considered a
square (or rectangle).

If there are 3 or more points with the same letter, a circle
(or ellipse) is defined.

Same letter at two points defines a line.

Else a (closed) curve is defined as a sequence of consecutive
letters.  Thus in the above example, 2-3-4-5-6-2 defines a
closed curve.

A curve or a dot may be followed or preceded a line, square
or circle but not a curve or a dot. Thus in the above example
we skipped 8 to define a dot as letter 9.
