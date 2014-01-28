package main

import (
	"errors"
	"fmt"
	"sync"
)

type Issue struct {
	Title string `json:"title" form:"title"`
	Body  string `json:"body" form:"body"`
	Id    int    `json:"id" form:"id"`
}

var (
	ErrAlreadyExists = errors.New("issue already exists")
	db               DB
)

type DB interface {
	Get(id int) *Issue
	GetAll() []*Issue
	Find(title string) []*Issue
	Add(a *Issue) (int, error)
	Update(a *Issue) error
	Delete(id int)
}

type issuesDB struct {
	sync.RWMutex
	m   map[int]*Issue
	seq int
}

func init() {
	db = &issuesDB{
		m: make(map[int]*Issue),
	}
	// Fill the database
	db.Add(&Issue{Id: 1, Title: "Test 1", Body: "Test"})
	db.Add(&Issue{Id: 2, Title: "Test 2 ", Body: "Test"})
	db.Add(&Issue{Id: 3, Title: "Test 3", Body: "Test"})
}

func (db *issuesDB) GetAll() []*Issue {
	db.RLock()
	defer db.RUnlock()
	if len(db.m) == 0 {
		return nil
	}
	ar := make([]*Issue, len(db.m))
	i := 0
	for _, v := range db.m {
		ar[i] = v
		i++
	}
	return ar
}

func (db *issuesDB) Find(title string) []*Issue {
	db.RLock()
	defer db.RUnlock()
	var res []*Issue
	for _, v := range db.m {
		if v.Title == title || title == "" {
			res = append(res, v)
		}
	}
	return res
}

func (db *issuesDB) Get(id int) *Issue {
	db.RLock()
	defer db.RUnlock()
	return db.m[id]
}

func (db *issuesDB) Add(a *Issue) (int, error) {
	db.Lock()
	defer db.Unlock()
	if !db.isUnique(a) {
		return 0, ErrAlreadyExists
	}
	// Get the unique ID
	db.seq++
	a.Id = db.seq
	// Store
	db.m[a.Id] = a
	return a.Id, nil
}

func (db *issuesDB) Update(a *Issue) error {
	db.Lock()
	defer db.Unlock()
	if !db.isUnique(a) {
		return ErrAlreadyExists
	}
	db.m[a.Id] = a
	return nil
}

func (db *issuesDB) Delete(id int) {
	db.Lock()
	defer db.Unlock()
	delete(db.m, id)
}

func (db *issuesDB) isUnique(i *Issue) bool {
	for _, v := range db.m {
		if v.Title == i.Title && v.Id != i.Id {
			return false
		}
	}
	return true
}

func (a *Issue) String() string {
	return fmt.Sprintf("%d - %s (%s)", a.Id, a.Title, a.Body)
}
