package helpers

import "strings"

func TagConverter(tagsStr string) []string {
	var tagList []string

	str := strings.Replace(tagsStr, " ", "", 0)
	for _, v := range strings.Split(str, "\n") {
		if v != "" {
			tagList = append(tagList, v)
		}
	}

	return tagList
}
