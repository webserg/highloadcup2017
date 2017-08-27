package main

import (
//    "fmt"
    "sync"
)

type visit_update struct {
    Id          * int
    Location    * int
    User        * int
    Mark        * int
    Visited_at  * int
}

type visit struct {
    Id          int
    Location    int
    User        int
    Mark        int
    Visited_at  int
}

type visit1 struct {
    Id          int
    Location    int
    User        int
    Mark        int
    Visited_at  int
}

var visits map[int]*visit
var visitsMutex sync.RWMutex

const visitsMaxCount = 10000740+40000
var visitsCount int
//var visits1[visitsMaxCount+1]visit1
var visits1[1]visit1

func getVisit(Visit int) (*visit, bool) {
    visitsMutex.RLock()
    l, err := visits[Visit]
    visitsMutex.RUnlock()
    return l, err
}

func insertRawVisit(Visit int, v * visit_update) {
    visitsMutex.Lock()
    var vn visit
    visits[Visit] = &vn
    vn.Id = Visit
    vn.Location = *v.Location
    vn.User = *v.User
    vn.Mark = *v.Mark
    vn.Visited_at = *v.Visited_at
    visitsMutex.Unlock()
}
