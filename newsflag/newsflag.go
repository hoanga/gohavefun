package main

import ui "github.com/gizak/termui"

import (
	"time"
	"strings"
	"bytes"
	"io/ioutil"
	"net/http"

	"golang.org/x/net/html"	
)

func getData(t html.Token, dt html.Token) (ok bool, href string, descr string) {
	for _, a := range t.Attr {
		if a.Key == "href" {
			href = a.Val
			descr = dt.String()
		}
		if a.Key == "class" {
			if strings.HasPrefix(a.Val, "title may-blank") {
				ok = true
			}
		}
	}
	return
}

type Loi struct {
	url string
	descr string
}

func (l *Loi) String() string {
	return l.url + "|" + l.descr
}

func (l *Loi) IsInteresting() bool {
	soi := []string{ "paris", "france" }

	for _, s := range soi {
		if strings.Contains(strings.ToLower(l.url), s) {
			return true
		}
		if strings.Contains(strings.ToLower(l.descr), s) {
			return true
		}
	}
	return false
}

func readLinks(r *bytes.Reader, il []Loi) []Loi {
	ht := html.NewTokenizer(r)


htmlsearch:
	for {
		tt := ht.Next()

		switch {
		case tt == html.ErrorToken:
			break htmlsearch
		case tt == html.StartTagToken:
			t := ht.Token()
			if t.Data != "a" {
				continue
			}
			tt = ht.Next()
			ok, url, descr := getData(t, ht.Token())
			if ok {
				l := Loi{url, descr}
				if l.IsInteresting() {
					il = append(il, l)
				}
			}
		}
	}

	return il
}

func fetchNews(uc chan bool, lc chan Loi, links []Loi, url string) {
	var doUpdates bool

	for {
		select {
		case doUpdates = <- uc:
			break
		default:
			if doUpdates {
				r, e := http.Get(url)
				if e != nil {
					break
				}
				b, e2 := ioutil.ReadAll(r.Body)
				if e2 != nil {
					break
				}
				links = readLinks(bytes.NewReader(b), links)
				for _, l := range links {
					lc <- l
				}
			}
		}
		time.Sleep(time.Millisecond * 500)
	}
}

func rotateLists(lists []*ui.List, fwd bool) []*ui.List {
	nl := []*ui.List{}
	if fwd {
		nl = append(nl, lists[2])
		nl = append(nl, lists[0])
		nl = append(nl, lists[1])
	} else {
		nl = append(nl, lists[1])
		nl = append(nl, lists[2])
		nl = append(nl, lists[0])
	}
	return nl
}

func main() {
	////
	//
	// Non-UI section
	//
	////
	var tick int
	var updateUi bool
	var updateNews bool
	var lc chan Loi = make(chan Loi)
	var uc chan bool = make(chan bool, 3)
	links := make([]Loi, 0)

	////
	//
	// Setup UI elements
	//
	////	
	e := ui.Init()
	if e != nil {
		panic(e)
	}
	defer ui.Close()

	infoStr := "Help - '?' Show/Hide help, 'u' Show curr url, 'l' Toggle live feed, <space> Toggle ui updates"
	bstr := make([]string, ui.TermHeight() - 1)
	wstr := make([]string, ui.TermHeight() - 1)
	rstr := make([]string, ui.TermHeight() - 1)
	fillb := make([]byte, ui.TermWidth()/3-2)
	for i := 0; i < ui.TermHeight() - 1; i ++ {
		bs := make([]byte, ui.TermWidth()/3-2)
		ws := make([]byte, ui.TermWidth()/3-2)
		rs := make([]byte, ui.TermWidth()/3-2)		
		for j := 0; j < ui.TermWidth()/3-2; j++ {
			bs[j] = '0'
			ws[j] = '0'
			rs[j] = '0'			
		}
		bstr[i] = string(bs)
		wstr[i] = string(ws)
		rstr[i] = string(rs)
	}
	copy(fillb, bstr[0])
	fillstr := string(fillb)
	blst := ui.NewList()
	blst.HasBorder = true
	blst.Border.FgColor = ui.ColorBlue
	blst.Border.BgColor = ui.ColorBlue
	blst.ItemFgColor = ui.ColorBlue
	blst.Items = bstr
	blst.Height = ui.TermHeight() - 1

	wlst := ui.NewList()
	wlst.HasBorder = true	
	wlst.Border.FgColor = ui.ColorWhite
	wlst.Border.BgColor = ui.ColorWhite
	wlst.ItemFgColor = ui.ColorWhite
	wlst.Items = wstr	
	wlst.Height = ui.TermHeight() - 1
	
	rlst := ui.NewList()
	rlst.HasBorder = true
	rlst.Border.FgColor = ui.ColorRed
	rlst.Border.BgColor = ui.ColorRed
	rlst.ItemFgColor = ui.ColorRed
	rlst.Items = rstr	
	rlst.Height = ui.TermHeight() - 1

	inf := ui.NewPar("")
	inf.HasBorder = false
	inf.Height = 1
	
	// build layout
	ui.Body.AddRows(
		ui.NewRow(
			ui.NewCol(4, 0, blst),
			ui.NewCol(4, 0, wlst),
			ui.NewCol(4, 0, rlst)),
		ui.NewRow(
			ui.NewCol(12, 0, inf)) )
	// Calc layout
	ui.Body.Align()
	ev := ui.EventCh()
	
	rotateList := []*ui.List{blst, wlst, rlst}	
	// Render
	draw := func(t int) {
		if updateUi {
			select {
			case ll := <- lc:
				rotateList[0].Items[t] = ll.url
				rotateList[1].Items[t] = ll.descr
				rotateList[2].Items[t] = fillstr
			default:
			}
		}
		ui.Render(ui.Body)	
	}
	
	////
	//
	// Main loop
	//
	////
	go fetchNews(uc, lc, links, "https://reddit.com/r/news")
	for {
		select {
		case evt := <- ev:
			if evt.Type == ui.EventKey &&
				evt.Ch == 'i' {
			} else if evt.Type == ui.EventKey &&
				evt.Ch == 'l' {
				updateNews = !updateNews
				uc <- updateNews
			} else if evt.Type == ui.EventKey &&
				evt.Ch == 'u' {
				ss := rotateList[0].Items[0]
				if !strings.HasPrefix(ss, fillstr) {
					inf.Text = ss
				}
			} else if evt.Type == ui.EventKey &&
				evt.Ch == '?' {
				if strings.HasPrefix(inf.Text, infoStr) {
					inf.Text = ""
				} else {
					inf.Text = infoStr
				}
			} else if evt.Type == ui.EventKey &&
				evt.Key == ui.KeySpace {
				updateUi = !updateUi
			} else if evt.Type == ui.EventKey &&
				evt.Key == ui.KeyArrowLeft {
				rotateList = rotateLists(rotateList, true)
			} else if evt.Type == ui.EventKey &&
				evt.Key == ui.KeyArrowRight {
				rotateList = rotateLists(rotateList, false)	
			} else	if evt.Type == ui.EventKey &&
				evt.Ch == 'Q' {
				return
			}			
		default:
			time.Sleep(time.Millisecond * 400)
			tick += 1
			if tick > (ui.TermHeight()-3) {
				tick = 0
			}			
			draw(tick)
		}
	}	

}
