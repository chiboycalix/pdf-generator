package main

import (
	"fmt"
	"os"

	"github.com/chiboycalix/pdf-generator/content"
	"github.com/chiboycalix/pdf-generator/utils"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
)

func main() {
	m := pdf.NewMaroto(consts.Portrait, consts.A3)
	utils.SetPageMargins(m)
	utils.LoadFonts(m)
	m.Row(200, func() {
		content.LeftContent(m)
		content.RightContent(m)
	})
	err := m.OutputFileAndClose("pdfs/resume.pdf")
	if err != nil {
		fmt.Println("Could not save PDF:", err)
		os.Exit(1)
	}
	fmt.Print("PDF generated successfully")
}
