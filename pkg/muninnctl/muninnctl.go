package muninnctl

import (
	tea "github.com/charmbracelet/bubbletea"
    "github.com/tenebresus/muninn/pkg/muninnctl/pages"
	"github.com/tenebresus/muninn/pkg/config"
)

type dmiMsg []byte

type MuninnModel struct {

    HomePage pages.HomePage
    HostsPage pages.Hosts
    DmiPage pages.Dminfo
    CurrentPage string

}

type CursorMover interface {

    CursorUp()
    CursorDown()

}

func (m MuninnModel) Init() tea.Cmd {
    return tea.Batch(pages.GetHostList, tea.EnterAltScreen)
}

func (m MuninnModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    
    switch msg := msg.(type) {

    // HOME PAGE
    case pages.ListHostsMsg:

        m.HomePage.HostIPs = pages.GetHostListFromMsg(string(msg))

    // HOST TS CurrentPage
    case pages.ListTSMsg:

        m.HostsPage.Timestamps = pages.GetHostTSFromMsg(string(msg))

    case pages.GetDmiMsg:

        m.DmiPage.Dmidecode = msg
        m.DmiPage.UpdateDMITypes()
        m.DmiPage.FilterDMITypes()
        
    // ON KEY PRESS
    case tea.KeyMsg:

        switch msg.String(){

        case "q":

            return m, tea.Batch(tea.Quit, tea.ExitAltScreen)

        case "k", "up":

            cursorMover := m.getCursorMover()
            cursorMover.CursorUp()

        case "j", "down":

            cursorMover := m.getCursorMover()
            cursorMover.CursorDown()

        case "e":

            cmdMsg := m.onEnterKey(m.CurrentPage)
            return m, tea.Batch(tea.ClearScreen, cmdMsg)

        case "h":

            if m.CurrentPage != "home" {
                m.CurrentPage = "home"
                return m, tea.ClearScreen
            }

        }

    }

    return m, nil

}

func (m *MuninnModel) getCursorMover() CursorMover {

    if m.CurrentPage == "home" {
        return &m.HomePage
    }

    if m.CurrentPage == "listHostTS" {
        return &m.HostsPage
    }

    if m.CurrentPage == "getTSDmi" {
        return &m.DmiPage
    }

    return nil 

}

func (m MuninnModel) View() string {

    if m.CurrentPage == "listHostTS" {
       return m.HostsPage.View() 
    }

    if m.CurrentPage == "getTSDmi" {
       return m.DmiPage.View() 
    }

    return m.HomePage.View() 
}

func (m *MuninnModel) onEnterKey(page string) tea.Cmd {

    switch page {

    case "home":

        m.CurrentPage = "listHostTS"
        m.HomePage.SelectIp()
        m.HostsPage.UpdateSelectIp(m.HomePage.SelectedIP)

        return pages.GetHostTS

    case "listHostTS":
        
        m.CurrentPage = "getTSDmi"
        m.HostsPage.SelectTS()
        m.DmiPage.UpdateSelectedTs(m.HostsPage.SelectedTS)

        return pages.GetTSDmi

    }

    return nil

}

func InitializeModel() MuninnModel {

    config := config.Get()
    ret := MuninnModel{
        HomePage: pages.HomePage{
            MuninnHostIp: config.MuninnHostIp,
        },
        CurrentPage: "home",
    }

    return ret

}

