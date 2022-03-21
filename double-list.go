package main

import (
	"math"
	"sort"
	"strconv"
	"sync"
	"time"
)

// doubleList is a structure for doubling integers
type doubleList struct {
	dblArr []int
	dblMap map[int]int
	mutex  *sync.Mutex
}

func newDoubleList() *doubleList {
	return &doubleList{
		dblArr: []int{},
		dblMap: make(map[int]int),
		mutex:  &sync.Mutex{},
	}
}

func (d *doubleList) canHandle(num int) bool {
	return abs(num) <= math.MaxInt/2
}

func (d *doubleList) add(num int) {
	if d.has(num) {
		return
	}

	d.dblArr = append(d.dblArr, num)
	sort.Ints(d.dblArr)

	go func() {
		time.Sleep(10 * time.Second)
		d.mutex.Lock()
		defer d.mutex.Unlock()
		d.dblMap[num] = num * 2
	}()
}

func (d *doubleList) has(num int) bool {
	for _, v := range d.dblArr {
		if num == v {
			return true
		}
	}

	return false
}

func (d *doubleList) list() [][2]string {
	dbls := [][2]string{}

	d.mutex.Lock()
	defer d.mutex.Unlock()

	for _, num := range d.dblArr {
		dbl, ok := d.dblMap[num]

		dblStr := "⏳"
		if ok {
			dblStr = strconv.Itoa(dbl)
		}

		dbls = append(dbls, [2]string{strconv.Itoa(num), dblStr})
	}

	return dbls
}
