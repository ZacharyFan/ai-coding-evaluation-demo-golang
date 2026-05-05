package shipping

import "fmt"

type Shipment struct {
	Items     int
	Remote    bool
	Expedited bool
	Zone      string
}

func Rate(shipment Shipment) (int, error) {
	if shipment.Items <= 0 {
		return 0, fmt.Errorf("items must be positive")
	}
	total := 500 + shipment.Items*100
	if shipment.Remote {
		total += shipment.Items * 250
	}
	if shipment.Expedited {
		total += 700
	}
	return total, nil
}
