package main

import "fmt"

type CategoryCountMessage struct {
	Category      string
	Count         int
	TerminalError interface{}
}

func processCategories(categories []string, outChan chan<- CategoryCountMessage) {
	defer func() {
		if arg := recover(); arg != nil {
			fmt.Println(arg)
			outChan <- CategoryCountMessage{
				TerminalError: arg,
			}
			close(outChan)
		}
	}()
	channel := make(chan ChannelMessage, 10)
	go Products.TotalPriceAsync(categories, channel)
	for message := range channel {
		if message.CategoryError == nil {
			outChan <- CategoryCountMessage{
				Category: message.Category,
				Count:    int(message.Total),
			}
		} else {
			panic(message.CategoryError)
		}
	}
	close(outChan)
}
func main() {
	categories := []string{"Watersports", "Chess", "Running"}
	channel := make(chan CategoryCountMessage)
	go processCategories(categories, channel)
	for message := range channel {
		if message.TerminalError == nil {
			fmt.Println(message.Category, "Total:", message.Count)
		} else {
			fmt.Println("A terminal error occured")
		}
	}
}

// func main() {
// 	// Adding Running here creates ambigous result
// 	// categories := []string{"Watersports", "Chess", "Running"}
// 	// for _, cat := range categories {
// 	// 	total, err := Products.TotalPrice(cat)
// 	// 	if err == nil {
// 	// 		fmt.Println(cat, "Total:", ToCurrency(total))
// 	// 	} else {
// 	// 		fmt.Println(cat, "(no such category)", err)
// 	// 	}
// 	// }

// 	// Panic recovery func:
// 	defer func() {
// 		if arg := recover(); arg != nil {
// 			if err, ok := arg.(error); ok {
// 				fmt.Println("Error:", err.Error())
// 				// Panic after recovery
// 				panic(err)
// 			} else if str, ok := arg.(string); ok {
// 				fmt.Println("Message:", str)
// 			} else {
// 				fmt.Println("Panic recovered")
// 			}
// 		}
// 	}()

// 	categories := []string{"Watersports", "Chess", "Running"}
// 	channel := make(chan ChannelMessage, 10)
// 	go Products.TotalPriceAsync(categories, channel)
// 	for message := range channel {
// 		if message.CategoryError == nil {
// 			fmt.Println(message.Category, "Total:", ToCurrency(message.Total))
// 		} else {
// 			// fmt.Println(message.Category, "(no such category)", message.CategoryError)

// 			// Unrecoverable errors
// 			panic(message.CategoryError)
// 		}
// 	}

// }
