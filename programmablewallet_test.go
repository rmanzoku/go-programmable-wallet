package programmablewallet

import (
	"context"
	"log"
	"os"
	"testing"
)

var (
	p *ProgrammableWallet

	apikey = os.Getenv("APIKEY")
	ctx    = context.TODO()
)

func TestGetConfigurationForEntity(t *testing.T) {
	appId, err := p.GetConfigurationForEntity(ctx)
	if err != nil {
		t.Fatalf("failed to get configuration for entity: %v", err)
	}

	t.Log(appId)
}

func TestMain(m *testing.M) {
	var err error
	p, err = NewProgrammableWallet(apikey)
	if err != nil {
		log.Fatalf("failed to create ProgrammableWallet: %v", err)
	}

	os.Exit(m.Run())
}
