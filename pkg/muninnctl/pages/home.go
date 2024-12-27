package pages

import (
    tea "github.com/charmbracelet/bubbletea"
    "github.com/tenebresus/muninn/pkg/config"
    "strings"
    "io"
    "net/http"
)

type ListHostsMsg []byte

type HomePage struct {

   MuninnHostIp string
   HostIPs []string 
   CurrentCursorPos int
   SelectedIP string

}

func (h HomePage) View() string {

    ret := "\t\t\tWelcome to Muninn!\n\n\n"
    ret += "\t\tCurrent muninn host: " + h.MuninnHostIp  + "\n\n"

    for i, ip := range h.HostIPs {

        cursor := "  "
        if h.CurrentCursorPos == i {
            cursor = "> "
        }

        ret += cursor + ip + "\n"

    }

    ret += "\n\t\t[q] quit"

    return ret

}

func (h *HomePage) CursorUp() {

    if h.CurrentCursorPos > 0 {
        h.CurrentCursorPos--
    }

}

func (h *HomePage) CursorDown() {

    if h.CurrentCursorPos < len(h.HostIPs) - 1 {
        h.CurrentCursorPos++
    }

}

func (h *HomePage) SelectIp() {

   h.SelectedIP = h.HostIPs[h.CurrentCursorPos] 

}

func GetHostListFromMsg(msg string) []string {

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

func GetHostList() tea.Msg {

    muninnHost := config.GetMuninnHost()
    ret, _ := http.Get("http://" + muninnHost + ":8081/scavenge") 
    lists, _ := io.ReadAll(ret.Body)

    return ListHostsMsg(lists) 

}

