package handle

import (
	"net/http"
)

// 结构体
type pathURL struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}

func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path, ok := pathsToUrls[r.URL.Path]
		if ok {
			http.Redirect(w, r, path, http.StatusFound)
		} else {
			fallback.ServeHTTP(w, r)
		}
	}
}

/*
func YAMLHandler(yamlbuyes []byte, fallback http.Handler) (http.HandlerFunc, error) {
	pathURLs, errr := ParseYaml(yamlbuyes)
	if errr != nil {
		log.Fatal(errr)
		return nil, errr
	}
	pathsToUrls := buildMap(pathURLs)
	return MapHandler(pathsToUrls, fallback), nil
}

// 解析Yaml,把数据编程Yaml对应的那个结构体
func ParseYaml(data []byte) ([]pathURL, error) {
	var pathURLs []pathURL
	err := yaml.Unmarshal(data, &pathURLs)
	if err != nil {
		log.Fatal("yaml parse error", zap.Error(err))
		return nil, err
	}
	return pathURLs, nil
}

// 把所有的一堆ParseYaml得到的结合在一起
func buildMap(pathURLs []pathURL) map[string]string {
	pathToUrl := make(map[string]string)
	for _, pathURL := range pathURLs {
		pathToUrl[pathURL.Path] = pathURL.URL
	}
	return pathToUrl
}
*/
