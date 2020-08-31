package m2j

import (
	"fmt"
	"regexp"
	"strings"
)

type jiration struct {
	re   *regexp.Regexp
	repl interface{}
}

// MDToJira takes a string in Github Markdown, and outputs Jira text styling
func MDToJira(str string) string {
	jirations := []jiration{
		{ // bold and italics
			re: regexp.MustCompile(`([*_]{1,2})(\S.*?)([*_]{1,2})`),
			repl: func(groups []string) string {
				match, start, content, end := groups[0], groups[1], groups[2], groups[3]

				switch {
				// Without backreferences in RE2, this is the best we can do
				case start != end:
					return match
				case len(start) == 1:
					return fmt.Sprintf("_%s_", content)
				default:
					return fmt.Sprintf("*%s*", content)
				}
			},
		},
		{ // code blocks
			//re: regexp.MustCompile("(?s)```(.+\n)?((?:.|\n)*?)```"),
			re: regexp.MustCompile("(?s)```(.+?\n)?(.*?)```"),
			repl: func(groups []string) string {
				// If there is syntax, it will be in groups[1], and content in groups[2].
				// Otherwise, content is in groups[1] and groups[2] is empty.

				if groups[2] == "" {
					return "{code}" + groups[1] + "{code}"
				} else {
					return fmt.Sprintf("{code:%s}\n", strings.TrimSpace(groups[1])) + groups[2] + "{code}"
				}
			},
		},
		{ // monospace
			re:   regexp.MustCompile("`([^`]+)`"),
			repl: "{{$1}}",
		},
		{ // headings
			re: regexp.MustCompile(`(?m)^(#{1,6})\s`),
			repl: func(groups []string) string {
				return fmt.Sprintf("h%d. ", len(groups[1]))
			},
		},
		{ // unnamed links
			re:   regexp.MustCompile(`<([^>]+?)>`),
			repl: "[$1]",
		},
		{ // named links
			re:   regexp.MustCompile(`\[([^\]]+)\]\(([^)]+)\)`),
			repl: "[$1|$2]",
		},
		{ // unordered lists
			re: regexp.MustCompile(`(?m)^([ \t]*)\*\s+`),
			repl: func(groups []string) string {
				indent := 0
				indentStr := groups[1]
				if indentStr != "" {
					if string(indentStr[0]) == " " {
						indent = len(indentStr) / 2
					} else {
						indent = len(indentStr)
					}
				}
				return strings.Repeat("*", indent+1) + " "
			},
		},
	}
	for _, jiration := range jirations {
		switch v := jiration.repl.(type) {
		case string:
			str = jiration.re.ReplaceAllString(str, v)
		case func([]string) string:
			str = replaceAllStringSubmatchFunc(jiration.re, str, v)
		default:
			fmt.Printf("I don't know about type %T!\n", v)
		}
	}
	return str
}

// https://gist.github.com/elliotchance/d419395aa776d632d897
func replaceAllStringSubmatchFunc(re *regexp.Regexp, str string, repl func([]string) string) string {
	result := ""
	lastIndex := 0

	for _, v := range re.FindAllSubmatchIndex([]byte(str), -1) {
		groups := []string{}
		for i := 0; i < len(v); i += 2 {
			groups = append(groups, str[v[i]:v[i+1]])
		}

		result += str[lastIndex:v[0]] + repl(groups)
		lastIndex = v[1]
	}

	return result + str[lastIndex:]
}
