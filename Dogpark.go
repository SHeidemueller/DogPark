package main

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

type Dog struct {
	Id    int  `json:"id"`
	Name  string `json:"name"`
	Owner string `json:"owner"`
}

var Dogs = make([]Dog, 0)
var id int = 0

func main() {

	//prefilling "Dog Park" on start up

	Dogs = append(Dogs, Dog{nextId(), "Buffy", "Susann"})
	Dogs = append(Dogs, Dog{nextId(), "Snoopy", "Charlie Brown"})

	router := gin.Default()

	router.GET("/allGet", gettingAll)
	router.GET("/oneGet/:id", gettingOne)
	router.POST("/somePost", posting)
	router.DELETE("/someDelete/:id", deleting)

	// Listen and serve on 0.0.0.0:8080
	router.Run(":8080")

}

func nextId() int {
	id += 1
	return id
}

func gettingAll(c *gin.Context) {

	//serve http://localhost:8080/allGet

	c.JSON(200, Dogs)

}

func gettingOne(c *gin.Context) {

	id := c.Param("id")
	dog_id, err := strconv.Atoi(id)

	if err != nil {
		panic(err)
	}

	content, index := getDogbyID(dog_id)

	//serve http://localhost:8080/oneGet/id
	if index == -1 {
		c.JSON(422, gin.H{"error": "index not found"})
	} else {

		c.JSON(200, content)
	}
}

//search through Dogs for correct Dog and Index
func getDogbyID(id int) (Dog, int) {
	for i, e := range Dogs {
		if e.Id == id {
			return e, i
		}
	}
	return Dog{}, -1
}

func posting(c *gin.Context) {
	// curl -i -X POST -H "Content-Type: application/json" -d "{ \"name\": \"Lassie\", \"owner\": \"Joe Carraclough\" }" http://localhost:8080/somePost
	var dog Dog
	c.Bind(&dog)

	if dog.Name != "" && dog.Owner != "" {
		dog.Id = nextId()
		Dogs = append(Dogs, dog)
		//return new dog so ID is available for client
		c.JSON(200, dog)

	} else {
		//serve if user imputs crap
		c.JSON(422, gin.H{"error": "name and owner must not be empty"})
	}
}

func deleting(c *gin.Context) {
	id := c.Param("id")

	dog_id, err := strconv.Atoi(id)

	if err != nil {
		panic(err)
	}

	//find book
	_, index := getDogbyID(dog_id)

	if index == -1 {
		c.JSON(422, gin.H{"error": "id does not exist"})
	} else {

		Dogs = append(Dogs[:index], Dogs[index+1:]...)
		c.JSON(200, "dog deleted")
	}
}
