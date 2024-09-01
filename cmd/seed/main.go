package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/wkdwilliams/go-blog/internal/adapters/secondary/mysql"
	"github.com/wkdwilliams/go-blog/internal/domain/services"
	"github.com/wkdwilliams/go-blog/internal/infrastructure"
)

func main() {
	// Load env
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	db, err := infrastructure.NewMysql()
	if err != nil {
		log.Fatal(err)
	}

	userRepo := mysql.NewUserRepository(db)
	usersService := services.NewUserService(userRepo)

	postRepo := mysql.NewPostRepository(db)
	postService := services.NewPostService(postRepo)

	user, err := usersService.CreateAccount("admin", "pass", "lewis")
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range posts {
		_, err = postService.Create(v[0], v[1], user.ID)
		if err != nil {
			log.Fatal(err)
		}
	}
}

var posts [][]string = [][]string{
	{
		"Decoding a struct in Golang",
		`<p>Decoding structs to a JSON string is easy. We can do this using go's built-in encoding package. Here is an example of creating a small HTTP server and writing the json to the response writer</p>
        <pre><code class="language-golang">package main

import (
	"encoding/json"
	"net/http"
	"log"
)

// Define a simple struct
type User struct {
	ID    int    ` + "`json:\"id\"`" + `
	Name  string ` + "`json:\"name\"`" + `
	Email string ` + "`json:\"email\"`" + `
}

func main() {
	http.HandleFunc("/user", userHandler)
	log.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Handler function to serve the User struct as JSON
func userHandler(w http.ResponseWriter, r *http.Request) {
	// Create an instance of User
	user := User{
		ID:    1,
		Name:  "John Doe",
		Email: "john.doe@example.com",
	}

	// Set the Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Encode the User struct to JSON and write it to the response
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
		return
	}
}
</code></pre>`,
	},
	{
		"Understanding Async/Await in JavaScript",
		`<p>Async/await is a powerful feature in JavaScript that simplifies working with asynchronous code. Here's how it works:</p>
        <pre><code class="language-javascript">async function fetchData() {
    try {
        const response = await fetch('https://api.example.com/data');
        const data = await response.json();
        console.log(data);
    } catch (error) {
        console.error('Error fetching data:', error);
    }
}
</code></pre>`,
	},
}
