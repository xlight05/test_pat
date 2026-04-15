package handlers

import (
	"encoding/json"
	"html"
	"html/template"
	"net/http"
	"strings"
)

var homeTmpl = template.Must(template.New("home").Parse(`<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<meta name="viewport" content="width=device-width, initial-scale=1.0">
<title>TestFlow</title>
<style>
*,*::before,*::after{box-sizing:border-box}
body{
  margin:0;padding:2rem 1rem;
  font-family:-apple-system,BlinkMacSystemFont,"Segoe UI",Roboto,Helvetica,Arial,sans-serif;
  font-size:1.2rem;line-height:1.6;
  color:#333;background:#fff;
}
main{max-width:600px;margin:0 auto;text-align:center}
h1{font-size:2rem;margin-bottom:1.5rem}
form{display:flex;flex-direction:column;align-items:center;gap:0.75rem}
label{font-weight:600}
input[type="text"]{
  padding:0.5rem 0.75rem;font-size:1rem;
  border:2px solid #555;border-radius:4px;
  width:100%;max-width:300px;
  color:#333;background:#fff;
}
input[type="text"]:focus{outline:2px solid #0066cc;outline-offset:2px;border-color:#0066cc}
button{
  padding:0.6rem 1.5rem;font-size:1rem;
  background:#0066cc;color:#fff;
  border:none;border-radius:4px;cursor:pointer;
}
button:hover{background:#0052a3}
button:focus{outline:2px solid #0066cc;outline-offset:2px}
</style>
</head>
<body>
<main>
  <h1>{{.Greeting}}</h1>
  <form method="GET" action="/">
    <label for="name-input">Enter your name</label>
    <input type="text" id="name-input" name="name" placeholder="Your name" value="{{.Name}}">
    <button type="submit">Greet</button>
  </form>
</main>
</body>
</html>`))

var notFoundTmpl = template.Must(template.New("404").Parse(`<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<meta name="viewport" content="width=device-width, initial-scale=1.0">
<title>TestFlow - Not Found</title>
<style>
*,*::before,*::after{box-sizing:border-box}
body{
  margin:0;padding:2rem 1rem;
  font-family:-apple-system,BlinkMacSystemFont,"Segoe UI",Roboto,Helvetica,Arial,sans-serif;
  font-size:1.2rem;line-height:1.6;
  color:#333;background:#fff;
}
main{max-width:600px;margin:0 auto;text-align:center}
h1{font-size:2rem;color:#cc0000}
a{color:#0066cc;text-decoration:underline}
a:focus{outline:2px solid #0066cc;outline-offset:2px}
</style>
</head>
<body>
<main>
  <h1>404 — Page Not Found</h1>
  <p>The page you requested does not exist.</p>
  <p><a href="/">Go back home</a></p>
</main>
</body>
</html>`))

type homeData struct {
	Greeting string
	Name     string
}

// Home handles GET / with optional ?name= query parameter.
func Home(w http.ResponseWriter, r *http.Request) {
	// Go 1.22 ServeMux: "/" is a catch-all, so reject non-root paths.
	if r.URL.Path != "/" {
		notFound(w, r)
		return
	}

	name := strings.TrimSpace(r.URL.Query().Get("name"))
	name = html.EscapeString(name)

	greeting := "Hello World"
	if name != "" {
		greeting = "Hello, " + name + "!"
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	homeTmpl.Execute(w, homeData{Greeting: greeting, Name: name})
}

// Health returns {"status":"ok"} as JSON.
func Health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusNotFound)
	notFoundTmpl.Execute(w, nil)
}
