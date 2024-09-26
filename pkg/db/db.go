package db

import (
	"sync"
)

type PreSearch struct {
	ID       int
	EtpID    string
	Article  string
	Brand    string
	PartName string
}

type DB struct {
	m     sync.Mutex        //мьютекс для синхронизации доступа
	id    int               // текущее значение ID для нового заказа
	store map[int]PreSearch // БД заказов
}

func New() *DB {
	db := DB{
		id:    1, // первый номер заказа
		store: map[int]PreSearch{},
	}
	return &db
}

func (db *DB) PreSearches() []PreSearch {
	db.m.Lock()
	defer db.m.Unlock()
	var data []PreSearch
	for _, v := range db.store {
		data = append(data, v)
	}
	return data
}

func (db *DB) NewPreSearch(o PreSearch) int {
	db.m.Lock()
	defer db.m.Unlock()
	o.ID = db.id
	db.store[o.ID] = o
	db.id++
	return o.ID
}
