package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Person struct {
	firstname string
	lastname  string
}
type Info struct {
	job string
	age int
}

var persons = map[Person]Info{}

func register(c *gin.Context) {
	firstname := c.PostForm("firstname")
	if firstname == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "firtname is required"})
		return
	}
	lastname := c.PostForm("lastname")
	if lastname == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "lastname is required"})
		return
	}
	p := Person{firstname: firstname, lastname: lastname}
	_, ok := persons[p]
	if ok {
		c.JSON(http.StatusConflict, gin.H{"message": fmt.Sprintf("%s %s registered before", firstname, lastname)})
		return
	}

	job := c.DefaultPostForm("job", "Unknown")
	ageStr := c.DefaultPostForm("age", "18")
	age, err := strconv.Atoi(ageStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "age should be integer"})
		return
	}
	persons[p] = Info{job: job, age: age}
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("%s %s registered successfully", firstname, lastname)})
}

func hello(c *gin.Context) {
	firstname := c.Param("firstname")
	lastname := c.Param("lastname")
	p := Person{firstname: firstname, lastname: lastname}
	info , ok := persons[p]
	if !ok {
		c.String(http.StatusNotFound, "%s %s is not registered", firstname, lastname)
		return
	}
	c.String(http.StatusOK, "Hello %s %s; Job: %s; Age: %d", firstname, lastname, info.job, info.age)	
}
