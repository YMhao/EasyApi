package serv

import "strings"

func formatDescript(description string) string {
	description = strings.Replace(description, "\n", " ", -1)
	description = strings.Replace(description, "\t", " ", -1)
	return description
}
