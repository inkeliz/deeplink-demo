package main

import (
	"gioui.org/app"
	"gioui.org/font"
	"gioui.org/io/transfer"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/text"
	"gioui.org/widget"
	"image"
	"image/color"
	"log"
	"os"
)

func main() {
	w := &app.Window{}
	lt := text.NewShaper()

	var texts = []string{"Schemes:"}

	ops := new(op.Ops)

	go func() {
		for {
			evt := w.Event()

			switch e := evt.(type) {
			case app.DestroyEvent:
				os.Exit(0)
				return
			case app.FrameEvent:
				gtx := app.NewContext(ops, e)

				for {
					evt, ok := gtx.Event(transfer.URLFilter{})
					if !ok {
						break
					}

					if e, ok := evt.(transfer.URLEvent); ok {
						texts = append(texts, e.URL.String())
					}
				}

				log.Println("FrameEvent")
				s := clip.RRect{Rect: image.Rectangle{Max: gtx.Constraints.Max}}.Push(gtx.Ops)
				paint.ColorOp{Color: color.NRGBA{R: 255, G: 255, A: 255}}.Add(gtx.Ops)
				paint.PaintOp{}.Add(gtx.Ops)
				s.Pop()

				offset := image.Pt(0, 0)

				p := op.Record(gtx.Ops)
				paint.ColorOp{Color: color.NRGBA{A: 255}}.Add(gtx.Ops)
				painter := p.Stop()

				for _, txt := range texts {
					gtx.Constraints.Min.Y = 0
					gtx.Constraints.Min.X = 0

					o := op.Offset(offset).Push(gtx.Ops)
					dims := widget.Label{}.Layout(gtx, lt, font.Font{}, 16, txt, painter)
					o.Pop()
					offset.Y += dims.Size.Y
				}

				e.Frame(ops)
			}
		}
	}()

	app.Main()
}
