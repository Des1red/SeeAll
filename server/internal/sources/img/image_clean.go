package img

import "strings"

func CleanImageURL(image string) string {

	if image == "" {
		return ""
	}

	image = strings.TrimSpace(image)

	// decode HTML entities
	image = strings.ReplaceAll(image, "&amp;", "&")

	// protocol-less
	if strings.HasPrefix(image, "//") {
		image = "https:" + image
	}

	// require absolute URL
	if !strings.HasPrefix(image, "http://") && !strings.HasPrefix(image, "https://") {
		return ""
	}

	return image
}
