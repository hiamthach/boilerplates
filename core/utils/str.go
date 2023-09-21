package utils

import (
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type str struct{}

var Str str

func (st *str) ToIntSlice(s string) (x []int) {
	slc := strings.Split(s, ",")
	for _, v := range slc {
		i, e := strconv.Atoi(v)
		if e != nil {
			continue
		}
		x = append(x, i)
	}
	return
}
func (st *str) IsEmail(e string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(e)
}
func (st *str) RandomCode(length int) string {
	const charset = "ABCDEFGHIJKLMNPQRSTUVWXYZ123456789"
	var seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}
func (st *str) Random(length int) string {
	const charset = "abcdefghijklmnpqrstuvwxyzABCDEFGHIJKLMNPQRSTUVWXYZ123456789"
	var seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}
func (st *str) ToInt(x string) int {
	if i, e := strconv.Atoi(x); e == nil {
		return i
	}
	return 0
}
func (st *str) IsGuess(mail string) bool {
	isGuest, err := regexp.MatchString("guest-([a-z0-9]+)@skydancer.com", mail)
	if err != nil || !isGuest {
		return false
	}
	return true
}

func (st *str) IsRangeString(e string) bool {
	emailRegex := regexp.MustCompile(`^(\(|\[)\d{0,},\d{0,}(\))$|^(\[|\[)\d{0,},\d{0,}(\])$`)
	is := emailRegex.MatchString(e)
	if is {
		return true
	}
	emailRegex = regexp.MustCompile(`^\d*-\d*$`)
	return emailRegex.MatchString(e)
}

func (st *str) SplitRangeString(e string) (int, int) {
	e = strings.Replace(e, "(", "", 1)
	e = strings.Replace(e, ")", "", 1)
	e = strings.Replace(e, "[", "", 1)
	e = strings.Replace(e, "]", "", 1)
	strs := strings.Split(e, "-")
	var num1 int = 0
	var num2 int = 0
	if len(strs) == 2 {
		num1, _ = strconv.Atoi(strs[0])
		num2, _ = strconv.Atoi(strs[1])
	}
	return num1, num2
}
func (st *str) SplitStrToInt64(e string, sep string) []int64 {
	strs := strings.Split(e, sep)
	result := make([]int64, 0)
	for _, v := range strs {
		num1, _ := strconv.Atoi(v)
		result = append(result, int64(num1))
	}
	return result
}
