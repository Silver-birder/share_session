package shareCookie

import (
        "context"
        "encoding/json"
        "html/template"
        "net/http"
        "time"

        "google.golang.org/appengine"
        "google.golang.org/appengine/datastore"
)
type ShareCookie struct {
        Key string `json:"key"`
        Value string `json:"value"`
        Domain  string `json:"domain"`
        Path string `json:"path"`
        Expire time.Time `json:"expire"`
}

func init() {
        http.HandleFunc("/api/1.0/get/cookie", getCookie)
        http.HandleFunc("/api/1.0/set/cookie", setCookie)
        http.HandleFunc("/api/1.0/delete/cookie", deleteCookie)
        http.HandleFunc("/maintenance", maintenance)
}

func getCookie(w http.ResponseWriter, r *http.Request) {
        // どのブラウザからでもアクセス可能にする
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
        c := appengine.NewContext(r)
        q := datastore.NewQuery("ShareCookie").Filter("Domain=", r.FormValue("Domain"))
        shareCookies := make([]ShareCookie, 0)
        if _, err := q.GetAll(c, &shareCookies); err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
        }
        res, err := json.Marshal(shareCookies)
        if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
        }
        w.Header().Set("Content-Type", "application/json")
        w.Write(res)
}

func setCookie(w http.ResponseWriter, r *http.Request) {
        c := appengine.NewContext(r)
        var s ShareCookie
        q := datastore.NewQuery("ShareCookie").Filter("Domain=", r.FormValue("Domain")).Filter("Key=", r.FormValue("Key"))
        iter := q.Run(c)
        key, err := iter.Next(&s)
        if key == nil {
                key = datastore.NewIncompleteKey(c, "ShareCookie", shareCookieKey(c))
                s.Domain = r.FormValue("Domain")
                s.Key = r.FormValue("Key")
                s.Path = r.FormValue("Path")
        }
        s.Value = r.FormValue("Value")
        format := "2006-01-02 15:04:05"
        s.Expire, err =  time.Parse(format, r.FormValue("Expire")+" 00:00:00")
        if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
        }
        datastore.Put(c, key, &s)
}

func deleteCookie(w http.ResponseWriter, r *http.Request) {
        c := appengine.NewContext(r)
        var s ShareCookie
        q := datastore.NewQuery("ShareCookie").Filter("Domain=", r.FormValue("Domain")).Filter("Key=", r.FormValue("Key"))
        iter := q.Run(c)
        key, _ := iter.Next(&s)
        _ = datastore.Delete(c, key)
}
func maintenance(w http.ResponseWriter, r *http.Request) {
        c := appengine.NewContext(r)
        q := datastore.NewQuery("ShareCookie").Ancestor(shareCookieKey(c))
        shareCookies := make([]ShareCookie, 0)
        if _, err := q.GetAll(c, &shareCookies); err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
        }
        if err := maintenanceTemplate.Execute(w, shareCookies); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
func shareCookieKey(c context.Context) *datastore.Key {
        return datastore.NewKey(c, "ShareCookie", "default_shareCookie", 0, nil)
}

var maintenanceTemplate = template.Must(template.ParseFiles("maintenance.html"))
