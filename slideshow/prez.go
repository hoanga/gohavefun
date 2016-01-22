package main

import tui "github.com/gizak/termui"
import "strconv"
import "math"

func calcPercent(curr int, total int) int {
	return (curr * 100) / (total - 1)
}

func cosData() []float64 {
	n := 200
	ps := make([]float64, n)
	for i := range ps {
		ps[i] = 1 + math.Cos(float64(i)/5)
	}
	return ps
}


type SlideShow struct {
	Slides []Slide
	Title tui.Bufferer
	Footer tui.Bufferer
}

type Slide struct {
	Title string
	Widgets []tui.Bufferer
}

func NewSlideShow(title tui.Bufferer, footer tui.Bufferer) SlideShow {
	s := SlideShow{
		Slides: make([]Slide, 0),
		Title: title,
		Footer: footer,
	}
	return s
}

func (sl* SlideShow) AddSlide(title string, slides []tui.Bufferer) {
	newslide := Slide{Title: title,
		Widgets: make([]tui.Bufferer, 0)}

	newslide.Widgets = append(newslide.Widgets, sl.Title)
	newslide.Widgets = append(newslide.Widgets, slides...)
	newslide.Widgets = append(newslide.Widgets, sl.Footer)
	sl.Slides = append(sl.Slides, newslide)
}

func (sl* SlideShow) Length() int {
	return len(sl.Slides)
}

func (sl* SlideShow) At(i int) Slide {
	return sl.Slides[i]
}

func main() {
	err := tui.Init()
	if err != nil {
		panic(err)
	}
	defer tui.Close()

	///////////
	// 
	//  Create UI components
	//
	///////////

	// Header
	pr_th := 3
	pr_title := tui.NewPar("Text Console User Interfaces")
	pr_title.Width = tui.TermWidth()
	pr_title.Height = pr_th
	pr_title.BorderFg = tui.ColorBlue

	// Footer
	g_h := 5
	g := tui.NewGauge()
	g.Percent = 1
	g.Width = tui.TermWidth()
	g.Height = g_h
	g.Y = tui.TermHeight() - g_h
	g.BorderLabel = "Progress"
	g.Label = "{{percent}} - Start!"
	g.LabelAlign = tui.AlignRight
	g.BarColor = tui.ColorGreen
	g.BorderFg = tui.ColorBlue
	g.BorderLabelFg = tui.ColorWhite

	// Slide 1
	txtlst1 := "Introduction\n\no Myself\n\no Interests in Go"
	se1_1 := tui.NewPar(txtlst1)
	se1_1.Width = tui.TermWidth()
	se1_1.Height = (tui.TermHeight() / 2) - (pr_th + g_h)
	se1_1.Y = pr_th
	se1_2 := tui.NewPar("")
	se1_2.Width = tui.TermWidth()
	se1_2.Height = (tui.TermHeight() / 2) - (pr_th + g_h)
	se1_2.Y = pr_th + se1_1.Height

	// Slide 2
	txtlst2 := "The Termui Library\n\no Console library UI\n\n"
	txtlst2 += "o A widget library for dashboard building in the terminal\n\n"
	txtlst2 += "o Cross Platform\n\n  o Runs on Linux, OSX, and Windows"
	se2_1 := tui.NewPar(txtlst2)
	se2_1.Width = tui.TermWidth()
	se2_1.Height = (tui.TermHeight() / 2) - (pr_th + g_h)
	se2_1.Y = pr_th

	// Slide 3
	txtlst3 := "More Info\n\n"
	txtlst3 += "o Built on top of termbox library\n\n"
	txtlst3 += "o Inherits handlers, events, and cross platform compatibility"
	se3_1 := tui.NewPar(txtlst3)
	se3_1.Width = tui.TermWidth()
	se3_1.Height = (tui.TermHeight() / 2) - (pr_th + g_h)
	se3_1.Y = pr_th

	// Slide 4
	txtlst4 := "Features\n\n"
	txtlst4 += "o Multiple widgets available\n\n"
	txtlst4 += "o Automatic grid layout\n\n"
	txtlst4 += "o 多言語可能 (multi-lang possible)"
	se4_1 := tui.NewPar(txtlst4)
	se4_1.Width = tui.TermWidth()
	se4_1.Height = (tui.TermHeight() / 2) - (pr_th + g_h)
	se4_1.Y = pr_th

	// Slide 5
	txtlst5 := "Widget Features\n\n"
	txtlst5 += "o Can be surrounded by borders\n\n"
	txtlst5 += "o Can have labels associated with it\n\n"
	txtlst5 += "o Borders can also have labels\n\n"
	txtlst5 += "o Color"
	se5_1 := tui.NewPar(txtlst5)
	se5_1.Width = tui.TermWidth()
	se5_1.Height = (tui.TermHeight() / 2) - (pr_th + g_h)
	se5_1.Y = pr_th

	// Slide 6
	txtlst6 := "Widgets - Par\n\no Par - aka Textbox\n\n"
	txtlst6 += "o Basic textbox widget\n\n"
	txtlst6 += "   p := termui.NewPar(\"World\")\n"
	txtlst6 += "   p.BorderLabel(\"Hello\")"
	se6_1 := tui.NewPar(txtlst6)
	se6_1.Width = tui.TermWidth()
	se6_1.Height = (tui.TermHeight() / 2) - (pr_th + g_h)
	se6_1.Y = pr_th
	se6_2 := tui.NewPar("World")
	se6_2.BorderLabel = "Hello"
	se6_2.BorderFg = tui.ColorYellow
	se6_2.BorderLabelFg = tui.ColorWhite
	se6_2.Width = tui.TermWidth()
	se6_2.Height = (tui.TermHeight() / 2) - (pr_th + g_h)
	se6_2.Y = pr_th + se6_1.Height

	// Slide 7
	txtlst7 := "Widgets - Lists\n\no List - A text list\n\n"
	txtlst7 += "o Text Lists\n\n"
	txtlst7 += "   tl := termui.NewList()\n"
	txtlst7 += "   tl.Items = textlist\n"
	se7_2lst := []string {
		"* List Elems",
		"* Are Just",
		"* Lists of",
		"* Strings",
		"* [and support](fg-blue)",
		"* [colors](fg-green,bg-black)"}
	se7_1 := tui.NewPar(txtlst7)
	se7_1.Width = tui.TermWidth()
	se7_1.Height = (tui.TermHeight() / 2) - (pr_th + g_h)
	se7_1.Y = pr_th
	se7_2 := tui.NewList()
	se7_2.Items = se7_2lst
	se7_2.BorderFg = tui.ColorYellow
	se7_2.BorderLabelFg = tui.ColorWhite
	se7_2.Width = tui.TermWidth()
	se7_2.Height = (tui.TermHeight() / 2) - (pr_th + g_h)
	se7_2.Y = pr_th + se7_1.Height

	// Slide 8
	txtlst8 := "Widgets - Line Charts\n\n"
	txtlst8 += "o Draw linecharts\n\n"
	txtlst8 += "   lc := termui.NewLineChart()\n"
	txtlst8 += "   lc.Data = cosdata"
	se8_1 := tui.NewPar(txtlst8)
	se8_1.Width = tui.TermWidth()
	se8_1.Height = (tui.TermHeight() / 2) - (pr_th + g_h)
	se8_1.Y = pr_th
	se8_2 := tui.NewLineChart()
	se8_2.Data = cosData()
	se8_2.BorderFg = tui.ColorYellow
	se8_2.BorderLabelFg = tui.ColorWhite
	se8_2.Width = tui.TermWidth()
	se8_2.Height = (tui.TermHeight() / 2) - (pr_th + g_h)
	se8_2.Y = pr_th + se8_1.Height

	// Slide 9
	txtlst9 := "Widgets - Bar Charts\n\n"
	txtlst9 += "o Draw bar charts\n\n"
	txtlst9 += "   data := []int{4, 5, 6, 7, 8, 6, 5}\n"
	txtlst9 += "   bc := termui.NewBarChart()\n"
	txtlst9 += "   bc.Data = data"
	se9_1 := tui.NewPar(txtlst9)
	se9_1.Width = tui.TermWidth()
	se9_1.Height = (tui.TermHeight() / 2) - (pr_th + g_h)
	se9_1.Y = pr_th
	se9_2 := tui.NewBarChart()
	se9_2.Data = []int{4, 5, 6, 7, 8, 6, 5}
	se9_2.DataLabels = []string{"S0", "S1", "S2", "S3", "S4", "S5", "S6", "S7"}
	se9_2.BorderFg = tui.ColorYellow
	se9_2.BorderLabelFg = tui.ColorWhite
	se9_2.Width = tui.TermWidth()
	se9_2.Height = (tui.TermHeight() / 2) - (pr_th + g_h)
	se9_2.Y = pr_th + se9_1.Height

	// Slide 10
	txtlst10 := "Widgets - Sparklines\n\n"
	txtlst10 += "o Draw sparklines\n\n"
	txtlst10 += "   data := []int{4, 5, 6, 7, 8, 6, 5}\n"
	txtlst10 += "   sp := termui.NewSparkline()\n"
	txtlst10 += "   sp.Data = data\n"
	txtlst10 += "   spl := termui.NewSparklines(sp)"
	se10_1 := tui.NewPar(txtlst10)
	se10_1.Width = tui.TermWidth()
	se10_1.Height = (tui.TermHeight() / 2) - (pr_th + g_h)
	se10_1.Y = pr_th
	sp10_2 := tui.NewSparkline()
	sp10_2.Data = []int{4, 5, 6, 7, 8, 6, 5}
	sp10_2.LineColor = tui.ColorRed
	se10_2 := tui.NewSparklines(sp10_2)
	se10_2.Width = tui.TermWidth()
	se10_2.Height = (tui.TermHeight() / 2) - (pr_th + g_h)
	se10_2.Y = pr_th + se10_1.Height

	// Slide 11
	txtlst11 := "General Workflow\n\n"
	txtlst11 += "o Setup\n\no Create & Setup UI elems\n\n"
	txtlst11 += "o Setup handlers\n\nLoop"
	se11_1 := tui.NewPar(txtlst11)
	se11_1.Width = tui.TermWidth()
	se11_1.Height = (tui.TermHeight() / 2) - (pr_th + g_h)
	se11_1.Y = pr_th
	txtlst11_2 := "   termui.Init()\n"
	txtlst11_2 += "   p := termui.NewPar(\"Hello World\")\n"
	txtlst11_2 += "   termui.Render(p)\n"
	txtlst11_2 += "   termui.Handle(\"/sys/kbd/Q\", func(termui.Event) {\n"
	txtlst11_2 += "             termui.StopLoop() })\n"
	txtlst11_2 += "   termui.Loop()"
	se11_2 := tui.NewPar(txtlst11_2)
	se11_2.BorderFg = tui.ColorYellow
	se11_2.BorderLabelFg = tui.ColorWhite
	se11_2.Width = tui.TermWidth()
	se11_2.Height = (tui.TermHeight() / 2) - (pr_th + g_h)
	se11_2.Y = pr_th + se11_1.Height

	// Slide 12
	txtlst12 := "Extra Notes\n\n"
	txtlst12 += "o V1 vs V2\n\n"
	txtlst12 += "o Timers\n\n"
	se12_1 := tui.NewPar(txtlst12)
	se12_1.Width = tui.TermWidth()
	se12_1.Height = (tui.TermHeight() / 2) - (pr_th + g_h)
	se12_1.Y = pr_th
	se12_2 := tui.NewPar("")
	se12_2.Border = false
	se12_2.Width = tui.TermWidth()
	se12_2.Height = (tui.TermHeight() / 2) - (pr_th + g_h)
	se12_2.Y = pr_th + se12_1.Height

	///////////
	// 
	//  Build the slideshow
	//
	///////////
	slides := NewSlideShow(pr_title, g)
	slides.AddSlide("start", []tui.Bufferer{})
	slides.AddSlide("Intro", []tui.Bufferer{se1_1})
	slides.AddSlide("Termui", []tui.Bufferer{se2_1})
	slides.AddSlide("More Info", []tui.Bufferer{se3_1})
	slides.AddSlide("Features", []tui.Bufferer{se4_1})
	slides.AddSlide("WidgetFeatures", []tui.Bufferer{se5_1})
	slides.AddSlide("Par()", []tui.Bufferer{se6_1, se6_2})
	slides.AddSlide("List()", []tui.Bufferer{se7_1, se7_2})
	slides.AddSlide("LineChart()", []tui.Bufferer{se8_1, se8_2})
	slides.AddSlide("BarChart()", []tui.Bufferer{se9_1, se9_2})
	slides.AddSlide("Sparkline()", []tui.Bufferer{se10_1, se10_2})
	slides.AddSlide("quickdemo", []tui.Bufferer{se11_1, se11_2})
	slides.AddSlide("gotchas", []tui.Bufferer{se12_1, se12_2})

	slides_num := slides.Length()
	slides_idx := 0

	maxttl := 360
	ttl := maxttl

	draw := func() {
		tui.Render(slides.At(slides_idx).Widgets...)
	}

	tui.Render(pr_title, g)

	tui.Handle("/sys/kbd/Q", func(tui.Event) {
		tui.StopLoop()
	})
	tui.Handle("/sys/kbd/<left>", func(tui.Event) {
	})
	tui.Handle("/sys/kbd/<right>", func(tui.Event) {
		ttl = maxttl
		slides_idx++
		if slides_idx > (slides_num-1) {
			slides_idx = 0
		}
		g.Percent = calcPercent(slides_idx, slides_num)

		lbl := "Progress - " + strconv.Itoa(g.Percent) + "%" +
			" TTL: " + strconv.Itoa(ttl)

		g.BorderLabel = lbl
		g.Label = "{{percent}} - " + slides.At(slides_idx).Title
	})
	tui.Handle("/sys/kbd/<space>", func(tui.Event) {
		ttl = maxttl
		slides_idx++
		if slides_idx > (slides_num-1) {
		 	slides_idx = 0
		}
		g.Percent = calcPercent(slides_idx, slides_num)

		lbl := "Progress - " + strconv.Itoa(g.Percent) + "%" +
			" TTL: " + strconv.Itoa(ttl)

		g.BorderLabel = lbl
		g.Label = "{{percent}} - " + slides.At(slides_idx).Title
	})
	tui.Handle("/timer/1s", func(e tui.Event) {
		ttl--
		if ttl <= 0 {
			if slides_idx < (slides_num-1) {
				if slides_idx > slides.Length() - 1 {
					slides_idx++
				}
			}
			g.Percent = calcPercent(slides_idx, slides_num)
			ttl = maxttl
		}
		lbl := "Progress - " + strconv.Itoa(g.Percent) + "%" +
			" TTL: " + strconv.Itoa(ttl)
		g.BorderLabel = lbl

		draw()
	})

	tui.Loop()
}
