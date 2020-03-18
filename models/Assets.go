package models

import (
	"database/sql"
	"github.com/jinzhu/gorm"
)

type Assets struct {
	Base
	gorm.Model

	Name     string `gorm:"unique_index;not null"`
	Path     string
	FullPath string
	Type     string
	Size     string
	Status   sql.NullBool
}

func GetAssetsList(page uint, pageSize uint, query interface{}, args ...interface{}) (uint, []Assets) {
	db := GetORM()
	assets := []Assets{}
	offset := 0
	if page > 0 {
		offset = int((page - 1) * pageSize)
	}
	count := uint(0)

	db.Model(&Assets{}).Where(query, args...).Count(&count)
	if count > 0 {
		db.Where(query, args...).Limit(pageSize).Offset(offset).Find(&assets)
	}
	return count, assets
}

func (asset *Assets) GetByAssetsName(assetname string) *Assets {
	db := GetORM()
	db.Where("name=?", assetname).First(&asset)
	return asset
}

func (asset *Assets) GetByID(ID uint) *Assets {
	db := GetORM()
	db.Where("ID=?", ID).First(&asset)
	return asset
}

func (asset *Assets) GetAssetsWhere(query string, args ...interface{}) *Assets {
	db := GetORM()
	db.Where(query, args...).First(&asset)
	return asset
}
