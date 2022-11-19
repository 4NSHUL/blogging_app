package main

import (
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

const region = "Asia/Kolkata"

// TableName overrides the table name used by User to `profiles`
func (Blog) TableName() string {
	return "tbl_blog"
}

func FetchBlogsData() (blogs []Blog, err error) {
	Instance.Raw(`SELECT * From tbl_blog`).Scan(&blogs)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(blogs)
	return
}

func FetchBlogData(id int) (blog Blog, err error) {
	Instance.Raw(`SELECT * From tbl_blog where id = ?`, id).Scan(&blog)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(blog)
	return
}

func CreateBlogData(blogReqBody Blog) error {
	//TODO further validations can be added and appropriate error can be returned.
	blogReqBody.CreatedAt, _ = time.Parse("2006-01-02 15:04", time.Now().Format("2006-01-02 15:04"))
	result := Instance.Create(&blogReqBody).Omit("id")
	if result.Error != nil {
		fmt.Println(result.Error)
	}

	return result.Error
}

func DeleteBlogData(id int) error {
	result := Instance.Where("id = ?", id).Delete(&Blog{})
	if result.Error != nil {
		fmt.Println(result.Error)
	} else if result.RowsAffected == 0 {
		return gorm.ErrInvalidValue
	}

	return result.Error
}

func UpdateBlogData(id int, blogUpdateReqBody BlogUpdate) error {
	result := Instance.Model(&Blog{}).Where("id = ?", id).Updates(map[string]interface{}{"title": blogUpdateReqBody.Title, "content": blogUpdateReqBody.Content})
	if result.Error != nil {
		fmt.Println(result.Error)
	} else if result.RowsAffected == 0 {
		return gorm.ErrInvalidValue
	}

	return result.Error
}

func GeneratehashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
