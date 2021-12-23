package categoryfunc

import "github.com/gin-gonic/gin"
import (
	DataModel "WebNote/datamodel"
	DatabaseHandler "WebNote/db"
	Responses "WebNote/responses"
	"github.com/google/uuid"
	"net/http"
	"time"
)

var id string

func PostCategory(c *gin.Context) {
	id = uuid.New().String()
	var createCategoryJSON DataModel.CategoryJSON
	c.BindJSON(&createCategoryJSON)

	var categoryQuerry DataModel.Category
	DatabaseHandler.DB.Find(&categoryQuerry, "`Name` = ? ", createCategoryJSON.Name)
	if categoryQuerry.Name == "" {
		DatabaseHandler.DB.Create(&DataModel.Category{
			ID:       id,
			Name:     createCategoryJSON.Name,
			Tag:      createCategoryJSON.Tag,
			Author:   createCategoryJSON.Author,
			DateTime: createCategoryJSON.DateTime,
		})
		c.IndentedJSON(http.StatusCreated, Responses.EntityCreatedSuccessFullyJSON{
			ID:       id,
			Message:  "Category created successfully",
			DateTime: time.Now().String(),
		})

	} else {
		c.IndentedJSON(http.StatusConflict, Responses.EntityAllreadyExistJSON{
			ID:       categoryQuerry.ID,
			Message:  "Category with the Name " + categoryQuerry.Name + " allready exists!",
			DateTime: time.Now().String(),
		})

	}

}

func GetCategorys(c *gin.Context) {
	var category []DataModel.Category
	DatabaseHandler.DB.Find(&category)
	var categorysJSON []DataModel.CategoryJSON
	for _, category := range category {

		DatabaseHandler.DB.First(&category, "id = ?", category.ID)
		var buffercategory = DataModel.CategoryJSON{
			ID:       category.ID,
			Name:     category.Name,
			Tag:      category.Tag,
			Author:   category.Author,
			DateTime: category.DateTime,
		}

		categorysJSON = append(categorysJSON, buffercategory)
	}
	c.IndentedJSON(http.StatusOK, categorysJSON)
}
