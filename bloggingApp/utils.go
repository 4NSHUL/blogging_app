package main

import "fmt"

func FetchBlogsData() (blogs []Blog, err error) {
	Instance.Raw(`SELECT * From blog`).Scan(&blogs)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(blogs)
	return
}
