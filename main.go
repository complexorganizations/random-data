package main

import (
	"bytes"
	"crypto/rand"
	"flag"
	"io"
	"log"
	"math/big"
	"os"
	"strconv"
)

var (
	assetsLocation = "assets/"
	stringLocation = assetsLocation + "string"
	byteLocation   = assetsLocation + "byte"
	intLocation    = assetsLocation + "int"
	boolLocation   = assetsLocation + "bool"
	update         bool
)

func init() {
	if len(os.Args) > 1 {
		tempUpdate := flag.Bool("update", false, "Make changes to the random data file.")
		flag.Parse()
		update = *tempUpdate
	} else {
		log.Fatalln("Error: No flags provided. Please use -help for more information.")
	}
}

func main() {
	if update {
		// Generate random string with a specified length and characters and write to file
		writeToFile(stringLocation, []byte(randomStringSpecifiedOfGivenLength(generateRandomInt(100000000))))
		// Generate random bytes and write to file
		writeToFile(byteLocation, []byte(randomBytesArray(generateRandomInt(100000000))))
		// Generate random int and write to file
		writeToFile(intLocation, []byte(generateRandomBigInt(100000000).String()))
		// Generate a random bool and write to file
		writeToFile(boolLocation, []byte(strconv.FormatBool(generateRandomBool())))
	}
}

// Generate a random string of a given length and return it.
func randomStringSpecifiedOfGivenLength(length int) string {
	// ABCDEFGHIJKLMNOPQRSTUVWXYZ, abcdefghijklmnopqrstuvwxyz, 0123456789 ~!@#$%^&*()-_+={}][|\`,./?;:'"<>
	letters := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789~!@#$%^&*()-_+={}][|\\`,./?;:'\"<>")
	randomString := make([]rune, length)
	for i := range randomString {
		randomInt, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			log.Fatalln(err)
		}
		randomString[i] = letters[randomInt.Int64()]
	}
	return string(randomString)
}

// Write to a file.
func writeToFile(path string, content []byte) {
	file, err := os.Create(path)
	if err != nil {
		log.Fatalln(err)
	}
	_, err = io.Copy(file, bytes.NewReader(content))
	if err != nil {
		log.Fatalln(err)
	}
	err = file.Close()
	if err != nil {
		log.Fatalln(err)
	}
}

// Generate a random int between 0 and a given max
func generateRandomInt(max int64) int {
	randomInt, err := rand.Int(rand.Reader, big.NewInt(max))
	if err != nil {
		log.Fatalln(err)
	}
	return int(randomInt.Int64())
}

// Generate a random byte array and return it.
func randomBytesArray(length int) []byte {
	randomBytes := make([]byte, length)
	_, err := rand.Read(randomBytes)
	if err != nil {
		log.Fatalln(err)
	}
	return randomBytes
}

// Generate a random int between 0 and a given max
func generateRandomBigInt(max int64) *big.Int {
	randomInt, err := rand.Int(rand.Reader, big.NewInt(max))
	if err != nil {
		log.Fatalln(err)
	}
	return randomInt
}

// Generate a random bool and than return it.
func generateRandomBool() bool {
	randomInt, err := rand.Int(rand.Reader, big.NewInt(int64(2)))
	if err != nil {
		log.Fatalln(err)
	}
	return int(randomInt.Int64()) == int(1)
}
