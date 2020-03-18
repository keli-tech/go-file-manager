package assets

import (
	"github.com/labstack/echo/v4"
	"go-file-manager/handler"
	"go-file-manager/models"
	"strconv"
)

func Index(c echo.Context) (err error) {
	//nowatime := time.Now().Nanosecond()

	path := c.QueryParam("path")
	if path == "" {
		path = "/"
	}

	count, assets := models.GetAssetsList(uint(1), uint(100), "path = ?", path)
	type Data struct {
		List  []models.Assets
		Total uint
	}
	return handler.SuccessResponse(c, "success", Data{assets, count})

}

func Delete(c echo.Context) (err error) {

	assetIDVal := c.Param("assetID")
	assets := new(models.Assets)
	assetID, _ := strconv.Atoi(assetIDVal)

	assets = assets.GetByID(uint(assetID))

	if err := assets.Delete(&assets).Error; err == nil {
		return handler.SuccessMsgResponse(c, "删除成功")
	} else {
		return handler.SuccessMsgResponse(c, "删除失败，请重试")
	}

}
