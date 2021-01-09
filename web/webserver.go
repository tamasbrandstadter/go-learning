package main

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type User struct {
	name, password string
	age            int
}

func sayName(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
	} else {
		for k, v := range r.Form {
			fmt.Println("key:", k)
			fmt.Println("val:", strings.Join(v, ""))
		}
		// send data to client side
		if _, err := fmt.Fprintf(w, "Hello %s!", r.Form.Get("name")); err != nil {
			log.Println(err)
		}
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		timeNow := time.Now().Unix()
		hash := md5.New()
		_, err := io.WriteString(hash, strconv.FormatInt(timeNow, 10))
		token := fmt.Sprintf("%x", hash.Sum(nil))

		t, _ := template.ParseFiles("web/template/login.gtpl")
		err = t.Execute(w, token)
		if err != nil {
			log.Println(err)
		}
	} else {
		token := r.FormValue("token")
		if token != "" {
			// check token validity
		} else {
			fmt.Println("No token")
		}

		name := r.FormValue("username")
		password := r.FormValue("password")
		age := r.FormValue("age")

		nameEmpty := len(name) == 0
		pwEmpty := len(password) == 0
		ageEmpty := len(age) == 0

		if nameEmpty {
			log.Println("Name can't be empty")
		} else if pwEmpty {
			log.Println("Password can't be empty")
		} else if ageEmpty {
			log.Println("Age can't be empty")
		} else {
			age, err := strconv.Atoi(r.Form.Get("age"))
			if err != nil {
				log.Println("Can't parse age as it is not a valid number")
			} else {
				if age > 100 {
					log.Println("Age too high")
				} else {
					u := User{
						name:     name,
						password: password,
						age:      age,
					}
					log.Printf("Created new user with name %s\n", u.name)
				}
			}
		}
	}
}

func main() {
	// set router
	http.HandleFunc("/", sayName)
	http.HandleFunc("/login", login)

	// set listen port
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
