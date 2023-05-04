package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

const (
	host   = "localhost"
	port   = "5431"
	dbName = "prueba"

	rolName     = "postgres"
	rolPassword = "postgres"
)

func main() {
	dbConn, err := establishDbConnection()
	fmt.Println("establishDbConnection dbConn:", dbConn, ",err:", err)

	/*_ , err = dbConn.Exec(`INSERT into items(code, customer_name, order_date, product, quantity, price) values
												(124, 'pedro', 1683211111, 'fideos con tuco', '20', 600)`)
	fmt.Println("INSERT err:", err)*/
	//QueryRow

	now := time.Now()
	fmt.Println(now, now.Unix())

	var id int
	row := dbConn.QueryRow(`INSERT into items(code, customer_name, order_date, product, quantity, price) values
												(125, $2, $1, 'fideos con pesto', '5', 800) RETURNING ID`, now.Unix(), "manolo")
	err = row.Scan(&id)
	fmt.Println("INSERT err:", err, ", id:", id)

	_ , err = dbConn.Exec("UPDATE items set product = 'polenta' where id = 1;")
	fmt.Println("UPDATE err:", err)

	printAllItems(dbConn)
}

func establishDbConnection() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, rolName, rolPassword, dbName)
	dbConn, err := sql.Open("postgres", psqlInfo)
	return dbConn, err
}

func printAllItems(dbConn *sql.DB) {
	query := `
		select id, code, customer_name, order_date, product, quantity, price
		from items`

	rows, err := dbConn.Query(query)
	if (err != nil) {
		fmt.Println(err)
		return
	}

	var items []Item = make([]Item, 0)
	for rows.Next() {
		var item Item
		err := rows.Scan(&item.Id, &item.Code, &item.CustomerName, &item.OrderDate, &item.Product, &item.Quantity, &item.Price)
		if (err != nil) {
			fmt.Println(err)
			continue
		}
		items = append(items,item)
	}

	fmt.Printf("items: %+v\n", items)
}

type Item struct {
	Id int
	Code int
  CustomerName string
  OrderDate int
  Product string
  Quantity int
  Price int
}