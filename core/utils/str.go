package utils

import (
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode"
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
	var seededRand = rand.New(rand.NewSource(time.Now().UnixNano() + int64(rand.Intn(100))))

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

func (st *str) EllipticalTruncate(text string, maxLen int) string {
	lastSpaceIx := maxLen
	len := 0
	for i, r := range text {
		if unicode.IsSpace(r) {
			lastSpaceIx = i
		}
		len++
		if len > maxLen {
			return text[:lastSpaceIx] + "..."
		}
	}
	// If here, string is shorter or equal to maxLen
	return text
}
