package main

import (
	//"encoding/json"
	//"log"
	// "io/ioutil"
	//"net/http"
	"os"

	//"strconv"
	//"fmt"

	//"github.com/gin-gonic/gin"
	"nats-stream-cqrs-frontend/controller"
)
 
func main() {
	r := controller.RegisterRoutes()
	//r.Static("/public", "./public")
	r.Run(port())


	// engine := gin.Default()

	// engine.LoadHTMLGlob("./templates/*.html")

	// engine.GET("/", func(c *gin.Context) {
	// 	log.Println("Root page is called.")
	// 	c.HTML(http.StatusOK, "index.html", gin.H{
	// 		"title": "Ecommerce Frontend",
	// 	})
	// })

	// engine.Run(port())
}

// 	engine.GET("/products", func(c *gin.Context) {
// 		productsAPI := fmt.Sprintf("http://%s:%s/api/products", storeHost, storePort)
// 		fmt.Printf("Products API: %s", productsAPI)

// 		//prods, err := readProducts("http://127.0.0.1:8080/api/products")
// 		prods, err := readProducts(productsAPI)
// 		if err != nil {

// 			c.String(http.StatusBadRequest, fmt.Sprintf("URL data error - %v", err))
// 			return
// 		}
// 		var products []Product

// 		//fmt.Printf("Data received:\n %s ", prods)

// 		err = json.Unmarshal(prods, &products)
// 		if err != nil {
// 			c.String(http.StatusInternalServerError, "Products unmarshal error")
// 			return
// 		}

// 		fmt.Printf("Products received:\n %v ", products)

// 		params := map[string]interface{}{
// 			"products": products,
// 		}

// 		c.HTML(http.StatusOK, "index.html", params)

// 	})

// 	// the hello message endpoint with JSON response from map
// 	engine.GET("/hello", func(c *gin.Context) {
// 		c.JSON(http.StatusOK, gin.H{"message": "Hello Gin Framework."})
// 	})

// 	engine.Run(port())
// }

// func readProducts(url string) ([]byte, error) {

// 	resp, err := http.Get(url)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer resp.Body.Close()

// 	return ioutil.ReadAll(resp.Body)
// }

func port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	return ":" + port
}
