package main

import (
	"encoding/json"
	"fmt"
	"math"
	"math/rand"
	"os"
	"strings"
	"time"
)

var (
	random = map[string]func()interface{}{
		"string": randomString,
		"int": randomInt,
		"int32": randomInt32,
		"int64": randomInt64,
		"float32": randomFloat32,
		"float64": randomFloat64,
		"bool": randomBool,
	}
	letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

const (
	maxInt8 = 255
	maxInt32 = 2147483647
	maxInt64 = 9223372036854775807
	decimalCount = 2
)


func main()  {
	argsWithoutProg := os.Args[1:]
	a := argsWithoutProg[0]

	a = strings.Replace(a, "\t", " ", -1)
	a = strings.Replace(a, "\n", " ", -1)

	rand.Seed(time.Now().UnixNano())

	var key string
	var value string
	keyTurn := true
	start := false
	sm := make(map[string]string)
	for i, v := range a {
		s := string(v)
		if s == "{" {
			start = true
			continue
		}
		if start {
			if s == "}" {
				break
			}
			if s == " " {
				continue
			}
			if keyTurn {
				key = key + s
				if string(a[i+1]) == " " {
					keyTurn = false
					continue
				}
			}
			if !keyTurn {
				value = value + s
				if string(a[i+1]) == " " {
					sm[key] = value
					key = ""
					value = ""
					keyTurn = true
				}
			}
		}
	}

	r := fill(sm)
	rb, err := json.Marshal(r)
	if err != nil {
		return
	}
	fmt.Println(string(rb))
}

func randomString() interface{} {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, 10)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func randomInt() interface{} {
	return rand.Intn(maxInt8)
}

func randomInt32() interface{} {
	return rand.Int31n(maxInt32)
}

func randomInt64() interface{} {
	return rand.Int63n(maxInt64)
}

func randomFloat64() interface{} {
	f := math.Pow10(decimalCount)
	return math.Round(rand.Float64()*f)/f
}

func randomFloat32() interface{} {
	f := math.Pow10(decimalCount)
	return math.Round(float64(rand.Float32())*f)/f
}

func randomBool() interface{} {
	if time.Now().Unix() % 2 == 0 {
		return true
	}
	return false
}

func randomSlice(t string) [2]interface{} {
	f := random[t]
	return [2]interface{}{f(), f()}
}


func fill(ll map[string]string) map[string]interface{} {
	r := make(map[string]interface{})
	for k, v := range ll {
		if z, ok := random[k]; ok {
			k = fmt.Sprintf("%v", z())
		}
		f, ok := random[v]
		if ok {
			r[k] = f()
		}
		if !ok && (v[:2] == "[]") {
			r[k] = randomSlice(v[2:])
		}
		if !ok && (v[:3] == "map") {
			var to string
			var tt string
			m := make(map[string]string)
			for i, l := range v[4:] {
				s := string(l)
				if s == "]" {
					tt = v[i+5:]
					break
				}
				to = to + s
			}
			m[to] = tt
			r[k] = fill(m)
		}
	}
	return r
}


