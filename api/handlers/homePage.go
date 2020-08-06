package handlers

import (
	"fmt"
	"net/http"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	var innerHtml = `
	<html>
		<body>
			<a href="/auth/google/login" >Login with Google</a>
		</body>
	</html>
	`
	fmt.Fprintf(w, innerHtml)
}
