# Resume Maker

Simple CLI resume generator that takes a YAML file (see `example.yaml`) and converts it to a PDF.

Currently, there is one template available, but the system is designed to be easily extended with additional templates.

---

## Requirements

This project uses a Go wrapper around `wkhtmltopdf`, so you must install `wkhtmltopdf` on your machine.

For Debian-based distributions:

```bash
sudo apt install wkhtmltopdf
```

---

## Build

```bash
go build -o resume-maker cmd/main.go
```

---

## Usage

### Show Available Templates

```bash
resume-maker templates
```

### Build Resume

```bash
resume-maker build resume.yaml [-template=template-name] [-lang=en] [-out=output-file.pdf]
```
