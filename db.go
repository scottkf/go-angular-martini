package main

import (
	"errors"
	"fmt"
	"regexp"
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
	db.Add(&Issue{Id: 2, Title: "Test 2", Body: "Test"})
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
	re := regexp.MustCompile("(?i)" + title)
	for _, v := range db.m {
		if re.MatchString(v.Title) || title == "" {
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

func (db *issuesDB) Add(i *Issue) (int, error) {
	db.Lock()
	defer db.Unlock()
	if !db.isUnique(i) {
		return 0, ErrAlreadyExists
	}
	// Get the unique ID
	db.seq++
	i.Id = db.seq
	// Store
	db.m[i.Id] = i
	return i.Id, nil
}

func (db *issuesDB) Update(i *Issue) error {
	db.Lock()
	defer db.Unlock()
	if !db.isUnique(i) {
		return ErrAlreadyExists
	}
	db.m[i.Id] = i
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

func (i *Issue) String() string {
	return fmt.Sprintf("%d - %s (%s)", i.Id, i.Title, i.Body)
}
