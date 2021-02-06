package gen

import (
	"bytes"
	"io"
	"os"
	"strings"
	"text/template"

	"golang.org/x/xerrors"
)

// config is input data for templates.
type config struct {
	Layer      int
	Package    string
	Structs    []structDef
	Interfaces []interfaceDef
	Registry   []bindingDef
	Errors     []errCheckDef
}

// FileSystem represents a directory of generated package.
type FileSystem interface {
	WriteFile(baseName string, source []byte) error
}

// outFileName returns file name of generated go source file based on namespace
// and baseName in snake case.
func outFileName(baseName string, namespace []string) string {
	var s strings.Builder
	s.WriteString("tl_")
	for _, ns := range namespace {
		s.WriteString(rules.Underscore(ns))
		s.WriteString("_")
	}
	s.WriteString(rules.Underscore(baseName))
	s.WriteString("_gen.go")
	return s.String()
}

func (g *Generator) hasUpdateClass() bool {
	for _, s := range g.structs {
		if s.Interface == "UpdateClass" {
			return true
		}
	}
	return false
}

type writer struct {
	pkg   string
	fs    FileSystem
	t     *template.Template
	buf   *bytes.Buffer
	wrote map[string]bool

	wroteConstructors map[string]struct{}
}

func (w *writer) Generate(templateName, name string, cfg config) error {
	if cfg.Package == "" {
		cfg.Package = w.pkg
	}
	if w.wrote[name] {
		return xerrors.Errorf("name collision (already wrote %s)", name)
	}

	w.buf.Reset()
	if err := w.t.ExecuteTemplate(w.buf, templateName, cfg); err != nil {
		return xerrors.Errorf("failed to execute template %s for %s: %w", templateName, name, err)
	}
	if err := w.fs.WriteFile(name, w.buf.Bytes()); err != nil {
		_, _ = io.Copy(os.Stderr, w.buf)
		return xerrors.Errorf("failed to write file %s: %w", name, err)
	}
	w.wrote[name] = true

	return nil
}

func (w *writer) WriteInterfaces(interfaces []interfaceDef) error {
	for _, class := range interfaces {
		cfg := config{
			Package:    w.pkg,
			Interfaces: []interfaceDef{class},
			Structs:    class.Constructors,
		}
		for _, s := range cfg.Structs {
			w.wroteConstructors[s.Name] = struct{}{}
		}

		name := outFileName(class.BaseName, class.Namespace)
		if err := w.Generate("main", name, cfg); err != nil {
			return err
		}
	}
	return nil
}

func (w *writer) WriteStructs(structs []structDef) error {
	for _, s := range structs {
		if _, ok := w.wroteConstructors[s.Name]; ok {
			continue
		}
		cfg := config{
			Package: w.pkg,
			Structs: []structDef{s},
		}
		name := outFileName(s.BaseName, s.Namespace)
		if w.wrote[name] {
			// Name collision.
			name = outFileName(s.BaseName+"_const", s.Namespace)
		}
		if err := w.Generate("main", name, cfg); err != nil {
			return err
		}
	}

	return nil
}

// WriteSource writes generated definitions to fs.
func (g *Generator) WriteSource(fs FileSystem, pkgName string, t *template.Template) error {
	w := &writer{
		pkg:   pkgName,
		fs:    fs,
		t:     t,
		buf:   new(bytes.Buffer),
		wrote: map[string]bool{},

		wroteConstructors: map[string]struct{}{},
	}

	if err := w.WriteInterfaces(g.interfaces); err != nil {
		return xerrors.Errorf("interfaces: %w", err)
	}
	if err := w.WriteStructs(g.structs); err != nil {
		return xerrors.Errorf("structs: %w", err)
	}

	if g.hasUpdateClass() {
		if err := w.Generate("handlers", "tl_handlers_gen.go", config{
			Structs: g.structs,
		}); err != nil {
			return err
		}
	}

	cfg := config{
		Registry: g.registry,
		Layer:    g.schema.Layer,
		Errors:   g.errorChecks,
	}

	if err := w.Generate("registry", "tl_registry_gen.go", cfg); err != nil {
		return err
	}
	if err := w.Generate("client", "tl_client_gen.go", cfg); err != nil {
		return err
	}
	if len(cfg.Errors) > 0 {
		if err := w.Generate("errors", "tl_errors_gen.go", cfg); err != nil {
			return err
		}
	}

	return nil
}
