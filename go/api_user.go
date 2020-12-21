/*
 * simple blog
 *
 * A Simple Blog
 *
 * API version: 1.0.0
 * Contact: apiteam@swagger.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/boltdb/bolt"
	"github.com/dgrijalva/jwt-go"
)

/*
func SignIn(w http.ResponseWriter, r *http.Request) {
	db, err := bolt.Open("my.db", 0600, nil)
	fatal(err)
	defer db.Close()

	var user User
	//err = json.NewDecoder(r.Body).Decode(&user)
	///*
		temp := json.NewDecoder(r.Body)
		t, err := temp.Token()
		if err != nil {
			response := ErrorResponse{"aa"} //omit

			//response := InlineResponse404{err.Error()}
			JsonResponse(response, w, http.StatusBadRequest)
			return
		}
		fmt.Printf("%T: %v\n", t, t)
		err = temp.Decode(&user)
	///
	rrr, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response := ErrorResponse{"bb"} //omit

		//response := InlineResponse404{err.Error()}
		JsonResponse(response, w, http.StatusBadRequest)
		return
	}

	fmt.Fprintln(w, len(string(rrr)))

	err = json.NewDecoder(r.Body).Decode(&user)
	/////////////////////////////////////////
	if err != nil {
		response := ErrorResponse{"hhh"} //omit

		//response := InlineResponse404{err.Error()}
		JsonResponse(response, w, http.StatusBadRequest)
		return
	}

	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("User"))
		if b != nil {
			v := b.Get([]byte(user.Username))
			if ByteSliceEqual(v, []byte(user.Password)) {
				return nil
			} else {
				return errors.New("Username and Password do not match")
			}
		} else {
			return errors.New("Username and Password do not match")
		}
	})

	if err != nil {
		response := InlineResponse404{err.Error()}
		JsonResponse(response, w, http.StatusNotFound)
		return
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()
	claims["iat"] = time.Now().Unix()
	token.Claims = claims

	if err != nil {
		fatal(err)
	}

	tokenString, err := token.SignedString([]byte(user.Username))
	if err != nil {
		fatal(err)
	}

	response := Token{tokenString}
	JsonResponse(response, w, http.StatusOK)
}

func ByteSliceEqual(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	if (a == nil) != (b == nil) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}*/

func SignIn(w http.ResponseWriter, r *http.Request) {
	db, err := bolt.Open("my.db", 0600, nil)
	fatal(err)
	defer db.Close()

	u, err := url.Parse(r.URL.String())
	fatal(err)
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(m)
	var user User
	user.Username = m["username"][0]
	user.Password = m["password"][0]
	fmt.Println(user.Username, user.Password)

	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("User"))
		if b != nil {
			v := b.Get([]byte(user.Username))
			if ByteSliceEqual(v, []byte(user.Password)) {
				return nil
			} else {
				return errors.New("Username and Password do not match")
			}
		} else {
			return errors.New("Username and Password do not match")
		}
	})

	if err != nil {
		response := InlineResponse404{err.Error()}
		JsonResponse(response, w, http.StatusNotFound)
		return
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()
	claims["iat"] = time.Now().Unix()
	token.Claims = claims

	if err != nil {
		fatal(err)
	}

	tokenString, err := token.SignedString([]byte(user.Username))
	if err != nil {
		fatal(err)
	}

	response := Token{tokenString}
	JsonResponse(response, w, http.StatusOK)
}

func ByteSliceEqual(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	if (a == nil) != (b == nil) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
