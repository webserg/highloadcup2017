package main

import (
    "container/list"
    "sync"
)

// list of indexes 'User' -> index
// index itself is mapping 'Location' -> LocationAvg[Visited_at, Birth_date, Gender, Mark]
// this index is used in request /locations/:id/avg
// 800 is User
//for e := IdxLocation[900].Front(); e != nil; e = e.Next() {
//    fmt.Println(e.Value)
//}
var IdxLocation map[int]*list.List
// TODO: try to save it into *user

var idxLocationMutex sync.RWMutex

func getIdxLocation(User int) (*list.List) {
    idxLocationMutex.RLock()
    il, ok := IdxLocation[User]
    idxLocationMutex.RUnlock()
    if !ok {
        // IdxLocation[User] was not existed, now creating. There were no visits of this user.
        il = list.New()
        idxLocationMutex.Lock()
        IdxLocation[User] = il
        idxLocationMutex.Unlock()
    }
    return il
}

func getIdxLocationLoad(User int) (*list.List) {
    il, ok := IdxLocation[User]
    if !ok {
        // IdxLocation[User] was not existed, now creating. There were no visits of this user.
        il = list.New()
        IdxLocation[User] = il
    }
    return il
}

func UpdateIdxLocation(User int, Age int, Gender rune) {
    il := getIdxLocation(User)

    for e := il.Front(); e != nil; e = e.Next() {
        idx := e.Value.(*locationsAvg)

        idx.Age = Age
        idx.Gender = Gender
    }
}
