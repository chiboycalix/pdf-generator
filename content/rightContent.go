package content

import (
	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
)

func RightContent(m pdf.Maroto) {
	m.Col(9, func() {
		m.Text("Lorem right Lorem right Lorem right Lorem rightLorem right Lorem right Lorem right Lorem right Lorem right Lorem rightLorem right Lorem right", props.Text{
			Color: color.NewBlack(),
		})
	})
}
