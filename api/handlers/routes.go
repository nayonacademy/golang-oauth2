package handlers

func (s *Server) initializeRoutes() {
	// Home Route
	//s.Router.Handle("/",  http.FileServer(http.Dir("api/templates/")))
	s.Router.HandleFunc("/",s.Home).Methods("GET")
	// Login Route
	s.Router.HandleFunc("/auth/google/login", s.handleGoogleLogin).Methods("GET")

	//Users routes
	s.Router.HandleFunc("/auth/google/callback", s.handleGoogleCallback).Methods("GET")

}
