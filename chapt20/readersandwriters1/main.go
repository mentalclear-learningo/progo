package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func processData(reader io.Reader) {
	b := make([]byte, 2)
	for {
		count, err := reader.Read(b)
		if count > 0 {
			Printfln("Read %v bytes: %v", count, string(b[0:count]))
		}
		if err == io.EOF {
			Printfln("EOF reached")
			break
		}
	}
}

func processDataWithWriter(reader io.Reader, writer io.Writer) {
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

func processDataCopy(reader io.Reader, writer io.Writer) {
	count, err := io.Copy(writer, reader)
	if err == nil {
		Printfln("Read %v bytes", count)
	} else {
		Printfln("Error: %v", err.Error())
	}
}

func noBuffRead() {
	text := "It was a boat. A small boat."
	var reader io.Reader = NewCustomReader(strings.NewReader(text))
	var writer strings.Builder
	slice := make([]byte, 5)
	// Adding buffer
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

func additionalBuffReadMeths() {
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

func buffedWrites() {
	text := "It was a boat. A small boat."
	var builder strings.Builder

	// introducing buffer:
	// var writer = NewCustomWriter(&builder)
	var writer = bufio.NewWriterSize(NewCustomWriter(&builder), 20)
	for i := 0; true; {
		end := i + 5
		if end >= len(text) {
			writer.Write([]byte(text[i:]))
			writer.Flush() // Flush the buffer here
			break
		}
		writer.Write([]byte(text[i:end]))
		i = end
	}
	Printfln("Written data: %v", builder.String())
}

func main() {
	r := strings.NewReader("Kayak")
	//processData(r)

	var builder strings.Builder
	processDataCopy(r, &builder)
	Printfln("String builder contents: %s", builder.String())

	// Pipes demo:
	pipeReader, pipeWriter := io.Pipe()
	Printfln("Tyes: %T, %T", pipeReader, pipeWriter)
	// go func() {
	// 	GenerateData(pipeWriter)
	// 	pipeWriter.Close()
	// }()

	// After improvement:
	go GenerateData(pipeWriter) // No need use Close() here any more.

	ConsumeData(pipeReader)

	// Concatenating Multiple Readers:
	Printfln("Concatenating Multiple Readers:")
	r1 := strings.NewReader("Kayak")
	r2 := strings.NewReader("Lifejacket")
	r3 := strings.NewReader("Canoe")
	concatReader := io.MultiReader(r1, r2, r3)
	ConsumeData(concatReader)

	Printfln("\nCombining Multiple Writers:")
	var w1 strings.Builder
	var w2 strings.Builder
	var w3 strings.Builder
	combinedWriter := io.MultiWriter(&w1, &w2, &w3)
	GenerateData(combinedWriter)
	Printfln("Writer #1: %v", w1.String())
	Printfln("Writer #2: %v", w2.String())
	Printfln("Writer #3: %v", w3.String())

	Printfln("\nEchoing Reads to a Writer")
	r4 := strings.NewReader("Kayak")
	r5 := strings.NewReader("Lifejacket")
	r6 := strings.NewReader("Canoe")
	concatReader1 := io.MultiReader(r4, r5, r6)
	var writer strings.Builder
	teeReader := io.TeeReader(concatReader1, &writer)
	ConsumeData(teeReader)
	Printfln("Echo data: %v", writer.String())

	Printfln("\nLimiting Read Data:")
	r7 := strings.NewReader("Kayak")
	r8 := strings.NewReader("Lifejacket")
	r9 := strings.NewReader("Canoe")
	concatReader2 := io.MultiReader(r7, r8, r9)
	limited := io.LimitReader(concatReader2, 5)
	ConsumeData(limited)

	Printfln("\nBuffering Data:")
	noBuffRead()

	Printfln("\nUsing the Additional Buffered Reader Methods:")
	additionalBuffReadMeths()

	Printfln("\nPerforming Buffered Writes:")
	buffedWrites()

	fmt.Println("\nScanning Values from a Reader:")
	scanFromReaderExample()
	fmt.Println("\nSingle")
	scanFromReaderExampleSingle()
	fmt.Println("\nWriting Formatted Strings to a Writer:")

}

func scanFromReader(reader io.Reader, template string,
	vals ...interface{}) (int, error) {
	return fmt.Fscanf(reader, template, vals...)
}

func scanFromReaderExample() {
	reader := strings.NewReader("Kayak Watersports $279.00")
	var name, category string
	var price float64
	scanTemplate := "%s %s $%f"
	_, err := scanFromReader(reader, scanTemplate, &name, &category, &price)
	if err != nil {
		Printfln("Error: %v", err.Error())
	} else {
		Printfln("Name: %v", name)
		Printfln("Category: %v", category)
		Printfln("Price: %.2f", price)
	}
}

func scanSingle(reader io.Reader, val interface{}) (int, error) {
	return fmt.Fscan(reader, val)
}

func scanFromReaderExampleSingle() {
	reader := strings.NewReader("Kayak Watersports $279.00")
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
}

func writeFormatted(writer io.Writer, template string, vals ...interface{}) {
	fmt.Fprintf(writer, template, vals...)
}
