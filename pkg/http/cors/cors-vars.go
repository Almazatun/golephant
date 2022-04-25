package cors

var (
	CORS_ALLOWED_HEADERS = []string{"*", "X-Requested-With", "Content-Type", "Authorization"}
	CORS_ALLOWED_METHODS = []string{"GET", "POST", "PUT", "DELETE"}
	CORS_ALLOWED_ORIGINS = []string{"*"}
)
