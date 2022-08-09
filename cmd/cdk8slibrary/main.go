//TODO: add ability to merge yamls
package main

import (
	"log"
	"os"

	ck8slibrary "github.com/coopnorge/cdk8slibrary"
	"github.com/creasty/defaults"
	"github.com/mitchellh/mapstructure"
	"gopkg.in/yaml.v2"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

type MyChartProps struct {
	cdk8s.ChartProps
}

func NewMyChart(scope constructs.Construct, id string, props *ck8slibrary.ChartProps) cdk8s.Chart {
	var cprops cdk8s.ChartProps

	chart := cdk8s.NewChart(scope, jsii.String(id), &cprops)

	ck8slibrary.Chart(chart, jsii.String("hello"), props)

	return chart
}

func main() {
	// Load the file; returns []byte
	f, err := os.ReadFile("test.yaml")
	if err != nil {
		log.Fatal(err)
	}
	// Create an empty Car to be are target of unmarshalling
	var c ck8slibrary.ChartProps
	if err := defaults.Set(&c); err != nil {
		panic(err)
	}

	var raw interface{}

	// Unmarshal our input YAML file into empty interface
	if err := yaml.Unmarshal(f, &raw); err != nil {
		log.Fatal(err)
	}

	// Use mapstructure to convert our interface{} to Car (var c)
	decoder, _ := mapstructure.NewDecoder(&mapstructure.DecoderConfig{WeaklyTypedInput: true, Result: &c})
	if err := decoder.Decode(raw); err != nil {
		log.Fatal(err)
	}

	// Print out the new struct
	//fmt.Printf("%+v\n", *c.Name)
	app := cdk8s.NewApp(nil)
	NewMyChart(app, "test", &c)
	app.Synth()
}
