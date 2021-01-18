package handlers
 
import (
    "log"
    "net/http"
    "os"
	"github.com/gorilla/mux"
	"github.com/rs/cors" 
	"github.com/CintiaMurashima/microblogg/middlew"
    "github.com/CintiaMurashima/microblogg/routers"

)
/*Manejadores: seteo mi puerto, el handler y pongo a escuchar al servidor*/
func Manejadores(){
    router := mux.NewRouter()
	/*Por cada endPoint vamos a tener un reglon de cod que permita */
	router.HandleFunc("/registro", middlew.ChequeoBD(routers.Registro)).Methods("POST")
	PORT := os.Getenv("PORT")
    if PORT == ""{
        PORT = "8080"
    }
    handler := cors.AllowAll().Handler(router)
    log.Fatal(http.ListenAndServe(":"+PORT,handler))
}
