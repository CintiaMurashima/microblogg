package main
 
import (

    "log"
    "github.com/CintiaMurashima/microblogging/handlers"
    "github.com/CintiaMurashima/microblogging/bd"
)
func main(){
    if bd.ChequeoConnection() == 0 {
        log.Fatal("Sin conexión a la BD")
        return
    }
    handlers.Manejadores()
 
}
