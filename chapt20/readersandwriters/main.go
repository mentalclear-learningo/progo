package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func processData(reader io.Reader) {
	b := make([]byte, 2) // Creating a slice of 2 bytes
	for {
		count, err := reader.Read(b) // for loop is reading what is passed to the func
		// if count > 0 print string rep of the byte read.
		if count > 0 {
			Printfln("Read %v bytes: %v", count, string(b[0:count]))
		}
		// When there's nothing to read EOF error thrown, this breaks.
		if err == io.EOF {
			break
		}
	}
}

func processDataWriter(reader io.Reader, writer io.Writer) {
	b := make([]byte, 2)
	for {
		count, err := reader.Read(b)
		if count > 0 {
			writer.Write(b[0:count])
			Printfln("Read %v bytes: %v", count, string(b[0:count]))
		}
		if err == io.EOF {
			break
		}
	}
}

func processDataUtil(reader io.Reader, writer io.Writer) {
	count, err := io.Copy(writer, reader)
	if err == nil {
		Printfln("Read %v bytes", count)
	} else {
		Printfln("Error: %v", err.Error())
	}
}

func pipeReading() {
	pipeReader, pipeWriter := io.Pipe()
	go func() {
		GenerateData(pipeWriter)
		// pipeWriter.Close()  now Close() moved to the GenerateData func
	}()
	ConsumeData(pipeReader)
}

func concatMultipleReaders() {
	r1 := strings.NewReader("Kayak")
	r2 := strings.NewReader("Lifejacket")
	r3 := strings.NewReader("Canoe")
	concatReader := io.MultiReader(r1, r2, r3)
	ConsumeData(concatReader)
}

func combineMultiWriters() {
	var w1 strings.Builder
	var w2 strings.Builder
	var w3 strings.Builder
	combinedWriter := io.MultiWriter(&w1, &w2, &w3)
	GenerateData(combinedWriter)
	Printfln("Writer #1: %v", w1.String())
	Printfln("Writer #2: %v", w2.String())
	Printfln("Writer #3: %v", w3.String())
}

func echoingReadToWrite() {
	r1 := strings.NewReader("Kayak")
	r2 := strings.NewReader("Lifejacket")
	r3 := strings.NewReader("Canoe")
	concatReader := io.MultiReader(r1, r2, r3)
	var writer strings.Builder
	teeReader := io.TeeReader(concatReader, &writer)
	ConsumeData(teeReader)
	Printfln("Echo data: %v", writer.String())
}

func limitingReadData() {
	r1 := strings.NewReader("Kayak")
	r2 := strings.NewReader("Lifejacket")
	r3 := strings.NewReader("Canoe")
	concatReader := io.MultiReader(r1, r2, r3)
	limited := io.LimitReader(concatReader, 5)
	ConsumeData(limited)
}

func main() {
	// to create a reader based on a string, the strings package
	// provides a NewReader constructor function
	r := strings.NewReader("Kayak")
	Printfln("%T", r)
	//processData(r)

	var builder strings.Builder
	processDataWriter(r, &builder)
	Printfln("String builder contents: %s", builder.String())

	Printfln("\nUsing the Utility Functions for Readers and Writers:")
	var builder2 strings.Builder
	r2 := strings.NewReader("Kayak")
	processDataUtil(r2, &builder2)
	Printfln("String builder contents: %s", builder2.String())

	Printfln("\nUsing Pipes:")
	pipeReading()

	Printfln("\nConcatenating Multiple Readers:")
	concatMultipleReaders()

	Printfln("\nCombining Multiple Writers:")
	combineMultiWriters()

	Printfln("\nEchoing Reads to a Writer:")
	echoingReadToWrite()

	Printfln("\nLimiting Read Data:")
	limitingReadData()

	Printfln("\nBuffering Data:")
	usingReaderWrapper()

	Printfln("\nUsing the Additional Buffered Reader Methods:")
	usingAdditionalBuffReadMethods()

	Printfln("\nPerforming Buffered Writes:")
	perfBuffWrites()

	Printfln("\nScanning Values from a Reader:")
	scanningValues()

	Printfln("\nWriting Formatted Strings to a Writer:")
	writeFormattedMain()

	Printfln("\nUsing a Replacer with a Writer:")
	usingReplacer()
}

func usingReplacer() {
	text := "It was a boat. A small boat."
	subs := []string{"boat", "kayak", "small", "huge"}
	var writer strings.Builder
	writeReplaced(&writer, text, subs...)
	fmt.Println(writer.String())
}

func writeReplaced(writer io.Writer, str string, subs ...string) {
	replacer := strings.NewReplacer(subs...)
	replacer.WriteString(writer, str)
}

func writeFormatted(writer io.Writer, template string, vals ...interface{}) {
	fmt.Fprintf(writer, template, vals...)
}

func writeFormattedMain() {
	var writer strings.Builder
	template := "Name: %s, Category: %s, Price: $%.2f"
	writeFormatted(&writer, template, "Kayak", "Watersports", float64(279))
	fmt.Println(writer.String())
}

func scanFromReader(reader io.Reader, template string,
	vals ...interface{}) (int, error) {
	return fmt.Fscanf(reader, template, vals...)
}

func scanSingle(reader io.Reader, val interface{}) (int, error) {
	return fmt.Fscan(reader, val)
}

func scanningValues() {
	reader := strings.NewReader("Kayak Watersports $279.00")

	// With scanSingle:
	for {
		var str string
		_, err := scanSingle(reader, &str)
		if err != nil {
			if err != io.EOF {
				Printfln("Error: %v", err.Error())
			}
			break
		}
		Printfln("Value: %v", str)
	}

	// var name, category string
	// var price float64
	// scanTemplate := "%s %s $%f"
	// _, err := scanFromReader(reader, scanTemplate, &name, &category, &price)
	// if err != nil {
	// 	Printfln("Error: %v", err.Error())
	// } else {
	// 	Printfln("Name: %v", name)
	// 	Printfln("Category: %v", category)
	// 	Printfln("Price: %.2f", price)
	// }
}

func perfBuffWrites() {
	text := "It was a boat. A small boat."
	var builder strings.Builder
	//var writer = NewCustomWriter(&builder)
	// With Buffered Writer:
	var writer = bufio.NewWriterSize(NewCustomWriter(&builder), 20)
	for i := 0; true; {
		end := i + 5
		if end >= len(text) {
			writer.Write([]byte(text[i:]))
			writer.Flush()
			break
		}
		writer.Write([]byte(text[i:end]))
		i = end
	}
	Printfln("Written data: %v", builder.String())
}

func usingAdditionalBuffReadMethods() {
	text := "It was a boat. A small boat."
	var reader io.Reader = NewCustomReader(strings.NewReader(text))
	var writer strings.Builder
	slice := make([]byte, 5)
	buffered := bufio.NewReader(reader)
	for {
		count, err := buffered.Read(slice)
		if count > 0 {
			Printfln("Buffer size: %v, buffered: %v",
				buffered.Size(), buffered.Buffered())
			writer.Write(slice[0:count])
		}
		if err != nil {
			break
		}
	}
	Printfln("Read data: %v", writer.String())
}

func usingReaderWrapper() {
	text := "It was a boat. A small boat."
	var reader io.Reader = NewCustomReader(strings.NewReader(text))
	var writer strings.Builder
	slice := make([]byte, 7)
	// Adding a buffer:
	reader = bufio.NewReader(reader)
	for {
		count, err := reader.Read(slice)
		if count > 0 {
			writer.Write(slice[0:count])
		}
		if err != nil {
			break
		}
	}
	Printfln("Read data: %v", writer.String())
}
