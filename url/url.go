package url

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

// QueryMap is key value of url query string
type QueryMap map[string]interface{}

// UserInfo is a struct of url username and password
type UserInfo struct {
	Username string
	Password string
}

// String converts UserInfo object to string
func (u UserInfo) String() (str string) {
	if u.Username != "" {
		str = u.Username
		if u.Password != "" {
			str += ":" + u.Password + "@"
		}
	}

	return str
}

// URL is a struct represent segments of a url
type URL struct {
	Scheme    string
	UserInfo  *UserInfo
	Host      string
	Domain    string
	Subdomain string
	Port      string
	Path      string
	Query     string
	Fragment  string
}

// URI replaces current path and query string with given path and query data
func (u *URL) URI(path string, params QueryMap) *URL {
	u.ReplacePath(path)
	u.parseParams(params)

	return u
}

// URIAppend Appends given path to current path and replace given query data
func (u *URL) URIAppend(path string, params QueryMap) *URL {
	u.AppendPath(path)
	u.parseParams(params)

	return u
}

// AppendPath appends given path to current path with no changes to query string
func (u *URL) AppendPath(path string) *URL {
	u.Path = fmt.Sprintf(`%s/%s`, strings.TrimRight(u.Path, "/"), strings.TrimLeft(path, "/"))

	return u
}

// ReplacePath replaces given path with current path with no changes to query string
func (u *URL) ReplacePath(path string) *URL {
	u.Path = strings.TrimLeft(path, "/")

	return u
}

// GetHost returns Host part of URL containing subdomain and domain
func (u *URL) GetHost() string {
	return u.Host
}

// GetDomain returns only Domain part of URL without any subdomain or www
func (u *URL) GetDomain() string {
	return u.Domain
}

// GetPath returns only the Path part of url object
func (u *URL) GetPath() string {
	return u.Path
}

// GetBaseDomain return base domain which is Domain prefixed with www as subdomain
func (u *URL) GetBaseDomain() string {
	return fmt.Sprintf("www.%s", u.Domain)
}

// GetURLString convert url object into url string
// [scheme://][userInfo@]host[:port][/path][?query][#fragment]
func (u *URL) GetURLString() string {
	var scheme, port, path, query, fragment string

	if u.Scheme != "" {
		scheme = fmt.Sprintf("%s://", u.Scheme)
	}

	if u.Port != "" {
		port = fmt.Sprintf(":%s", u.Port)
	}

	if u.Path != "" {
		path = fmt.Sprintf("/%s", u.Path)
	}

	if u.Query != "" {
		query = fmt.Sprintf("?%s", u.Query)
	}

	if u.Fragment != "" {
		fragment = fmt.Sprintf("#%s", u.Fragment)
	}

	return scheme + u.UserInfo.String() + u.GetHost() + port + path + query + fragment
}

// String is an alias for GetURLString
func (u *URL) String() string {
	return u.GetURLString()
}

// parse QueryMap into query string value
func (u *URL) parseParams(params QueryMap) {
	var queryString string
	for key, value := range params {
		if v, ok := value.([]interface{}); ok {
			for _, val := range v {
				queryString += fmt.Sprintf("&%s[]=%v", key, val)
			}
		} else {
			queryString += fmt.Sprintf("&%s=%v", key, value)
		}
	}

	u.Query = strings.TrimLeft(queryString, "&")
}

// Parse the given url into URL object
func Parse(url string) (*URL, error) {
	// URL regex `^[scheme]?[user_info]?[host][port]?[path]?[query]?#?[fragment]?`
	regexRule := fmt.Sprintf(`(?i)^%s?%s?%s`, getSchemeRegexRule(), getMainRegexRule(), getURLTrailRegexRule())
	regex, _ := regexp.Compile(regexRule)

	matches := regex.FindStringSubmatch(url)

	// If given url matches the regex rule, result will have all 11 group
	if len(matches) != 11 {
		return nil, errors.New(fmt.Sprintf("%s is an invalid url!", url))
	}

	urlObject := &URL{
		Scheme:    matches[1],
		UserInfo:  &UserInfo{Username: matches[2], Password: matches[3]},
		Host:      matches[4],
		Domain:    matches[6],
		Subdomain: matches[5],
		Port:      matches[7],
		Path:      matches[8],
		Query:     matches[9],
		Fragment:  matches[10],
	}

	return urlObject, nil
}

// NewUserInfo make new instance of UserInfo struct with given data
func NewUserInfo(username string, password string) *UserInfo {
	return &UserInfo{
		Username: username,
		Password: password,
	}
}

func getSchemeRegexRule() string {
	return `(?:([a-zA-z0-9]+):(?:\/\/)?)`
}

func getMainRegexRule() string {
	return fmt.Sprintf(`%s?%s%s`, getUserInfoRegexRule(), getHostRegexRule(), getPortRegexRule())
}

func getUserInfoRegexRule() string {
	return `(?:([a-zA-z0-9\-]+):([a-zA-z0-9\-]*)?@)`
}

func getHostRegexRule() string {
	return `((?:([a-zA-z0-9\-\.]+)\.)?([a-zA-z0-9\-]+\.[a-zA-z]{2,}))`
}

func getPortRegexRule() string {
	return `(?::(\d+))`
}

func getURLTrailRegexRule() string {
	return fmt.Sprintf(`%s?%s?#?%s?`, getPathRegexRule(), getQueryRegexRule(), getFragmentRegexRule())
}

func getPathRegexRule() string {
	return `(?:\/([\+~%\/\.\w\-_]*)`
}

func getQueryRegexRule() string {
	return `(?:\?([\-\+=&;%@\.\w_]*)?)`
}

func getFragmentRegexRule() string {
	return `([\.\!\/\\\w\-_%]*)?)`
}
