package main

import (
	"fmt"
	"image"
	"image/color"
	"io"
	"os"
	"strings"

	"golang.org/x/tour/pic"
	"golang.org/x/tour/reader"
)

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.
func (i IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", i[0], i[1], i[2], i[3])
}

func task1() {
	fmt.Printf("\n## Task 1 ##\n")
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}

type ErrNegativeSqrt float64

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	return x, nil
}

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func task2() {
	fmt.Printf("\n## Task 2 ##\n")
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}

// ? task 3
// ? Add a Read([]byte) (int, error) method to MyReader.
// ? Implement a Reader type that emits an infinite stream of the ASCII character 'A'.
type MyReader struct{}

func (r MyReader) Read(b []byte) (int, error) {
	// fmt.Println(b)
	for i := range b {
		b[i] = 'A'
	}
	return len(b), nil
}
func task3() {
	fmt.Printf("\n## Task 3 ##\n")
	reader.Validate(MyReader{})
}

// ? Task 4: rotreader
type rot13Reader struct {
	r io.Reader
}

// ? Implement a rot13Reader that implements io.Reader and reads from an io.Reader, modifying the stream by applying the rot13 substitution cipher to all alphabetical characters.
func (r rot13Reader) Read(b []byte) (int, error) {

	n, err := r.r.Read(b)
	for i := 0; i < n; i++ {
		if (b[i] >= 'A' && b[i] < 'N') || (b[i] >= 'a' && b[i] < 'n') {
			b[i] += 13
		} else if (b[i] > 'M' && b[i] <= 'Z') || (b[i] > 'm' && b[i] <= 'z') {
			b[i] -= 13
		}
	}
	return n, err
}
func task4() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

// ? TASK 5 - Images
// ? Define your own Image type, implement the necessary methods, and call pic.ShowImage.
// ? Bounds should return a image.Rectangle, like image.Rect(0, 0, w, h).
// ? ColorModel should return color.RGBAModel.
// ? At should return a color; the value v in the last picture generator corresponds to color.RGBA{v, v, 255, 255} in this one.
type Image struct{}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 256, 256)
}
func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) At(x, y int) color.Color {
	return color.RGBA{uint8(x), uint8(y), 255, 255}
}

func task5() {
	m := Image{}
	pic.ShowImage(m)
}
