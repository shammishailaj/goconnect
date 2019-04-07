package goconnect_test

import (
	"context"
	"fmt"
	"github.com/autom8ter/goconnect"
	"github.com/autom8ter/objectify"
	"io/ioutil"
	"os"
	"testing"
)

var ctx = context.Background()
var tool = objectify.New()

type TestDoc struct {
	Content map[string]interface{}
	Results map[string]interface{}
}

var Doc = &TestDoc{
	Content: map[string]interface{}{
		"translate": []string{"Hello World!"},
	},
}

func Test(t *testing.T) {
	var err error
	g, err := goconnect.New(ctx, &goconnect.Config{
		ProjectID:       os.Getenv("PROJECT_ID"),
		JSONPath:        "credentials.json",
		TwilioAccount:   os.Getenv("TWILIO_ACCOUNT"),
		TwilioToken:     os.Getenv("TWILIO_TOKEN"),
		SendGridAccount: os.Getenv("SENDGRID_ACCOUNT"),
		SendGridToken:   os.Getenv("SENDGRID_TOKEN"),
		StripeAccount:   os.Getenv("STRIPE_ACCOUNT"),
		StripeToken:     os.Getenv("STRIPE_TOKEN"),
		SlackAccount:    os.Getenv("SLACK_ACCOUNT"),
		SlackToken:      os.Getenv("SLACK_TOKEN"),
		Scopes:          []string{"users"},
		InCluster:       false,
		MasterKey:       os.Getenv("PROJECT_ID"),
	})
	if err != nil {
		t.Fatal(err.Error())
	}
	if g == nil {
		t.Fatal("nil goconnect")
	}
	if g.GCP().Clients() == nil {
		t.Fatal("nil gcp clients")
	}
	if g.GCP().Services() == nil {
		t.Fatal("nil gcp services")
	}
	f := g.GCP().Clients().FireStore
	doc := f.Collection("test").Doc(tool.UUID())

	bits, err := ioutil.ReadFile("credentials.json")
	if err != nil {
		t.Fatal(err.Error())
	}
	Doc.Content["credentials"] = string(bits)

	if err != nil {
		t.Fatal(err.Error())
	}

	resp, err := doc.Create(ctx, tool.ToMap(Doc))
	if err != nil {
		t.Fatal(err.Error())
	}
	fmt.Println("Firestore Document Update Time:", tool.HumanizeTime(resp.UpdateTime), resp.UpdateTime.String())
}
