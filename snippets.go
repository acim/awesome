import "net/http"

func WithDB(s *mgo.Session, h http.Handler) http.Handler {
	return &dbWrapper{dbSession: s, h: h}
}

type dbWrapper struct {
	dbSession *mgo.Session
	h         http.Handler
}

func (dbWrapper *dbWrapper) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	dbCopy := h.dbSession.Copy()
	defer dbCopy.Close()

	context.Set(r, "db", dbCopy) // bad idea to put db in context, but shows up how to put something reasonable
	dbWrapper.h.ServeHTTP(w, r)
}
