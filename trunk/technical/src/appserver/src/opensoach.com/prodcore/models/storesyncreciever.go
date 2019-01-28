package models

import (
	"errors"
	"reflect"
)

func (r StoreEntityModel) GetUuid() string {
	return r.Uuid
}

func (r StoreSyncApplyRequestModel) GetDataItems() (error, []map[string]interface{}) {
	s := reflect.ValueOf(r.Data)
	if s.Kind() != reflect.Slice {
		return errors.New("error : GetDataItems() given a non-slice type"), nil
	}

	m := make([]map[string]interface{}, s.Len())

	for i := 0; i < s.Len(); i++ {
		m[i] = s.Index(i).Interface().(map[string]interface{})
	}

	return nil, m
}

func (r StoreSyncApplyRequestModel) GetDataItemsIStoreSync() (error, []IStoreSync) {
	s := reflect.ValueOf(r.Data)
	if s.Kind() != reflect.Slice {
		return errors.New("error : GetDataItems() given a non-slice type"), nil
	}

	m := make([]IStoreSync, s.Len())

	for i := 0; i < s.Len(); i++ {
		m[i] = s.Index(i).Interface().(IStoreSync)
	}

	return nil, m
}
