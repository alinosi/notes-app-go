package main

imprt (
		"encoding/json"
		"net/http"
)

// users dummy data
type Users struct {
	User_id    string `json:"id"`
	Email      string `json:"email"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Role 	   string `json:"role"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
	Deleted_at string `json:"deleted_at"`
}


// notes dummy data
type Notes struct {
	Note_Id    string `json:"id"`
	User_id    string `json:"email"`
	Text	   string `json:"text"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
	Deleted_at string `json:"deleted_at"`
}

func main() {

		// add the server mux or routes controller
		http.ListenAndServe(":8080", nil)
	
		// routes controller
		// http.handlerfunc("home", func () => ) {

		// }
}
