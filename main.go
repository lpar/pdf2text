package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"math"
	"os"

	"github.com/lpar/pdf"
)

func main() {
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	tmpfile, err := ioutil.TempFile("", "pdf2text")
	if err != nil {
		panic(err)
	}
	defer func() {
		os.Remove(tmpfile.Name())
	}()
	n, err := tmpfile.Write(data)
	if err != nil {
		panic(err)
	}
	if n < len(data) {
		panic(fmt.Errorf("wrote %d but expected to write %d", n, len(data)))
	}
	tmpfile.Close()
	err = dumpText(tmpfile.Name())
	if err != nil {
		panic(err)
	}
	return
}

func dumpText(path string) error {
	f, r, err := pdf.Open(path)
	defer f.Close()
	if err != nil {
		return err
	}
	numpages := r.NumPage()
	var buf bytes.Buffer
	for i := 1; i <= numpages; i++ {
		page := r.Page(i)
		if page.V.IsNull() {
			continue
		}
		texts := page.Content().Text
		var nx float64
		var ny float64
		for _, text := range texts {
			if text.Y != ny {
				buf.Write([]byte("\n"))
			}
			ny = text.Y
			c := text.S[0]
			dx := math.Abs(text.X - nx)
			if dx >= 1.0 {
				buf.Write([]byte(" "))
			}
			nx = text.W + text.X
			if c < 32 || c > 126 {
				buf.Write([]byte(" "))
			} else {
				buf.Write([]byte(text.S))
			}
		}
		fmt.Printf("%s", buf.String())
		buf.Reset()
	}
	return nil
}
