package flydown

import (
	"net/http"
	"os"
	"strings"
)

type clientCookies struct {
	lastPage string
	theme    string
}

func saveLastPageInCookies(lastPage string, w http.ResponseWriter) {
	if !strings.Contains(lastPage, "search?search_string") {
		if searchStrIndex := strings.Index(lastPage, "?search_string"); searchStrIndex > 0 {
			lastPage = lastPage[:searchStrIndex]
		}
		cookie := http.Cookie{Name: "last-page", Value: lastPage}
		http.SetCookie(w, &cookie)
	}
}

func isLastPageValid(r *http.Request) bool {
	userCookie, err := r.Cookie("last-page")
	if err == nil {
		// last page validation
		fullAddr := MdGenerator.rootMdFolder + string(os.PathSeparator) + userCookie.Value

		if _, err := os.Stat(fullAddr); !os.IsNotExist(err) {
			return true
		}
	}
	return false
}

func getUserPreferences(r *http.Request) (c clientCookies) {

	userCookie, err := r.Cookie("last-page")

	if isLastPageValid(r) {
		c.lastPage = userCookie.Value
	} else {
		c.lastPage = MdGenerator.readmeName
	}

	userCookie, err = r.Cookie("theme")
	if err != nil {
		c.theme = defaultColorTheme
	} else {
		c.theme = userCookie.Value
	}
	return
}
