package controller 

import (
	// "bytes"
	 
	"os"
	// "io/ioutil"
	"log"
	"net/http" 
	// "strings"
	// "sync"
 
	"github.com/gin-gonic/gin"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"

)
func RegisterRoutes() *gin.Engine {

	r := gin.Default()

	r.LoadHTMLGlob("templates/*.html")

	r.GET("/", func(c *gin.Context) {
	 
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.POST("/", func(c *gin.Context) {

		topic := c.PostForm("topic")
		message := c.PostForm("text-message")
		//submitButton := c.PostForm("submit-btn")

		// log.Println("... Topic:", topic)
		// log.Println("... Message:", message)
		// publish message on the given topic

		err := publish(topic, message)
		if err != nil {
			c.HTML(http.StatusOK, "error.html", nil)
		} else {
			c.HTML(http.StatusOK, "index.html",
				gin.H{"Message": message, "Topic": topic})
		}	
	})
	return r
}

func publish(topic, message string) error {

	natsServer := os.Getenv("ECOMMERCE_NATS_SERVICE_HOST")
	log.Println("NAT server --- "+ natsServer)

	natsURL := "nats://"+natsServer+":4222"
	//natsURL := "nats://10.104.28.139:4222"

	log.Println("NAT server conn URL --- "+ natsURL)

	//opts := []nats.Option{nats.Name("NATS Ecommerce Publisher")}


	// Connect to NATS
	nc, err := nats.Connect(natsURL)
	if err != nil {
			return err
	}
		
	defer nc.Close()
  
	sc, err := stan.Connect("ecommerce-stan", "client-id-122", stan.NatsConn(nc)) 
	if err != nil {
		return err
	}

	defer sc.Close()

  	log.Println("... Topic:", topic)
        log.Println("... Message:", message)

	sc.Publish(topic, []byte(message))
	//nc.Flush()

	if err := nc.LastError(); err != nil {
		return err
	} else {
		log.Printf("...Published [%s] : '%s'\n", topic, message)
	}
		 
	return nil

}
