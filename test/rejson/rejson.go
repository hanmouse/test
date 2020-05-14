package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"

	goredis "github.com/go-redis/redis"
	"github.com/gomodule/redigo/redis"
	"github.com/nitishm/go-rejson"
)

// Name - student name
type Name struct {
	First  string `json:"first,omitempty"`
	Middle string `json:"middle,omitempty"`
	Last   string `json:"last,omitempty"`
}

// Student - student object
type Student struct {
	Name Name `json:"name,omitempty"`
	Rank int  `json:"rank,omitempty"`
}

func Example_JSONSet(rh *rejson.Handler) {

	student := &Student{
		Name: Name{
			"Mark",
			"S",
			"Pronto",
		},
		Rank: 1,
	}

	/*
		student2 := &Student{
			Name: Name{
				"Mark2",
				"S2",
				"Pronto2",
			},
			Rank: 2,
		}
	*/

	var res interface{}
	var err error

	res, err = rh.JSONSet("student", ".", student)
	if err != nil {
		log.Fatalf("Failed to JSONSet: err=%#v", err.Error())
		return
	}

	if res.(string) == "OK" {
		fmt.Printf("JSONSet Success: %s\n", res)
	} else {
		fmt.Println("Failed to Set")
	}

	/*
			res, err = rh.JSONDel("student", ".")
			if err != nil {
				log.Fatalf("Failed to JSONDel: err=%#v", err.Error())
				return
			}

		if res.(int64) > 0 {
			fmt.Printf("JSONDel Success: %#v\n", res.(int64))
		} else {
			fmt.Printf("Failed to Del: res=%#v\n", res.(int64))
		}
	*/

	studentJSON, err := redis.Bytes(rh.JSONGet("student2", "."))
	if err != nil {
		log.Fatalf("Failed to JSONGet")
		return
	}

	readStudent := Student{}
	err = json.Unmarshal(studentJSON, &readStudent)
	if err != nil {
		log.Fatalf("Failed to JSON Unmarshal")
		return
	}

	fmt.Printf("Student read from redis : %#v\n", readStudent)
}

func main() {
	var addr = flag.String("Server", "localhost:6379", "Redis server address")

	rh := rejson.NewReJSONHandler()
	flag.Parse()

	/*
		// Redigo Client
		conn, err := redis.Dial("tcp", *addr)
		if err != nil {
			log.Fatalf("Failed to connect to redis-server @ %s", *addr)
		}

		log.Println("Connected to Redis server")

		defer func() {
			_, err = conn.Do("FLUSHALL")
			err = conn.Close()
			if err != nil {
				log.Fatalf("Failed to communicate to redis-server @ %v", err)
			}
		}()

		rh.SetRedigoClient(conn)

		fmt.Println("Executing Example_JSONSET for Redigo Client")

		Example_JSONSet(rh)
	*/

	// GoRedis Client
	cli := goredis.NewClient(&goredis.Options{Addr: *addr})
	defer func() {
		if err := cli.FlushAll().Err(); err != nil {
			log.Fatalf("goredis - failed to flush: %v", err)
		}
		if err := cli.Close(); err != nil {
			log.Fatalf("goredis - failed to communicate to redis-server: %v", err)
		}
	}()
	rh.SetGoRedisClient(cli)
	fmt.Println("\nExecuting Example_JSONSET for GoRedis Client")
	Example_JSONSet(rh)
}
