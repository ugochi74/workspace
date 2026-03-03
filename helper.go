package main

import (
    "strconv"
    "strings"
    "unicode"
)

func hexToDec(word string) string {
    val, err := strconv.ParseInt(word, 16, 64)
    if err != nil {
        return word
    }
    return strconv.FormatInt(val, 10)
}

func binToDec(word string) string {
    val, err := strconv.ParseInt(word, 2, 64)
    if err != nil {
        return word
    }
    return strconv.FormatInt(val, 10)
}

func applyCase(result []string, mode string, command string) []string {
    count := 1
    if strings.Contains(command, ",") {
        parts := strings.Split(command, ",")
        countStr := strings.Trim(parts[1], ")")
        c, err := strconv.Atoi(countStr)
        if err == nil {
            count = c
        }
    }

    for i := len(result) - count; i < len(result); i++ {
        if i >= 0 {
            switch mode {
            case "up":
                result[i] = strings.ToUpper(result[i])
            case "low":
                result[i] = strings.ToLower(result[i])
            case "cap":
                runes := []rune(result[i])
                if len(runes) > 0 {
                    runes[0] = unicode.ToUpper(runes[0])
                    for j := 1; j < len(runes); j++ {
                        runes[j] = unicode.ToLower(runes[j])
                    }
                    result[i] = string(runes)
                }
            }
        }
    }
    return result
}
