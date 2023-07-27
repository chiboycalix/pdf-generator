package utils

import "github.com/johnfercher/maroto/pkg/pdf"

func SetPageMargins(m pdf.Maroto) {
	m.SetPageMargins(10, 15, 10)
}
