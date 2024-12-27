package agent

import (
	"net/http"
	"os/exec"

	"github.com/Tenebresus/dmidegoder/parser"
)

func Run() {

    handler := http.NewServeMux()
    handler.HandleFunc("GET /dmidecode", dmi)

    server := http.Server{

        Addr: ":8080",
        Handler: handler,

    }

   server.ListenAndServe() 

}

func dmi(w http.ResponseWriter, r *http.Request) {
    w.Write(dmidecode())
}

func dmidecode() []byte {

    exec := exec.Command("dmidecode")
    output, _ := exec.Output()

    return parser.Parse(string(output))

}
