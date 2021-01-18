package main
 
import (

    "log"
    "github.com/CintiaMurashima/microblogg/handlers"
    "github.com/CintiaMurashima/microblogg/bd"
)
func main(){
    if bd.ChequeoConnection() == 0 {
        log.Fatal("Sin conexión a la BD")
        return
    }
    handlers.Manejadores()
 
}
