package test

import (
	"testing"

	"github.com/shinshin8/myFavorite/utils"
)

func TestIsName(t *testing.T) {
	var testName string
	testName = "testName"
	res := utils.IsName(testName)
	if !res {
		t.Fatal("failed test")
	}
}

func TestIsEmail(t *testing.T) {
	var emailAddress string
	emailAddress = "test@test.com"
	res := utils.IsEmailAddress(emailAddress)
	if !res {
		t.Fatal("failed test")
	}
}

func TestIsPassword(t *testing.T) {
	var password string
	password = "hogehogehoge"
	res := utils.IsPassword(password)
	if !res {
		t.Fatal("failed test")
	}
}

func TestIsID(t *testing.T) {
	var id int
	id = 1
	res := utils.IsID(id)
	if !res {
		t.Fatal("failed test")
	}
}

func TestIsTitle(t *testing.T) {
	var title string
	title = "test title"
	res := utils.IsTitle(title)
	if !res {
		t.Fatal("failed test")
	}
}

func TestIsContent(t *testing.T) {
	var content string
	content = "this is content example"
	res := utils.IsContent(content)
	if !res {
		t.Fatal("failed test")
	}
}

func TestIsComment(t *testing.T) {
	var comment string
	comment = "this is test comment"
	res := utils.IsComment(comment)
	if !res {
		t.Fatal("failed test")
	}
}

func TestIsBirthday(t *testing.T) {
	var birthday string
	birthday = "20001212"
	res := utils.IsBirthday(birthday)
	if !res {
		t.Fatal("failed test")
	}
}
