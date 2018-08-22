package main

import (
   "net/http"
   "gopkg.in/mgo.v2"
)


func main() {
}

type Adapter func(http.Handler) http.Handler
func Adapt(h http.Handler, adapters ...Adapter) http.Handler {
   for _, adapter := range adapters {
      h = adapter(h)
   }
   return h
}

func withDB(db *mgo.Session) Adapter {
   // return the Adapter
   return func(h http.Handler) http.Handler {
      // the adapter (when called) should return a new handler
      return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
         // copy the database session
         dbsession := db.Copy()
         defer dbsession.Close() // clean up
         // save it in the mux context
         context.Set(r, "database", dbsession)
         // pass execution to the original handler
         h.ServeHTTP(w, r)
      })
   }
}