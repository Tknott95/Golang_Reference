
// REF http://web.colby.edu/danelson/cs333-programming-languages/go/go-mandelbrot-set-tutorial/
// ADDED ONTO THE LOGIC OF http://web.colby.edu/danelson/cs333-programming-languages/go/go-mandelbrot-set-tutorial/
package main

import (
	"math"
	"os"
	"flag"
	"image"
	"image/color"
	"image/png"
)

func lenComplex(z complex128) float64 {
	a := real(z)
	b := imag(z)
	return math.Sqrt(a*a + b*b)
}


func attractor(a float64, b float64, max int) (int, float64) {	
	c := complex(a, b)
	iters := 0
	z := c
	len := lenComplex(c)
	for len < 10 {
		z = z*z + c
		len = lenComplex(z)		
		iters++
		if iters == max {
			return -1, len			 
		}
	}
	return iters, len
}


type mandelbrotWriter struct {
	f *os.File
	offset int64
} 



func createMandelbrotWriter(file string) mandelbrotWriter {
	var w mandelbrotWriter
	w.f, _ = os.OpenFile(file, os.O_CREATE|os.O_WRONLY, 0777)
	w.offset = 0
	return w
}


func (m *mandelbrotWriter) Write(p []uint8) (n int, err error) {
	n, _ = m.f.WriteAt(p, m.offset)  // here the underscore means discare the 2nd return value
	m.offset += int64(n)
	return n, nil
}



type Artist func(iterations int, value float64) color.RGBA


func makeArtist(size int, c1 color.RGBA, c2 color.RGBA) Artist {

	// lets prepare for a more interesting color scheme
	modification := [3]float64{ (float64(c2.R) - float64(int(c1.R)))/float64(size),
                            	(float64(c2.G) - float64(int(c1.G)))/float64(size),
                            	(float64(c2.B) - float64(int(c1.B)))/float64(size)  }

	table := make([]color.RGBA, size)
	
	// modify the colors
	for i:=0; i < size; i++ {
		table[i] = color.RGBA{ uint8(math.Floor(float64(int(c1.R)) + float64(i)*modification[0])),
									uint8(math.Floor(float64(int(c1.G)) + float64(i)*modification[1])),
									uint8(math.Floor(float64(int(c1.B)) + float64(i)*modification[2])),
									0xFF}
	}

	lnP := math.Log(2);
	tlnB := 2*math.Log(20);
	
	// here we create a function literal to use in the return block
	getId := func(iterations int, value float64) int {
			return iterations + int(math.Log(tlnB - math.Log(math.Log(value)))/lnP);
	};
	
	
	return func(iterations int, value float64) color.RGBA {
		if iterations == -1 {
			return color.RGBA{0, 0, 0, 0xFF};
		}
		index := (getId(iterations, value))*(size/100);
		if index < 0 {
			index = 0;
		}
		if index >= size {
			return color.RGBA{10, 0, 0, 0xFF};
		}
		return table[index];
	};
}


func main() {
	imgW := 200
	imgH := 200
	halfH := imgH/2
	halfW := imgW/2
	var scale float64
	var offx float64
	var offy float64
	var maxIters int
	var fOut string
	
	flag.Float64Var(&scale, "scale", 1, "The scale of the image")
	flag.Float64Var(&offx, "x", 0, "The x-offset of the image")
	flag.Float64Var(&offy, "y", 0, "The y-offset of the image")
	flag.IntVar(&maxIters, "maxIters", 400, "The maximim number of iterations")
	flag.StringVar(&fOut, "f", "mandelbrot_new.png",
					"The .png extension to save the image")
	
	flag.Parse()
		
	img1 := color.RGBA{0, 150, 200, 0xFF} // background color
	img2 := color.RGBA{0, 0, 171, 0xFF}
	
	// create our writer, image, and artist
	var w mandelbrotWriter = createMandelbrotWriter(fOut)
	var mandelbrot *image.RGBA = image.NewRGBA(image.Rect(0, 0, imgW, imgH))
	var artist Artist = makeArtist(2 << 8, img1, img2)
	
	const nCores = 4  // number of cores to use
	
	// set up some varibles for calculation
	scale *= 100
	k := 0
	step := imgW / nCores
	c := make(chan int, nCores)

	for k < imgW {
		// the goroutines
		go func(m int) {
			for i := m; i < (m + step); i++ {
				for j := 0; j < imgH; j++ {					
					// set the color
					mandelbrot.Set(i, j,
								artist(attractor(float64(i - halfW)/scale + float64(offx),
												 float64(j - halfH)/scale + float64(offy),
												 maxIters)))
				}
			}
			c <- 1  // send 1 on c (flowing in to c)
		}(k)  // we call the function with k
		
		k += step  // increment k accordingly
	}
	
	for i:=0; i < nCores; i++ {
		<-c  // receive value and discard it
	}
	
	// use the writer to write to the .png
	png.Encode(&w, mandelbrot)
	png.Encode(&w, mandelbrot)
	
}
