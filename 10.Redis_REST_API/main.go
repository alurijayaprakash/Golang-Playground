package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
)

const (
	PORT = ":8080"
)

type Student struct {
	Id         string  `json:"id"`
	Name       string  `json:"name"`
	Gpa        float64 `json:"gpa"`
	IsEligible bool    `json:"isEligible"`
}

type redisCache struct {
	Host    string
	Db      int
	Expires time.Duration
}

var ctx = context.Background()
var RCache *redisCache

func NewRedisCache(host string, db int, exp time.Duration) *redisCache {
	return &redisCache{
		Host:    host,
		Db:      db,
		Expires: exp,
	}
}

func (cache *redisCache) getClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     cache.Host,
		Password: "",
		DB:       cache.Db,
	})
}

func (cache *redisCache) Set(key string, st *Student) error {
	client := cache.getClient()

	// serialize Student object to JSON
	json, err := json.Marshal(st)
	if err != nil {
		fmt.Println("Unable marshal json ", err)
		return err
	}

	setErr := client.Set(ctx, key, json, cache.Expires*time.Second).Err()
	if err != nil {
		fmt.Println("Unable to set the key into Redis Cache: setErr ", setErr)
		return err
	}
	return err
}

func (cache *redisCache) Get(key string) (*Student, error) {
	client := cache.getClient()
	// fmt.Println("Data... ", cache, client)
	val, err := client.Get(ctx, key).Result()
	if err != nil {
		fmt.Println("Unable to get the key from Redis Cache: err ", err)
		return nil, err
	}

	// serialize JSON to Student object
	stobj := &Student{}
	err = json.Unmarshal([]byte(val), stobj)
	if err != nil {
		fmt.Println("Unable Unmarshal stobj ", err)
		panic(err)
	}

	return stobj, nil
}

//////////////////////////////CACHE END//////////////////////////////////////////////

func GetStudent(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetStudent .... invoked")
	params := mux.Vars(r)
	studentID := params["id"]
	stuData, getErr := RCache.Get(studentID)
	if getErr != nil {
		fmt.Println("Key not found in cache")
	}
	fmt.Println("GetStudent Func Sending => ", stuData)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stuData)

}

func SetStudent(w http.ResponseWriter, r *http.Request) {
	fmt.Println("SetStudent .... invoked")
	w.Header().Set("Content-Type", "application/json")
	var stuObj Student
	_ = json.NewDecoder(r.Body).Decode(&stuObj) // must pass pointer
	err := RCache.Set(stuObj.Id, &stuObj)
	if err != nil {
		fmt.Println("Unable to Set : Error getErr")
	}

}

func InitializeRouter() {
	routerX := mux.NewRouter()

	routerX.HandleFunc("/getstudent/{id}", GetStudent).Methods("GET")
	routerX.HandleFunc("/setstudent", SetStudent).Methods("POST")

	fmt.Println("Starting Server at ", PORT)
	log.Fatal(http.ListenAndServe(PORT, routerX))
}

func InitializeRedis() {
	RCache = NewRedisCache("localhost:6379", 0, 5*time.Minute)

	println("Redis-Cache created...! ", RCache)

	s1 := Student{
		Id:         "1",
		Name:       "Prakash",
		Gpa:        3.9,
		IsEligible: false,
	}

	RCache.Set(s1.Id, &s1)
	fmt.Println("Basic Data inserted with ID", s1.Id)
	firstVal, err := RCache.Get(s1.Id)
	if err != nil {
		fmt.Println("Unable to retrive data from Redis ", err)
	}
	fmt.Println("Basic Data retrived ", firstVal)

}

func main() {
	println("Hello, welcome to Redis-Cache")

	InitializeRedis()
	InitializeRouter()

}
