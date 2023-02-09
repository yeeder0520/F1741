package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type customer struct {
	UserName      string  `json:"username"`
	ShipTo        address `json:"shipto"`
	PurchaseOrder order   `json:"order"`
}

type address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
	ZipCode int    `json:"zipcode"`
}

type order struct {
	TotalPrice  int    `json:"total"`
	IsPaid      bool   `json:"paid"`
	Fragile     bool   `json:"fragil,omitempty"`
	OrderDetail []item `json:"orderdetail"`
}

type item struct {
	Name        string `json:"itemname"`
	Description string `json:"desc,omitempty"`
	Quantity    int    `json:"qty"`
	Price       int    `json:"price"`
}

func (c *customer) AddItem(name, description string, quantity, price int) {
	i := item{name, description, quantity, price}
	c.PurchaseOrder.OrderDetail = append(c.PurchaseOrder.OrderDetail, i)
}

func (c *customer) Total() {
	price := 0
	for _, item := range c.PurchaseOrder.OrderDetail {
		price += item.Quantity * item.Price
	}
	c.PurchaseOrder.TotalPrice = price
}

func main() {
	jsonData := []byte(`
{
	"username":"blackhat",
	"shipto":{
		"street":"Sulphur Springs Rd",
		"city":"Park City",
		"state":"VA",
		"zipcode":12345
	},
	"order":{
		"paid":false,
		"orderdetail":[]
	}
}
	`)

	if !json.Valid(jsonData) {
		fmt.Printf("JSON is not valid: %s", jsonData)
		os.Exit(1)
	}

	c := customer{}
	if err := json.Unmarshal(jsonData, &c); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	c.AddItem("A Guide to the World of zeros and ones", "book", 3, 50)
	c.AddItem("Final Fantasy The Zodiac Age", "Nintendo Switch Game", 1, 50)
	c.AddItem("Crystal Drinking Glass", "", 11, 25)
	c.Total()

	c.PurchaseOrder.IsPaid = true
	c.PurchaseOrder.Fragile = true

	customerOrder, err := json.MarshalIndent(c, "", "    ")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(string(customerOrder))
}
