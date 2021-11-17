// pdf2jpg convert a pdf file to one or multiple jpg files using MuPDF library
// Copyright (C) 2021 Brahim Hamdouni - <brahim@hamdouni.com>

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published
// by the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.

// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package main

import (
	"errors"
	"flag"
	"fmt"
	"image/jpeg"
	"os"

	"github.com/gen2brain/go-fitz"
)

func main() {

	in := flag.String("in", "", "-in filename")
	out := flag.String("out", "out%02d.jpg", "-out pattern")
	qual := flag.Int("qual", 10, "-qual number (1 to 100)")
	flag.Parse()

	err := process(*in, *out, *qual)
	if err != nil {
		fmt.Printf("error: %s\n", err)
		flag.Usage()
	}
}

func process(input, output string, quality int) error {

	if input == "" {
		return errors.New("no pdf file provided")
	}

	if quality < 1 || quality > 100 {
		return errors.New("qual must be in 1 to 100")
	}

	doc, err := fitz.New(input)
	if err != nil {
		return err
	}

	defer doc.Close()
	for n := 0; n < doc.NumPage(); n++ {
		img, err := doc.Image(n)
		if err != nil {
			return err
		}

		f, err := os.Create(fmt.Sprintf(output, n))
		if err != nil {
			return err
		}

		err = jpeg.Encode(f, img, &jpeg.Options{Quality: quality})
		if err != nil {
			return err
		}
		f.Close()
	}
	return nil
}
