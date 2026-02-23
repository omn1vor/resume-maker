package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/omn1vor/resume-maker/internal/parser"
	"github.com/omn1vor/resume-maker/internal/registry"
	"github.com/omn1vor/resume-maker/internal/renderer"
	"github.com/omn1vor/resume-maker/internal/validator"
)

const (
	defaultMessage = `usage: 
- resume-maker templates: shows available templates
- resume-maker build resume.yaml [-template=template-name] [-lang=en] [-out=pdf-output-file-name.pdf]: creates a pdf resume. 
Use resume-maker build -help for more info
`
)

func main() {
	reg, err := registry.New()
	if err != nil {
		closeWithErr(fmt.Errorf("failed to init templates registry: %w", err))
	}

	r := newRunner(reg)

	if len(os.Args) < 2 {
		closeWithErr(errors.New(defaultMessage))
	}

	switch os.Args[1] {

	case "build":

		if len(os.Args) < 3 {
			closeWithErr(errors.New("missing input YAML file"))
		}

		inputFile := os.Args[2]

		buildCmd := flag.NewFlagSet("build", flag.ExitOnError)

		templateName := buildCmd.String("template", "blueprint", "Template name")
		lang := buildCmd.String("lang", "en", "Language (en, ru)")
		out := buildCmd.String("out", "", "Output PDF file name")

		if err := buildCmd.Parse(os.Args[3:]); err != nil {
			closeWithErr(err)
		}

		if err := r.build(inputFile, *templateName, *lang, *out); err != nil {
			closeWithErr(err)
		}

	case "templates":
		r.listTemplates()

	default:
		closeWithErr(errors.New("expected 'build' or 'templates' subcommands"))
	}
}

func newRunner(reg *registry.Registry) *runner {
	return &runner{
		reg: reg,
	}
}

type runner struct {
	reg *registry.Registry
}

func (r *runner) build(inputFile, templateName, lang, out string) error {
	temp, err := r.reg.GetTemplate(templateName)
	if err != nil {
		return err
	}

	resume, err := parser.LoadResume(inputFile)
	if err != nil {
		return err
	}

	if out == "" {
		out = resume.FileName()
	}

	fmt.Println("Building resume:")
	fmt.Println("  Input:", inputFile)
	fmt.Println("  Template:", templateName)
	fmt.Println("  Output:", out)
	fmt.Println("  Language:", lang)

	if err := validator.Validate(resume); err != nil {
		return err
	}

	if err := renderer.RenderToPDF(temp.Path, resume, lang, out); err != nil {
		return err
	}

	fmt.Println("Success")
	return nil
}

func (r *runner) listTemplates() {
	fmt.Printf("available templates:\n%s\n", r.reg.List())
}

func closeWithErr(err error) {
	fmt.Println(err)
	os.Exit(1)
}
