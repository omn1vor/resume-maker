package renderer

import (
	"bytes"
	"html/template"
	"strings"

	wkhtml "github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/omn1vor/resume-maker/internal/model"
)

func RenderToPDF(templatePath string, r *model.Resume, lang string, output string) error {

	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		return err
	}

	data := struct {
		Resume model.Resume
		T      model.Translations
	}{
		Resume: *r,
		T:      model.GetTranslations(lang),
	}

	var htmlBuf bytes.Buffer
	if err := tmpl.Execute(&htmlBuf, data); err != nil {
		return err
	}

	pdfg, err := wkhtml.NewPDFGenerator()
	if err != nil {
		return err
	}

	page := wkhtml.NewPageReader(strings.NewReader(htmlBuf.String()))
	page.EnableLocalFileAccess.Set(true)

	pdfg.AddPage(page)

	if err := pdfg.Create(); err != nil {
		return err
	}

	return pdfg.WriteFile(output)
}
