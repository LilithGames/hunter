package main

import (
	"errors"
	pgs "github.com/lyft/protoc-gen-star"
	pgsgo "github.com/lyft/protoc-gen-star/lang/go"
	"text/template"
)

type seriesModule struct {
	*pgs.ModuleBase
	ctx pgsgo.Context
	tpl *template.Template
}

func newSeries() pgs.Module {
	return &seriesModule{ModuleBase: &pgs.ModuleBase{}}
}

func (m *seriesModule) Name() string {
	return "series"
}

func (m *seriesModule) InitContext(c pgs.BuildContext)  {
	m.ModuleBase.InitContext(c)
	m.ctx = pgsgo.InitContext(c.Parameters())
	tpl := template.New("series").Funcs(map[string]interface{}{
		"package": m.ctx.PackageName,
		"name": m.ctx.Name,
		"selectSubMsg": m.selectSubMsg,
		"lowerCamelCase": m.lowerCamelCase,
		"lowerSnakeCase": m.lowerSnakeCase,
		"upperCamelCase": m.upperCamelCase,
		"upperSnakeCase": m.upperSnakeCase,
	})
	m.tpl = template.Must(tpl.Parse(seriesTemplate))
}

func (m *seriesModule) Execute(targets map[string]pgs.File, packages map[string]pgs.Package) []pgs.Artifact {
	for _, f := range targets {
		if len(f.Messages()) == 0 {
			continue
		}

		needGenerate := false
		for _, m := range f.AllMessages() {
			for _, mm := range m.AllMessages(){
				if mm.Name() == "Tag" || mm.Name() == "Field" {
					needGenerate = true
				}
			}
		}

		if !needGenerate {
			continue
		}
		
		name := m.ctx.OutputPath(f).SetExt(".sr.go")
		m.AddGeneratorTemplateFile(name.String(), m.tpl, f)
	}
	return m.Artifacts()
}

func (m *seriesModule) lowerCamelCase(entity pgs.Entity) string {
	return entity.Name().LowerCamelCase().String()
}

func (m *seriesModule) lowerSnakeCase(entity pgs.Entity) string {
	return entity.Name().LowerSnakeCase().String()
}

func (m *seriesModule) upperCamelCase(entity pgs.Entity) string {
	return entity.Name().UpperCamelCase().String()
}

func (m *seriesModule) upperSnakeCase(entity pgs.Entity) string {
	return entity.Name().UpperSnakeCase().String()
}

func (m *seriesModule) selectSubMsg(msg pgs.Message, name string) ([]pgs.Field, error) {
	for _, m := range msg.Messages() {
		if m.Name().LowerCamelCase().String() == name {
			return m.Fields(), nil
		}
	}
	return nil, errors.New("not found message")
}


