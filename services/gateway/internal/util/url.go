package util

import "fmt"

func BuildUrl(host string, path string, query string) string {
  url := fmt.Sprintf("%s%s", host, path)

  if query != "" {
    url = fmt.Sprintf("%s?%s", url, query)
  }

  return url
}
