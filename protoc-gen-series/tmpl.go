package main

const seriesTemplate = `// Code generated by protoc-gen-series. DO NOT EDIT.

package {{ package . }}

import (
	influxdb "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
	"time"
)
{{ range .Messages }}{{- $project := name . }}{{ $lowerCamelProject :=  lowerCamelCase . }}{{ $lowerSnakeProject := lowerSnakeCase . }}
type {{ $project }}Writer interface {
	Flush()
	Errors() <- chan error
	{{- range .Messages -}}{{ $table := upperCamelCase . }}
	Write{{ $table }}Point(tag *{{ $project }}{{ $table }}Tag, field *{{ $project }}{{ $table }}Field)
	Write{{ $table }}PointWithTime(tag *{{ $project }}{{ $table }}Tag, field *{{ $project }}{{ $table }}Field, timestamp time.Time)
	{{- end }}
}

type {{ lowerCamelCase . }}Writer struct {
	api.WriteAPI
}

func New{{ $project }}Writer(endpoint string) {{ $project }}Writer {
	return &{{ lowerCamelCase . }}Writer{WriteAPI: influxdb.NewClient(endpoint, "").WriteAPI("", "")}
}

func (w *{{ lowerCamelCase . }}Writer) Flush() {
	w.WriteAPI.Flush()
}

func (w *{{ lowerCamelCase . }}Writer) Errors() <-chan error {
	return w.WriteAPI.Errors()
}{{ range .Messages }}{{ $table := upperCamelCase . }}{{ $lowerSnakeTable := lowerSnakeCase . }}

type {{ $project }}{{ $table }}Tag struct {
	{{ $project }}_{{ $table }}_Tag
}

type {{ $project }}{{ $table }}Field struct {
	{{ $project }}_{{ $table }}_Field
}

func (w *{{ $lowerCamelProject }}Writer) Write{{ $table }}Point(tag *{{ $project }}{{ $table }}Tag, field *{{ $project }}{{ $table }}Field) {
	w.Write{{ $table }}PointWithTime(tag, field, time.Now())
}

func (w *{{ $lowerCamelProject }}Writer) Write{{ $table }}PointWithTime(tag *{{ $project }}{{ $table }}Tag, field *{{ $project }}{{ $table }}Field, timestamp time.Time) {
	point := influxdb.NewPointWithMeasurement("{{ $lowerSnakeProject }}_{{ $lowerSnakeTable }}"){{ range ( selectSubMsg . "tag" )}}
	point.AddTag("{{ lowerSnakeCase . }}", tag.{{ upperCamelCase . }}){{ end }}{{ range ( selectSubMsg . "field")}}
	point.AddField("{{ lowerSnakeCase . }}", field.{{ upperCamelCase . }}){{ end }}
	w.WriteAPI.WritePoint(point)
}
{{- end -}}
{{ end }}
`
