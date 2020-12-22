/*
 * simple blog
 *
 * API version: 1.0.0
 */

package test

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"

	// Change this to your own fully-qualified import path
	sw "github.com/ServiceComputing-Project/Server/go"

	"github.com/boltdb/bolt"
)


func CreateTable() {
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Article"))
		if b == nil {
			//create table if not exits
			b, err = tx.CreateBucket([]byte("Article"))
			if err != nil {
				log.Fatal(err)
			}
		}
		if b != nil {
			var article sw.Article
			var tags []sw.Tag
			tags = append(tags, sw.Tag{"CS"})
			tags = append(tags, sw.Tag{"SC"})

			filePath := "./data"
			files, err := ioutil.ReadDir(filePath)
			if err != nil {
				log.Fatal(err)
			}
			for i := 1; i <= len(files); i++ {
				path := filePath + "/" + strconv.Itoa(i)
				fileInfoList, err := ioutil.ReadDir(path)
				var articleName string
				for i := 0; i < len(fileInfoList); i++ {
					if fileInfoList[i].IsDir() == false {
						articleName = fileInfoList[i].Name()
					}
				}
				if err != nil {
					log.Fatal(err)
				}
				content, err := ioutil.ReadFile(path + "/" + articleName)
				if err != nil {
					fmt.Println("获取文章内容失败！", err)
					return err
				}
				title := articleName[:len(articleName)-3]
				article = sw.Article{int32(i), title, tags, "2019", string(content)}
				v, err := json.Marshal(article)
				//insert rows
				err = b.Put(itob(i), v)
				if err != nil {
					log.Fatal(err)
				}
			}
		} else {
			return errors.New("列表文章不存在！")
		}

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
}

func GetArticleById(id int) {
	//connect to database
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//query the article by ID
	var article sw.Article
	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Article"))
		if b != nil {
			v := b.Get(itob(id))
			if v == nil {
				fmt.Println(id, "文章不存在！")
				return errors.New("文章不存在！")
			} else {
				_ = json.Unmarshal(v, &article)
				return nil
			}
		} else {
			fmt.Println("文章不存在！")
			return errors.New("文章不存在！")
		}
	})
}

func DeleteArticleById(id int) {
	//connect to database
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//delete the article by ID
	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Article"))
		if b != nil {
			c := b.Cursor()
			c.Seek(itob(id))
			err := c.Delete()
			if err != nil {
				return errors.New("删除文章出错！")
			}
		} else {
			fmt.Println("文章不存在！")
			return errors.New("文章不存在！")
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("Successfully Delete article ", id)
}

func GetArticles(p int) {
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//display 10 articles per page
	IdIndex := (p-1)*10 + 1
	var articles sw.ArticlesResponse
	var article sw.ArticleResponse
	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Article"))
		if b != nil {
			c := b.Cursor()
			k, v := c.Seek(itob(IdIndex))
			if k == nil {
				fmt.Println("页面索引出错！")
				return errors.New("页面索引出错！")
			}
			key := binary.BigEndian.Uint64(k)
			fmt.Print(key)
			if int(key) != IdIndex {
				fmt.Println("页面索引出错！")
				return errors.New("页面索引出错！")
			}
			count := 0
			var ori_artc sw.Article
			for ; k != nil && count < 10; k, v = c.Next() {
				err = json.Unmarshal(v, &ori_artc)
				if err != nil {
					return err
				}
				article.Id = ori_artc.Id
				article.Name = ori_artc.Name
				articles.Articles = append(articles.Articles, article)
				count = count + 1
			}
			return nil
		} else {
			return errors.New("文章不存在！")
		}
	})
	for i := 0; i < len(articles.Articles); i++ {
		fmt.Println(articles.Articles[i])
	}
}

func itob(value int) []byte {
	bytes := make([]byte, 8)
	binary.BigEndian.PutUint64(bytes, uint64(value))
	return bytes
}

func DBTestArticle() {
	fmt.Println()
	fmt.Println("DBTestArticle")
	CreateTable()
	GetArticleById(1)
	GetArticleById(5)
	GetArticles(1)
	CreateUser()
}
