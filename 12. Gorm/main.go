package main

import (
	"errors"
	"gorm.io/gorm"
)

type TLDRDBCached struct {
	Provider TLDRProvider
}

func (t *TLDRDBCached) Retrieve(key string) string {
	var tldr TLDREntity
	db := GetConnection()
	res := db.Where("Key = ?", key).First(&tldr)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		value := t.Provider.Retrieve(key)
		db.Create(&TLDREntity{Key: key, Val: value})
		return value
	}
	return tldr.Val
}

func (t *TLDRDBCached) List() []string {
	result := []string{}
	tldrs := []TLDREntity{}
	db := GetConnection()
	db.Find(&tldrs)
	for _, tldr := range tldrs {
		result = append(result, tldr.Key)
	}
	providerList := t.Provider.List()
	for _, key := range providerList {
		existed := false
		for _, res := range result {
			if key == res {
				existed = true
				break
			}
		}
		if !existed {
			result = append(result, key)
		}
	}
	return result
}

func NewTLDRDBCached(nonCachedProvider TLDRProvider) TLDRProvider {
	return &TLDRDBCached{Provider: nonCachedProvider}
}
