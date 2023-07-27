package content

import (
	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
)

var TopMargin = 55.00

func LeftContent(m pdf.Maroto) {
	m.Col(3, func() {
		m.FileImage("images/pic.jpg", props.Rect{
			Percent: 75,
		})
		m.Text("Igwe Chinonso", props.Text{
			Color:  color.NewBlack(),
			Family: "Outfit",
			Style:  consts.Bold,
			Top:    TopMargin,
			Size:   25,
		})
		m.Text("Full Stack Engineer", props.Text{
			Style:  consts.Normal,
			Top:    TopMargin + 12,
			Size:   20,
			Family: "Outfit",
			Color: color.Color{
				Red:   164,
				Green: 120,
				Blue:  232,
			},
		})

		m.Text("I am a full stack engineer", props.Text{
			Top:    TopMargin + 30,
			Family: "Outfit",
			Size:   16,
			Style:  consts.Style("Light"),
			Align:  consts.Left,
		})

		m.Col(1, func() {
			m.Text("Mail", props.Text{
				Top:    TopMargin*2 + 20,
				Size:   14,
				Family: "Outfit",
			})
		})

		m.Col(1, func() {
			m.Text("igwechinonso77@gmail.com", props.Text{
				Top:    TopMargin*2 + 20,
				Size:   14,
				Family: "Outfit",
			})
		})
	})
}
