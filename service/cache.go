package main

import (
	"log"
	"sync"
	"time"
)

type t_data struct {
	value *user
	save bool
	time time.Time
}

type t_cache struct {
	idMap map[uint64]*t_data
	loginMap map[string]*t_data
}

var (
	cache t_cache
	mutex sync.Mutex
)

func init() {
	cache = t_cache{}
	cache.idMap = map[uint64]*t_data{}
	cache.loginMap = map[string]*t_data{}
	mutex = sync.Mutex{}
}

func (cache t_cache)update(index uint64, age uint8) error {
	if err := cache.checkId(index); err != nil {
		return err
	}
	mutex.Lock()
	cache.idMap[index].value.Age = age
	cache.idMap[index].save = true
	mutex.Unlock()
	return nil
}

func (cache t_cache)get(index string, user *user) error {
	if err := cache.checkLogin(index); err != nil {
		return err
	}
	mutex.Lock()
	*user = *cache.loginMap[index].value
	mutex.Unlock()
	return nil
}

func (cache t_cache)checkId(id uint64) error {
	mutex.Lock()
	_, exist := cache.idMap[id]
	mutex.Unlock()
	if exist {
		return nil
	}
	user := user{}
	if err := db.C("users").FindId(id).One(&user); err != nil {
		return err
	}
	cache.add(&user)
	return nil
}

func (cache t_cache)checkLogin(login string) error {
	mutex.Lock()
	_, exist := cache.loginMap[login]
	mutex.Unlock()
	if exist {
		return nil
	}
	user := user{}
	if err := db.C("users").Find(obj{"login": login}).One(&user); err != nil {
		return err
	}
	cache.add(&user)
	return nil
}

func (cache t_cache)add(user *user){
	data := t_data{}
	data.value = user
	data.save = false
	data.time = time.Now()
	mutex.Lock()
	cache.idMap[user.Id] = &data
	cache.loginMap[user.Login] = &data
	mutex.Unlock()
}


func (cache t_cache)remove(user *user) {
	mutex.Lock()
	delete(cache.idMap, user.Id)
	delete(cache.loginMap, user.Login)
	mutex.Unlock()
}


func (cache t_cache)worker_save() {
	for ;; {
		tmp := []user{}
		for _, v := range cache.idMap {
			if v.save {
				tmp = append(tmp, *v.value)
				v.save = false
			}
		}
		for i := 0; i < len(tmp); i++ {
			if err := db.C("users").UpdateId(tmp[i].Id, tmp[i]); err != nil {
				log.Println(err)
			}
		}
		time.Sleep(10 * time.Second)
	}
}

func (cache t_cache)worker_clear() {
	for ;; {
		for _, v := range cache.idMap {
			if time.Now().Sub(v.time).Seconds() > 10 {
				cache.remove(v.value)
			}
		}
		time.Sleep(120 * time.Second)
	}
}


// (check flag to save) worker save to db
// worker clean up from cache