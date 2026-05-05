package shipping

import "testing"

func TestRateLocalShipment(t *testing.T) {
	got, err := Rate(Shipment{Items: 2})
	if err != nil {
		t.Fatalf("Rate returned error: %v", err)
	}
	if got != 700 {
		t.Fatalf("Rate = %d, want 700", got)
	}
}

func TestRateExpeditedShipment(t *testing.T) {
	got, err := Rate(Shipment{Items: 1, Expedited: true})
	if err != nil {
		t.Fatalf("Rate returned error: %v", err)
	}
	if got != 1300 {
		t.Fatalf("Rate = %d, want 1300", got)
	}
}
