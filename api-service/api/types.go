package api

import (
	"github.com/gin-gonic/gin"
	"github.com/go-neutrino/neutrino-core/api-service/db"
	"github.com/go-neutrino/neutrino-core/api-service/utils"
	"net/http"
	"github.com/go-neutrino/neutrino-core/api-service/notification"
	"github.com/go-neutrino/neutrino-core/models"
)

type TypesController struct {
}

func (t *TypesController) CreateTypeHandler(c *gin.Context) {
	body := utils.GetBody(c)
	typeName := body["name"]

	app := c.MustGet("app").(models.JSON)

	d := db.NewAppsDbService(c.MustGet("user").(string))
	d.UpdateId(app["_id"],
		models.JSON{
			"$push": models.JSON{
				"types": typeName,
			},
		},
	)
}

func (t *TypesController) DeleteType(c *gin.Context) {
	appId := c.Param("appId")
	typeName := c.Param("typeName")

	app := c.MustGet("app").(models.JSON)

	d := db.NewAppsDbService(c.MustGet("user").(string))
	d.UpdateId(app["_id"],
		models.JSON{
			"$pull": models.JSON{
				"types": typeName,
			},
		},
	)

	database := db.NewTypeDbService(appId, typeName)
	session, collection := database.GetCollection()
	defer session.Close()

	dropError := collection.DropCollection()

	if dropError != nil {
		RestError(c, dropError)
		return
	}
}

func (t *TypesController) InsertInTypeHandler(c *gin.Context) {
	appId := c.Param("appId")
	typeName := c.Param("typeName")
	body := utils.GetBody(c)

	d := db.NewTypeDbService(appId, typeName)
	err := d.Insert(body)

	if err != nil {
		RestError(c, err)
		return
	}

	notification.Notify(notification.Build(
		notification.OP_CREATE,
		notification.ORIGIN_API,
		body,
		nil,
	))

	RespondId(body["_id"], c)
}

func (t *TypesController) GetTypeDataHandler(c *gin.Context) {
	appId := c.Param("appId")
	typeName := c.Param("typeName")

	app := c.MustGet("app").(models.JSON)
	types := app["types"].([]interface{})
	found := false

	for _, t := range types {
		if value, ok := t.(string); ok && value == typeName {
			found = true
			break
		}
	}

	if !found {
		RestErrorNotFound(c)
		return
	}

	d := db.NewTypeDbService(appId, typeName)

	typeData, err := d.Find(nil, nil)

	if err != nil {
		RestError(c, err)
		return
	}

	c.JSON(http.StatusOK, typeData)
}

func (t *TypesController) GetTypeItemById(c *gin.Context) {
	appId := c.Param("appId")
	typeName := c.Param("typeName")
	itemId := c.Param("itemId")

	d := db.NewTypeDbService(appId, typeName)

	item, err := d.FindId(itemId, nil)

	if err != nil {
		RestError(c, err)
		return
	}

	c.JSON(http.StatusOK, item)
}

func (t *TypesController) UpdateTypeItemById(c *gin.Context) {
	appId := c.Param("appId")
	typeName := c.Param("typeName")
	itemId := c.Param("itemId")

	d := db.NewTypeDbService(appId, typeName)
	body := utils.GetBody(c)

	err := d.UpdateId(itemId, body)

	if err != nil {
		RestError(c, err)
		return
	}
}

func (t *TypesController) DeleteTypeItemById(c *gin.Context) {
	appId := c.Param("appId")
	typeName := c.Param("typeName")
	itemId := c.Param("itemId")

	d := db.NewTypeDbService(appId, typeName)

	err := d.RemoveId(itemId)

	if err != nil {
		RestError(c, err)
		return
	}
}