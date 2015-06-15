package models

import (
	"fmt"
	"htz_search/errors"
	"log"
	//	"database/sql"
)

var (
	TBL_ARTICLES = "articles"
)

type articleModel struct{}

var _articleModel = articleModel{}

func ArticleModel() *articleModel {
	return &_articleModel
}

func (*articleModel) Query(content string) ([]*Article, error) {
	sql := fmt.Sprintf("SELECT id, title,content,author,kind,ctime FROM "+TBL_ARTICLES+" WHERE content LIKE '%%%s%%'", content)
	log.Println("sql:", sql)
	rows, err := DB.Query(sql)
	if nil != err {
		log.Println(err)
		return nil, err
	}

	var as = []*Article{}
	for rows.Next() {
		var a Article
		err := rows.Scan(&a.Id, &a.Title, &a.Content, &a.Author, &a.Kind, &a.Ctime)
		if nil != err {
			log.Println(err)
			continue
		}

		r := []rune(a.Content)
		length := len(r)
		if 800 >= length {
			a.Abs = a.Content
		} else {
			a.Abs = string(r[:800]) + "..."
		}

		as = append(as, &a)
	}

	if 0 == len(as) {
		return nil, errors.E_ARTICLE_NOT_FOUND
	}

	return as, nil
}

func (*articleModel) InfoById(id string) (*Article, error) {
	sql := fmt.Sprintf("SELECT id, title,content,author,kind,ctime FROM "+TBL_ARTICLES+" WHERE id='%s'", id)
	log.Println("sql:", sql)
	row := DB.QueryRow(sql)
	var a Article
	err := row.Scan(&a.Id, &a.Title, &a.Content, &a.Author, &a.Kind, &a.Ctime)
	if nil != err {
		log.Println(err)
		return nil, err
	}

	return &a, nil
}

func (*articleModel) InfoByKind(kind string) ([]*Article, error) {
	sql := fmt.Sprintf("SELECT id, title,content,author,kind,ctime FROM "+TBL_ARTICLES+" WHERE kind='%s'", kind)
	log.Println("sql:", sql)
	rows, err := DB.Query(sql)
	if nil != err {
		log.Println(err)
		return nil, err
	}

	var as = []*Article{}
	for rows.Next() {
		var a Article
		err := rows.Scan(&a.Id, &a.Title, &a.Content, &a.Author, &a.Kind, &a.Ctime)
		if nil != err {
			log.Println(err)
			continue
		}
		r := []rune(a.Content)
		length := len(r)
		if 800 >= length {
			a.Abs = a.Content
		} else {
			a.Abs = string(r[:800]) + "..."
		}

		as = append(as, &a)
	}

	if 0 == len(as) {
		return nil, errors.E_ARTICLE_NOT_FOUND
	}

	return as, nil
}

func (*articleModel) InfoByAuthor(author string) ([]*Article, error) {
	sql := fmt.Sprintf("SELECT id, title,content,author,kind,ctime FROM "+TBL_ARTICLES+" WHERE author='%s'", author)
	log.Println("sql:", sql)
	rows, err := DB.Query(sql)
	if nil != err {
		log.Println(err)
		return nil, err
	}

	var as = []*Article{}
	for rows.Next() {
		var a Article
		err := rows.Scan(&a.Id, &a.Title, &a.Content, &a.Author, &a.Kind, &a.Ctime)
		if nil != err {
			log.Println(err)
			continue
		}
		r := []rune(a.Content)
		length := len(r)
		if 800 >= length {
			a.Abs = a.Content
		} else {
			a.Abs = string(r[:800]) + "..."
		}

		as = append(as, &a)
	}

	if 0 == len(as) {
		return nil, errors.E_ARTICLE_NOT_FOUND
	}

	return as, nil
}
