package api

import (
  "net/http"
  "io/ioutil"
)

var IndexHtml string

//  Serve assets/index.html. This means that the frontend will check the route
//  and send its 404-Error if necessary.
func Handle404(writer http.ResponseWriter, request *http.Request) {
    index, err := ioutil.ReadFile(IndexHtml)
    if err != nil {
        writer.Write([]byte("Internal Server error! Can't provide index.html."))
    }
    writer.Header().Set("Content-Type", "text/html")
    writer.Write(index)
}
