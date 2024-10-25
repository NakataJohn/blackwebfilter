package parseutil

import (
	"fmt"
	"net/url"
	"strings"
)

// getDomain从完整的URL中提取第一个点之后的域名部分
func GetDomain(urlStr string) (string, error) {
	u, err := url.Parse(urlStr)
	if err != nil {
		return "", err
	}
	domain := u.Hostname()

	// 查找第一个点的位置
	idx := strings.Index(domain, ".")
	if idx == -1 {
		return "", fmt.Errorf("no dot found in domain: %s", domain)
	}

	// 返回第一个点之后的域名部分，包括点
	return "." + domain[idx+1:], nil
}

// 返回所有子域名
func GetSubDomains(urlStr string) []string {
	u, err := url.Parse(urlStr)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return nil
	}
	host := strings.Trim(u.Hostname(), ".")

	// Split the host by dots and reverse to get the subdomains
	parts := strings.Split(host, ".")

	// Create the subdomain list
	subdomains := make([]string, 0, len(parts))
	for i := 0; i < len(parts); i++ {
		subdomains = append(subdomains, "."+strings.Join(parts[i:], "."))
	}
	return subdomains
}
