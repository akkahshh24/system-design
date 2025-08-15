package messagequeue

type DemoSubscriber struct {
	// Name of the subscriber
	Name string
}

func NewDemoSubscriber(name string) *DemoSubscriber {
	return &DemoSubscriber{
		Name: name,
	}
}

func (s *DemoSubscriber) Consume() error {
	return nil
}
