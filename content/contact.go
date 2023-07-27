package content

import (
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
)

func generateContact(m pdf.Maroto) {
	m.Row(50, func() {
		m.Col(1, func() {
			m.Text("Mail", props.Text{})
		})

		m.Col(2, func() {
			m.Text("igwechinonso77@gmail.com", props.Text{})
		})
	})
}
