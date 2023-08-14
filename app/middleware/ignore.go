package middleware

import "strings"

// publicPaths is a map of public paths
var publicPaths = map[string]bool{
	"/":     true,
	"/ping": true,
}

// devPaths is a map of dev paths
var devPaths = map[string]bool{}

func isPublicPath(path string) bool {
	return publicPaths[path]
}

func isDevPath(path string) bool {
	if isPublicPath(path) {
		return false
	}
	if strings.HasPrefix(path, "/v1/dev") {
		return true
	}

	return devPaths[path]
}
