package broker

func truePtr() *bool {
	b := true
	return &b
}

func falsePtr() *bool {
	b := false
	return &b
}

func (b *SpinnakerBroker) ValidateBrokerAPIVersion(version string) error {
	return nil
}
