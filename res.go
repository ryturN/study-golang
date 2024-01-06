package main

import (
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

type person struct {
	ID   string `json:"Id"`
	Name string `json:"Name"`
	Age  int    `json:"Age"`
}

var persons = []person{
	{
		ID:   "1",
		Name: "Ryan",
		Age:  14,
	},
	{
		ID:   "2",
		Name: "Ganteng",
		Age:  23,
	},
	{
		ID:   "3",
		Name: "Keren",
		Age:  12,
	},
}

func getPersons(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, persons)
}

func getPersonsById(c *gin.Context) {
	id := c.Param("id")

	for _, p := range persons {
		if p.ID == id {
			c.IndentedJSON(http.StatusOK, p)
			return
		}
	}
}

func updatePersons(c *gin.Context) {
	id := c.Param("id")
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	name := string(body)

	for _, p := range persons {
		if p.ID == id {
			p.Name = name
			c.IndentedJSON(http.StatusOK, p)
			return
		}
	}
}

func newPerson(c *gin.Context) {
	var newPerson person
	if err := c.BindJSON(&newPerson); err != nil {
		return
	}
	persons = append(persons, newPerson)
	c.IndentedJSON(http.StatusCreated, newPerson)
}
func main() {
	router := gin.Default()
	router.GET("/persons", getPersons)
	router.POST("/persons", newPerson)
	router.GET("/persons/:id", getPersonsById)
	router.PUT("/persons/:id", updatePersons)
	router.Run("localhost:8080")
}
