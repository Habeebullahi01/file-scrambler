package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
)

func main() {
	f := flag.String("f", "", "The path to the file to be scrambled")
	u := flag.Bool("u", false, "To Scramble or not to Scramble.")
	flag.Parse()

	log.Default()
	log.Println("This is File Scrambler")

	// sfp := strings.Split(*f, strconv.QuoteRune(os.PathSeparator))
	fp := filepath.Dir(*f)
	fn := filepath.Base(*f)
	// println(sfp)
	println("File Path: " + fp)
	println("File Name: " + fn)

	// Read file
	// (1) Ensure file exists
	if _, err := os.ReadFile(*f); os.IsNotExist(err) || *f == "" {
		log.Panicln("Please specify a valid file path.")
	}
	// Ensure path does not lead to a directory

	// filepath.WalkDir(filepath.Dir(*f), func(path string, d fs.DirEntry, err error) error {

	// })

	fileContent, _ := os.ReadFile(*f)
	// Scramble by changing it's 0s and 1s
	var newFile []byte
	if *u {

		newFile = unscramble(fileContent)
	} else {
		newFile = scramble(fileContent)

	}

	// Save

	var p = filepath.Join(fp, "s", fn)
	os.Mkdir(filepath.Join(fp, "s"), os.ModePerm)
	if err := os.WriteFile(p, newFile, 0600); err != nil {
		log.Println("An error occured")
		log.Println(err.Error())
	} else {
		// index += 1
		// log.Printf("File %d done.", index)
		println("Scrambled file saved.")
	}

	os.Exit(0)
}

func scramble(con []byte) []byte {
	var newCon []byte
	for _, c := range con {
		// println(c)
		c += 1
		newCon = append(newCon, c)
	}
	fb := newCon[0]
	lb := newCon[len(newCon)-1]

	// newCon = append(newCon, lb)
	println("First byte: ", fb)
	println("Last byte: ", lb)
	return newCon
}

func unscramble(con []byte) []byte {
	var newCon []byte
	for _, c := range con {
		// println(c)
		c -= 1
		newCon = append(newCon, c)
	}
	return newCon
}
