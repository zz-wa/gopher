package handle

import (
	"database/sql"
	"net/http"

	"go.uber.org/zap"
)

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

func GetURL(db *sql.DB, number int) (map[string]string, error) {
	rows, err := db.Query("SELECT originalURL, shorterURL   FROM urlShorter WHERE number= ?", number)
	if err != nil {
		zap.L().Error("error getting url", zap.Error(err))
		return nil, err
	}
	defer rows.Close()

	urls := make(map[string]string)

	for rows.Next() {
		var shorterURL, originalURL string
		err := rows.Scan(&shorterURL, &originalURL)
		if err != nil {
			zap.L().Error("error getting url", zap.Error(err))
		}
		urls[originalURL] = shorterURL
	}
	return urls, nil
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
