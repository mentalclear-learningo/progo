package main

import (
	"database/sql"
)

func main() {
	withStructToCategory()
	Printfln("\nWith placeholders:")
	withPlaceHolders()
	Printfln("\nWith single row:")
	withSingleRow()
}

type Category struct {
	Id   int
	Name string
}

type Product struct {
	Id   int
	Name string
	Category
	Price float64
}

func withSingleRow() {
	db, err := openDatabase()
	if err == nil {
		for _, id := range []int{1, 3, 10} {
			p := queryDatabaseSingleRow(db, id)
			Printfln("Product: %v", p)
		}
		db.Close()
	} else {
		panic(err)
	}
}

func queryDatabaseSingleRow(db *sql.DB, id int) (p Product) {
	row := db.QueryRow(`
        SELECT Products.Id, Products.Name, Products.Price,
                Categories.Id as Cat_Id, Categories.Name as CatName
                FROM Products, Categories
        WHERE Products.Category = Categories.Id
            AND Products.Id = ?`, id)
	if row.Err() == nil {
		scanErr := row.Scan(&p.Id, &p.Name, &p.Price,
			&p.Category.Id, &p.Category.Name)
		if scanErr != nil {
			Printfln("Scan error: %v", scanErr)
		}
	} else {
		Printfln("Row error: %v", row.Err().Error())
	}
	return
}

func withPlaceHolders() {
	db, err := openDatabase()
	if err == nil {
		for _, cat := range []string{"Soccer", "Watersports"} {
			Printfln("--- %v Results ---", cat)
			products := queryDatabaseWPlaceholders(db, cat)
			for i, p := range products {
				Printfln("#%v: %v %v %v", i, p.Name, p.Category.Name, p.Price)
			}
		}
		db.Close()
	} else {
		panic(err)
	}
}

func queryDatabaseWPlaceholders(db *sql.DB, categoryName string) []Product {
	products := []Product{}
	rows, err := db.Query(`
        SELECT Products.Id, Products.Name, Products.Price,
                Categories.Id as Cat_Id, Categories.Name as CatName
                FROM Products, Categories
        WHERE Products.Category = Categories.Id
            AND Categories.Name = ?`, categoryName)
	if err == nil {
		for rows.Next() {
			p := Product{}
			scanErr := rows.Scan(&p.Id, &p.Name, &p.Price,
				&p.Category.Id, &p.Category.Name)
			if scanErr == nil {
				products = append(products, p)
			} else {
				Printfln("Scan error: %v", scanErr)
				break
			}
		}
	} else {
		Printfln("Error: %v", err)
	}
	return products
}

func withStructToCategory() {
	// listDrivers()
	db, err := openDatabase()
	if err == nil {
		//queryDatabase(db)
		products := queryDatabaseToStructCategory(db)
		for i, p := range products {
			Printfln("#%v: %v", i, p)
		}
		db.Close()
	} else {
		panic(err)
	}
}

func queryDatabaseToStructCategory(db *sql.DB) []Product {
	products := []Product{}
	rows, err := db.Query(`
	SELECT Products.Id, Products.Name, Products.Price,
			Categories.Id as Cat_Id, Categories.Name as CatName
			FROM Products, Categories
	WHERE Products.Category = Categories.Id`)
	if err == nil {
		for rows.Next() {
			p := Product{}
			scanErr := rows.Scan(&p.Id, &p.Name, &p.Price, &p.Category.Id, &p.Category.Name)
			if scanErr == nil {
				products = append(products, p)
			} else {
				Printfln("Scan error: %v", scanErr)
				break
			}
		}
	} else {
		Printfln("Error: %v", err)
	}
	return products
}

// func queryDatabaseToStruct(db *sql.DB) []Product {
// 	products := []Product{}
// 	rows, err := db.Query("SELECT * from Products")
// 	if err == nil {
// 		for rows.Next() {
// 			p := Product{}
// 			scanErr := rows.Scan(&p.Id, &p.Name, &p.Category, &p.Price)
// 			if scanErr == nil {
// 				products = append(products, p)
// 			} else {
// 				Printfln("Scan error: %v", scanErr)
// 				break
// 			}
// 		}
// 	} else {
// 		Printfln("Error: %v", err)
// 	}
// 	return products
// }

// func queryDatabase(db *sql.DB) {
// 	rows, err := db.Query("SELECT * from Products")
// 	if err == nil {
// 		for rows.Next() {
// 			var id, category int
// 			var name string
// 			var price float64
// 			rows.Scan(&id, &name, &category, &price)
// 			Printfln("Row: %v %v %v %v", id, name, category, price)
// 		}
// 	} else {
// 		Printfln("Error: %v", err)
// 	}
// 	// if (err == nil) {
// 	//     for (rows.Next()) {
// 	//         var id, category int
// 	//         var name int
// 	//         var price float64
// 	//         scanErr := rows.Scan(&id, &name, &category, &price)
// 	//         if (scanErr == nil) {
// 	//             Printfln("Row: %v %v %v %v", id, name, category, price)
// 	//         } else {
// 	//             Printfln("Scan error: %v", scanErr)
// 	//             break
// 	//         }
// 	//     }
// 	// } else {
// 	//     Printfln("Error: %v", err)
// 	// }
// }

// func queryDatabase(db *sql.DB) {
// 	rows, err := db.Query("SELECT * from Products")
// 	if err == nil {
// 		for rows.Next() {
// 			var id, category string
// 			var name string
// 			var price string
// 			scanErr := rows.Scan(&id, &name, &category, &price)
// 			if scanErr == nil {
// 				Printfln("Row: %v %v %v %v", id, name, category, price)
// 			} else {
// 				Printfln("Scan error: %v", scanErr)
// 				break
// 			}
// 		}
// 	} else {
// 		Printfln("Error: %v", err)
// 	}
// }
