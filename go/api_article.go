/*
 * simple blog
 *
 * API version: 1.0.0
 * Contact: apiteam@swagger.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/boltdb/bolt"
)

// Error types
var ErrIndex = errors.New("页面索引出错！")
var ErrArticle = errors.New("文章不存在！")
var ErrDelete = errors.New("删除文章失败！")

func GetArticleById(w http.ResponseWriter, r *http.Request) {
	//connect to database
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//  /user/article/{id}
	articleId := strings.Split(r.URL.Path, "/")[4]

	//	string to int
	Id, err := strconv.Atoi(articleId)
	fmt.Println(Id)
	if err != nil {
		reponse := ErrorResponse{"ArticleId错误！"}
		JsonResponse(reponse, w, http.StatusBadRequest)
		return
	}

	//query the article by ID
	var article Article
	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Article"))
		if b != nil {
			v := b.Get(itob(Id))
			if v == nil {
				return ErrArticle
			} else {
				_ = json.Unmarshal(v, &article)
				return nil
			}
		} else {
			return ErrArticle
		}
	})

	if err != nil {
		response := ErrorResponse{err.Error()}
		JsonResponse(response, w, http.StatusNotFound)
		return
	}

	JsonResponse(article, w, http.StatusOK)
}

//  /user/articles
//  http://localhost:8080/user/articles?page=1
func GetArticles(w http.ResponseWriter, r *http.Request) {
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	u, err := url.Parse(r.URL.String())
	if err != nil {
		log.Fatal(err)
	}
	m, _ := url.ParseQuery(u.RawQuery)
	page := m["page"][0]
	IdIndex, err := strconv.Atoi(page)

	pageCount := 0
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Article")) //这个桶必须存在
		b.ForEach(func(k, v []byte) error {
			pageCount = pageCount + 1
			return nil
		})
		return nil
	})

	fmt.Println(pageCount)

	//display 10 articles per page
	IdIndex = (IdIndex-1)*10 + 1
	var articles ArticlesResponse
	var article ArticleResponse
	err = db.View(func(tx *bolt.Tx) error {
		articles.PageCount = pageCount
		b := tx.Bucket([]byte("Article"))
		var k, v []byte
		if b != nil {
			c := b.Cursor()
			k, v = c.First()
			err = json.Unmarshal(v, &article)
			for i := 1; i < IdIndex; i++ {
				k, v = c.Next()
				err = json.Unmarshal(v, &article)
				fmt.Println(article.Id)
				if k == nil {
					return ErrIndex
				}
			}
			count := 0
			for ; k != nil && count < 10; k, v = c.Next() {
				err = json.Unmarshal(v, &article)
				if err != nil {
					return err
				}
				articles.Articles = append(articles.Articles, article)
				count = count + 1
			}
			return nil
		} else {
			return ErrArticle
		}
	})
	if err != nil {
		response := ErrorResponse{err.Error()}
		JsonResponse(response, w, http.StatusNotFound)
		return
	}
	json, err := json.Marshal(articles)
	fmt.Println(string(json))
	JsonResponse(articles, w, http.StatusOK)
}

func DeleteArticleById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("call deletr")
	//connect to database
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//  /user/article/{id}
	articleId := strings.Split(r.URL.Path, "/")[4]

	Id, err := strconv.Atoi(articleId)
	fmt.Println(Id)
	if err != nil {
		response := ErrorResponse{"ArticleId错误！"}
		JsonResponse(response, w, http.StatusBadRequest)
		return
	}

	//delete the article by ID
	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Article"))
		if b != nil {
			c := b.Cursor()
			c.Seek(itob(Id))
			err := c.Delete()
			if err != nil {
				return ErrDelete
			}
		} else {
			return ErrArticle
		}
		return nil
	})

	if err != nil {
		response := ErrorResponse{err.Error()}
		JsonResponse(response, w, http.StatusNotFound)
		return
	}
	JsonResponse("", w, http.StatusOK)
}



