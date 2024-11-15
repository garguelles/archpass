package utils

import (
	"regexp"
	"strings"
)

// Slugify generates a URL-friendly slug from a given string.
func Slugify(name string) string {
	// Convert to lowercase
	name = strings.ToLower(name)

	// Replace spaces and underscores with hyphens
	name = strings.ReplaceAll(name, " ", "-")
	name = strings.ReplaceAll(name, "_", "-")

	// Remove special characters
	re := regexp.MustCompile(`[^a-z0-9-]+`)
	name = re.ReplaceAllString(name, "")

	// Trim any leading or trailing hyphens
	name = strings.Trim(name, "-")

	return name
}
