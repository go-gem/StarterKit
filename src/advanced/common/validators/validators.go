package validators

import "regexp"

var (
	regexpEmail *regexp.Regexp
)

func init() {
	var err error
	if regexpEmail, err = regexp.Compile(`([\w\-])+@([\w\-])+\.([a-zA-Z]){2,}`); err != nil {
		panic(err)
	}
}

func IsEmail(email string) bool {
	return regexpEmail.Match([]byte(email))
}
