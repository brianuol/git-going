package main

import (
	"context"
	"testing"
)

func TestCreateAnonymousClient(t *testing.T) {
	client, throttleMs := CreateClient(context.Background(), "")
	if client == nil {
		t.Error("Client not instantiated")
	}
	if throttleMs != anonDelay {
		t.Error("Expected anonymous throttling policy for anonymous github client")
	}
}

func TestCreateAuthTokenClient(t *testing.T) {
	client, throttleMs := CreateClient(context.Background(), "aToken!")
	if client == nil {
		t.Error("Client not instantiated")
	}
	if throttleMs != authDelay {
		t.Error("Expected authenticated throttling policy for oauth github client")
	}
}
