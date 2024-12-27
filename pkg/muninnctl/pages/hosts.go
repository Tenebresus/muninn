package pages

import (
	"io"
	"net/http"
    "strings"

    "github.com/tenebresus/muninn/pkg/config"
	tea "github.com/charmbracelet/bubbletea"
)

var selectedIp string

type ListTSMsg []byte

type Hosts struct {

    Timestamps []string
    SelectedIp string
    CurrentCursorPos int
    SelectedTS string

}

func (h Hosts) View() string {

    ret := "\t\tSelected IP: " + h.SelectedIp + "\n\n\n"

    for i, ts := range h.Timestamps {

        cursor := "  "
        if h.CurrentCursorPos == i {
            cursor = "> "
        }

        ret += cursor + ts + "\n"

    }

    ret += "\n\n\t\t[q] quit\t[h] home"
    
    return ret

}

func (h *Hosts) UpdateSelectIp(ip string) {

    selectedIp = ip
    h.SelectedIp = ip

}

func (h *Hosts) CursorUp() {

    if h.CurrentCursorPos > 0 {
        h.CurrentCursorPos--
    }

}

func (h *Hosts) CursorDown() {

    if h.CurrentCursorPos < len(h.Timestamps) - 1 {
        h.CurrentCursorPos++
    }

}

func (h *Hosts) SelectTS() {
    h.SelectedTS = h.Timestamps[h.CurrentCursorPos]
}

func GetHostTS() tea.Msg {

    muninnHost := config.GetMuninnHost()
    ret, _ := http.Get("http://" + muninnHost + ":8081/scavenge/" + selectedIp) 
    ipTS, _ := io.ReadAll(ret.Body)

    return ListTSMsg(ipTS)

}  

func GetHostTSFromMsg(msg string) []string {

    list := string(msg)
    lines := strings.Split(list, "\n")

    var listSlice []string

    for _, line := range lines {

        if line != "" {

            listSlice = append(listSlice, line)

        }

    }

    return listSlice

}
