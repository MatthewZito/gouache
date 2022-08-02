package session

import (
	"testing"
)

func TestRegisterProvider(t *testing.T) {
	providerName := "test"

	p := NewMockProvider()
	RegisterProvider(providerName, p)

	if _, ok := providers[providerName]; !ok {
		t.Errorf("expected to register provider %s but it failed\n", providerName)
	}
}

func TestRegisterProviderDedupe(t *testing.T) {
	providerName := "test"

	p := NewMockProvider()
	RegisterProvider(providerName, p)

	if err := RegisterProvider(providerName, p); err == nil {
		t.Errorf("expected an error when registering the already-registered provider %s but the err value was nil\n", providerName)
	}
}
