package dynamic

import (
	"controllers/auth"
	"controllers/testimonials"
	testimonialAdmin "controllers/testimonials"
	"flag"
	"fmt"
	"html/template"
	"net/http"
	"reflect"
	"strings"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/gzip"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"
)

var (
	listenAddr  = flag.String("port", ":8080", "http listen address")
	environment = flag.String("env", "", "current app environment")
)

func init() {
	flag.Parse()
	m := martini.Classic()
	m.Use(gzip.All())
	m.Use(martini.Static("public"))
	store := sessions.NewCookieStore([]byte("dynamic-fab"))
	m.Use(sessions.Sessions("my_session", store))
	m.Use(render.Renderer(render.Options{
		Directory:  "views",
		Layout:     "layout",
		Extensions: []string{".tmpl", ".html"},
		Funcs: []template.FuncMap{
			{
				"StringsEqual": func(a, b string) bool {
					if strings.EqualFold(a, b) {
						return true
					}
					return false
				},
				"ShortenString": func(s string, l int) string {
					if len(s) <= l {
						return s
					}
					return fmt.Sprintf("%s...", s[:l])
				},
				"IntegerGreater": func(x interface{}, y interface{}) bool {

					if x == nil || y == nil {
						return false
					}

					var xint int = 0
					var yint int = 0

					xtyp := reflect.TypeOf(x)
					switch xtyp.Kind() {
					case reflect.Int:
						xint = int(x.(int))
					case reflect.Int32:
						xint = int(x.(int32))
					case reflect.Int16:
						xint = int(x.(int16))
					case reflect.Int64:
						xint = int(x.(int64))
					}

					ytyp := reflect.TypeOf(y)
					switch ytyp.Kind() {
					case reflect.Int:
						yint = int(y.(int))
					case reflect.Int32:
						yint = int(y.(int32))
					case reflect.Int16:
						yint = int(y.(int16))
					case reflect.Int64:
						yint = int(y.(int64))
					}

					if xint <= yint {
						return false
					}

					return true
				},
			},
		},
		Delims:          render.Delims{"{{", "}}"},
		Charset:         "UTF-8",
		IndentJSON:      true,
		HTMLContentType: "text/html",
	}))

	// Backend tasks
	m.Group("/admin/testimonials", func(r martini.Router) {
		r.Get("", auth.Check, testimonialAdmin.Index)
		r.Get("/:id", auth.Check, testimonialAdmin.Edit)
		r.Post("/:id", auth.Check, testimonialAdmin.Save)
		r.Delete("/:id", auth.Check, testimonialAdmin.Delete)
	})
	m.Group("/api/testimonials", func(r martini.Router) {
		r.Get("", testimonials.All)
		r.Get("/:id", testimonials.Get)
	})

	// m.Get("/adduser", auth.AddUser)

	// Serve Frontend
	m.Get("/**", func(rw http.ResponseWriter, req *http.Request, r render.Render) {
		bag := make(map[string]interface{}, 0)
		bag["Host"] = req.URL.Host
		r.HTML(200, "index", bag)
	})

	http.Handle("/", m)
}
