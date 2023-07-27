package utils

import (
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
)

func LoadFonts(m pdf.Maroto) {
	m.AddUTF8Font("Outfit", consts.Normal, "fonts/Outfit/Outfit-Regular.ttf")
	m.AddUTF8Font("Outfit", consts.Bold, "fonts/Outfit/Outfit-Bold.ttf")
	m.AddUTF8Font("Outfit", consts.Style("Light"), "fonts/Outfit/Outfit-Light.ttf")
}
