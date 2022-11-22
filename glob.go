package glob

import "regexp"

// Compile returns a new, compiled regular expression from
// a traditional glob syntax. See regexp.Compile.
//
//     *       -> .*
//     ?       -> .
//     {3..22} -> (?:[3-9]|1[0-9]|2[0-2])
//     [abc]   -> [abc]
//     [0-9]   -> [0-9]
//
// Any character escaped with backslash (\) is treated normally.
func Compile(glob string) (*regexp.Regexp, error) {
	// TODO
	return nil, nil
}

// MustCompile calls Compile but panics instead of returning an error.
func MustCompile(glob string) *regexp.Regexp {
	re, err := Compile(glob)
	if err != nil {
		panic(err)
		return nil
	}
	return re
}
