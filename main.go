package main
import (
	"fmt"
	// "spotbuzz-backend/pkg/dbhelper"
	// "spotbuzz-backend/router"

	// "github.com/gin-gonic/gin"
	// _ "github.com/lib/pq"
	"net/http"
   	"os"
)

func main() {

	 port := os.Getenv("PORT")
	   if port == "" {
	      port = "8080"
	   }
	
	   http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	      fmt.Fprintln(w, "Hello, World!")
	   })
	
	   http.ListenAndServe(":"+port, nil)
	// fmt.Println("This is mY player score management systemK")
	// DB connection
	// db, err := dbhelper.CreateConnection()
	// if err != nil {
	// 	return
	// }
	// //    incase if you need to run locally uncommented this line
	// tab, err := dbhelper.CreateTable("Player")
	// fmt.Println(tab)
	// // database connection close
	// defer db.Close()
	// // create instance of gin framwork
	// r := gin.New()
	// // register routes to handle router
	// router.Router(r)
	// // server running on port 8000
	// r.Run(":8080")

}
