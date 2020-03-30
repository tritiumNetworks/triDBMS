package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"
)

// Redirect Root Page
func Redirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/docs", 308)
}

// Route API
func Route(w http.ResponseWriter, r *http.Request) {
	apiV := strings.Split(r.URL.Path, "/api")[1]
	if apiV == "/" {
		apiV = "/v1"
	}

	switch apiV {
	case "/v1", "/v1.0", "/v1.0.0":
		qTarget := r.URL.Query().Get("target")
		qType := r.URL.Query().Get("type")

		if len(qTarget) < 1 {
			qTarget = "neko"
		}

		if len(qType) < 1 {
			qType = "page"
		}

		v1(qTarget, qType, w, r)
		break

	default:
		w.WriteHeader(404)
		w.Write([]byte("<html>404 - API Not Found<hr />triDBMS v2, written in go</html>"))
		break
	}
}

func v1(qTarget, qType string, w http.ResponseWriter, r *http.Request) {
	switch qType {
	case "url":
		rand.Seed(time.Now().UnixNano())
		arr := fileRead(qTarget)
		if len(arr) > 0 {
			w.Write([]byte("/static/" + qTarget + "/" + arr[rand.Intn(len(arr))]))
		} else {
			w.WriteHeader(404)
			w.Write([]byte("<html>404 - API Not Found<hr />triDBMS v2, written in go</html>"))
		}
		break

	case "page":
		rand.Seed(time.Now().UnixNano())
		arr := fileRead(qTarget)
		if len(arr) > 0 {
			w.Write([]byte("<html><img height=\"700\" src=\"/static/" + qTarget + "/" + arr[rand.Intn(len(arr))] + "\"></html>"))
		} else {
			w.WriteHeader(404)
			w.Write([]byte("<html>404 - API Not Found<hr />triDBMS v2, written in go</html>"))
		}
		break

	case "redirect":
		rand.Seed(time.Now().UnixNano())
		arr := fileRead(qTarget)
		if len(arr) > 0 {
			http.Redirect(w, r, "/static/"+qTarget+"/"+arr[rand.Intn(len(arr))], 302)
		} else {
			w.WriteHeader(404)
			w.Write([]byte("<html>404 - API Not Found<hr />triDBMS v2, written in go</html>"))
		}
		break

	case "buffer":
		rand.Seed(time.Now().UnixNano())
		arr := fileRead(qTarget)
		if len(arr) > 0 {
			pwd, err := os.Getwd()
			erring(err)

			data, err := ioutil.ReadFile(pwd + "/datas/" + qTarget + "/" + arr[rand.Intn(len(arr))])
			erring(err)

			w.Write(data)
		} else {
			w.WriteHeader(404)
			w.Write([]byte("<html>404 - API Not Found<hr />triDBMS v2, written in go</html>"))
		}
		break

	case "list":
		rand.Seed(time.Now().UnixNano())
		arr := fileRead(qTarget)
		if len(arr) > 0 {
			dataJSON, err := json.Marshal(arr)
			erring(err)
			w.Write([]byte(dataJSON))
		} else {
			w.WriteHeader(404)
			w.Write([]byte("<html>404 - API Not Found<hr />triDBMS v2, written in go</html>"))
		}
		break

	case "listImg":
		arr := fileRead(qTarget)
		if len(arr) > 0 {
			str := ""
			for _, img := range arr {
				str += "<img width=\"100\" src=\"/static/" + qTarget + "/" + img + "\">"
			}
			w.Write([]byte("<html>" + str + "</html>"))
		} else {
			w.WriteHeader(404)
			w.Write([]byte("<html>404 - API Not Found<hr />triDBMS v2, written in go</html>"))
		}
		break
	}
}

func fileRead(qTarget string) (imgs []string) {
	pwd, err := os.Getwd()
	erring(err)

	files, err := ioutil.ReadDir(pwd + "/datas")
	erring(err)

	for _, f := range files {
		if f.IsDir() && f.Name() == qTarget {
			i, err := ioutil.ReadDir(pwd + "/datas/" + qTarget)
			erring(err)

			for _, j := range i {
				imgs = append(imgs, j.Name())
			}
		}
	}

	return imgs
}

func erring(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
