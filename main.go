package main

import (
	"bytes"
	"context"
	"text/template"

	"github.com/mitchellh/mapstructure"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

const cssTemplate = `
{{if .Color}}
.mb-header{
	background:{{.Color}};
}
{{end}}

{{if .Logo}}
.mb-logo{
	content:url({{.Logo}});
	padding: .4em !important;
	display: flex !important;
}
{{end}}

{{if .Icon}}
input.mb-query{
	height: 30px;
	background-image: url({{.Icon}});
	background-size: contain;
	background-repeat: no-repeat;
	background-position: -4px 50%;
}
{{end}}
`

type cssParams struct {
	Logo string
	Icon string
	Color string
}

var t *template.Template

func init() {
	t = template.Must(template.New("css").Parse(cssTemplate))
}

func main() {
	lambda.Start(handler)
}

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	params := cssParams{}
	css := &bytes.Buffer{}

	mapstructure.Decode(request.QueryStringParameters, &params)
	t.Execute(css, params)

	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers: map[string]string{"Content-type": "text/css"},
		Body:       css.String(),
	}, nil
}