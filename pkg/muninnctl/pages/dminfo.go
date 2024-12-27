package pages

import (
	"encoding/json"
	"io"
	"net/http"
	"slices"

	"github.com/Tenebresus/dmidegoder/parser"
	"github.com/tenebresus/muninn/pkg/config"
	tea "github.com/charmbracelet/bubbletea"
)

var selectedTS string

var dmiNameFilter = []string {

    "Processor",
    "Memory Device",

}

type GetDmiMsg []byte

type Dminfo struct {

    SelectedTS string
    Dmidecode []byte
    DmiTypes []parser.DMIType
    FilteredDmiTypes []parser.DMIType
    CurrentCursorPos int

}

func (d *Dminfo) UpdateDMITypes() {

    json.Unmarshal(d.Dmidecode, &d.DmiTypes)

}

func (d *Dminfo) FilterDMITypes() {

    var filteredDmiTypes []parser.DMIType

    for _, dmiType := range d.DmiTypes {

        if slices.Contains(dmiNameFilter, dmiType.Name) {

            filteredDmiTypes = append(filteredDmiTypes, dmiType)  

        }

    }

    d.FilteredDmiTypes = filteredDmiTypes 

}

func (d *Dminfo) UpdateSelectedTs(ts string) {

    selectedTS = ts
    d.SelectedTS = ts

}

func (d Dminfo) View() string {

    ret := "\t\tSelected TS: " + d.SelectedTS + "\n\n"

    for i, dmiType := range d.FilteredDmiTypes {

        cursor := "  "
        if i == d.CurrentCursorPos{
            cursor = "> "
        }

        ret += cursor + dmiType.Name + "\n"

    }

    ret += "\t\t[q] quit\t[h] home"

    return ret

}

func (d *Dminfo) CursorUp() {

    if d.CurrentCursorPos > 0  {
        d.CurrentCursorPos--
    }

}

func (d *Dminfo) CursorDown() {

    if d.CurrentCursorPos < len(d.FilteredDmiTypes) - 1  {
        d.CurrentCursorPos++
    }

}

func GetTSDmi() tea.Msg {

    muninnHost := config.GetMuninnHost()
    ret, _ := http.Get("http://" + muninnHost + ":8081/scavenge/" + selectedIp + "/" + selectedTS) 
    tsDmi, _ := io.ReadAll(ret.Body)

    return GetDmiMsg(tsDmi)

}


