package main

import (
	"errors"
	pgs "github.com/lyft/protoc-gen-star"
	pgsgo "github.com/lyft/protoc-gen-star/lang/go"
	"text/template"
)

type metricModule struct {
	*pgs.ModuleBase
	ctx pgsgo.Context
	tpl *template.Template
}

func NewMetric() pgs.Module {
	return &metricModule{ModuleBase: &pgs.ModuleBase{}}
}

func (m *metricModule) Name() string {
	return "metric"
}

func (m *metricModule) InitContext(c pgs.BuildContext)  {
	m.ModuleBase.InitContext(c)
	m.ctx = pgsgo.InitContext(c.Parameters())
	tpl := template.New("metric").Funcs(map[string]interface{}{
		"package": m.ctx.PackageName,
		"name": m.ctx.Name,
		"selectSubMsg": m.selectSubMsg,
		"lowerCamelCase": m.lowerCamelCase,
		"lowerSnakeCase": m.lowerSnakeCase,
		"upperCamelCase": m.upperCamelCase,
		"upperSnakeCase": m.upperSnakeCase,
	})
	m.tpl = template.Must(tpl.Parse(metricTemplate))
}

func (m *metricModule) Execute(targets map[string]pgs.File, packages map[string]pgs.Package) []pgs.Artifact {
	for _, f := range targets {
		if len(f.Messages()) == 0 {
			continue
		}
		name := m.ctx.OutputPath(f).SetExt(".mt.go")
		m.AddGeneratorTemplateFile(name.String(), m.tpl, f)
	}
	return m.Artifacts()
}

func (m *metricModule) lowerCamelCase(entity pgs.Entity) string {
	return entity.Name().LowerCamelCase().String()
}

func (m *metricModule) lowerSnakeCase(entity pgs.Entity) string {
	return entity.Name().LowerSnakeCase().String()
}

func (m *metricModule) upperCamelCase(entity pgs.Entity) string {
	return entity.Name().UpperCamelCase().String()
}

func (m *metricModule) upperSnakeCase(entity pgs.Entity) string {
	return entity.Name().UpperSnakeCase().String()
}

func (m *metricModule) selectSubMsg(msg pgs.Message, name string) ([]pgs.Field, error) {
	for _, m := range msg.Messages() {
		if m.Name().LowerCamelCase().String() == name {
			return m.Fields(), nil
		}
	}
	return nil, errors.New("not found message")
}


