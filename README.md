Simple cli resume maker which takes a yaml file (see example.yaml) and converts it to pdf.
Right now there is only one template which I use, but it can be extended easily.

project uses go wrapper of wkhtmltopdf, so you'll need to install it on your machine, e.g. (for debian-based distros):
sudo apt install wkhtmltopdf

Build: 
- go build -o resume-maker cmd/main.go

Usage: 
- resume-maker templates: shows available templates
- resume-maker build resume.yaml [-template=template-name] [-lang=en] [-out=pdf-output-file-name.pdf]: creates a pdf resume.
