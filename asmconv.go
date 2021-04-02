package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	readPtr := flag.String("f", "", "filename to read")
	writePtr := flag.String("o", "output.asm", "filename to write")
	flag.Parse()
	fmt.Println("Reading:", *readPtr)
	fmt.Println("Writing:", *writePtr)

	// Open file and create scanner on top of it
	file, err := os.Open(*readPtr)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	//open output file
	newFile, err := os.Create(*writePtr)
	if err != nil {
		log.Fatal(err)
	}
	defer newFile.Close()

	var pi byte
	//read file header
	for i := 0; i < 129; i++ {
		err = binary.Read(file, binary.LittleEndian, &pi)
		if err != nil {
			log.Fatal(err)
		}
	}

	//loop until end of file
	newline := 1
	for {
		err = binary.Read(file, binary.LittleEndian, &pi)
		if err != nil {
			if err == io.EOF {
				fmt.Println("Finished")
				return
			} else {
				log.Fatal(err)
			}
		}

		s := ""
		switch pi {
		case 255:
			return
		case 13:
			s = fmt.Sprintln("")
			//setup variable to indicate the next character will be a hex line number byte
			newline = 1
		default:
			if newline > 0 {
				//line numbers are stored as two hex digits
				s = fmt.Sprintf("%02x", pi)
				newline += 1
				if newline == 3 {
					newline = 0
				}
			} else {
				//print a normal ASCII character
				s = fmt.Sprintf("%c", pi)
			}
		}

		n, err := newFile.WriteString(s)
		if n < len(s) || err != nil {
			log.Println("Error writing to output file")
			return
		}
	}
}
