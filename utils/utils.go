package utils

import (
	"context"
	"douyin/pkg/constants"
	"douyin/pkg/errno"
	"douyin/pkg/minio"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"mime/multipart"
	"regexp"
	"strings"
)

// EncryptPassword 对密码进行加密
func EncryptPassword(password string) (string, error) {
	cost := 5

	hashed, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	return string(hashed), err
}

// VerifyPassword 验证密码
func VerifyPassword(pass, hashed string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(pass))
	if err != nil {
		return false
	}
	return true
}

// URLConvert Convert the path in the database into a complete url accessible by the front end
func URLConvert(ctx context.Context, c *app.RequestContext, path string) (fullURL string) {
	if len(path) == 0 {
		return ""
	}
	arr := strings.Split(path, "/")
	u, err := minio.GetObjectURL(ctx, arr[0], arr[1])
	if err != nil {
		hlog.CtxInfof(ctx, err.Error())
		return ""
	}
	u.Scheme = string(c.URI().Scheme())
	u.Host = string(c.URI().Host())
	u.Path = "/src" + u.Path
	return u.String()
}

// ReadFileBytes 读取文件字节数据
func ReadFileBytes(file *multipart.FileHeader) ([]byte, error) {
	s, err := file.Open()
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(s)
}

func CheckPassword(password string) (bool, *errno.ErrNo) {
	pats := []string{"[a-zA-Z]", "[0-9]", "[^\\d\\w]"}
	switch {
	case len(password) < constants.PasswordMinLen:
		return false, &errno.ErrPassWordBelowSize
	case len(password) > constants.PassWordMaxLen:
		return false, &errno.ErrPassWordOverSize
	}
	for _, pat := range pats {
		if ok, _ := regexp.MatchString(pat, password); !ok {
			return false, &errno.ErrPassWordSymbols
		}
	}
	return true, nil
}

type TrieNode struct {
	children map[rune]*TrieNode
	isEnd    bool
}

func NewTrieNode() *TrieNode {
	return &TrieNode{
		children: make(map[rune]*TrieNode),
		isEnd:    false,
	}
}

func InsertWord(root *TrieNode, word string) {
	node := root
	for _, ch := range word {
		if _, ok := node.children[ch]; !ok {
			node.children[ch] = NewTrieNode()
		}
		node = node.children[ch]
	}
	node.isEnd = true
}

func SearchText(root *TrieNode, text string) bool {
	node := root
	for _, ch := range text {
		if _, ok := node.children[ch]; !ok {
			return false
		}
		node = node.children[ch]
		if node.isEnd {
			return true
		}
	}
	return false
}

func SensitiveWordDetection(sensitiveWords []string, text string) bool {
	root := NewTrieNode()
	for _, word := range sensitiveWords {
		InsertWord(root, word)
	}
	return SearchText(root, text)
}
