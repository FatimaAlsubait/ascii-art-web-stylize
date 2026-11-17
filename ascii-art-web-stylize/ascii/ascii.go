package ascii

import (
	"errors"
	"os"
	"strings"
)

func Generate(text, banner string) (string, error) {
	if text == "" {
		return "", nil
	}

	bannerLines, err := loadBanner(banner)
	if err != nil {
		return "", err
	}

	cleanText := text

	// 1. First replace typed literal backslash sequences ("\r\n", "\r", "\n")
	cleanText = strings.ReplaceAll(cleanText, "\\r\\n", "\n")
	cleanText = strings.ReplaceAll(cleanText, "\\r", "\n")
	cleanText = strings.ReplaceAll(cleanText, "\\n", "\n")

	// 2. Replace real browser/system carriage returns with actual newlines
	cleanText = strings.ReplaceAll(cleanText, "\r\n", "\n")
	cleanText = strings.ReplaceAll(cleanText, "\r", "\n")

	inputLines := strings.Split(cleanText, "\n")

	var result strings.Builder

	for _, line := range inputLines {
		if line == "" {
			result.WriteString("\n")
			continue
		}

		lineOutput, err := buildAsciiLine(line, bannerLines)
		if err != nil {
			return "", err
		}
		result.WriteString(lineOutput)
	}

	return result.String(), nil
}

func loadBanner(banner string) ([]string, error) {
	bannerPath := "banner/" + banner + ".txt"
	content, err := os.ReadFile(bannerPath)
	if err != nil {
		return nil, errors.New("banner file not found")
	}

	normalizedContent := strings.ReplaceAll(string(content), "\r\n", "\n")
	normalizedContent = strings.ReplaceAll(normalizedContent, "\r", "\n")
	return strings.Split(normalizedContent, "\n"), nil
}

func buildAsciiLine(line string, bannerLines []string) (string, error) {
	var lineBuilder strings.Builder

	for row := 0; row < 8; row++ {
		for _, char := range line {
			if char < 32 || char > 126 {
				return "", errors.New("invalid character")
			}

			lineIndex := int(char-32)*9 + 1 + row
			if lineIndex < len(bannerLines) {
				lineBuilder.WriteString(bannerLines[lineIndex])
			}
		}
		lineBuilder.WriteString("\n")
	}

	return lineBuilder.String(), nil
}
