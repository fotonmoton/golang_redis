package warehouse

type Notifier interface {
	NotifyCustomers([]CustomerNotification)
}

type Subscriptions interface {
	FindSubscriptions(string) []CustomerSubscription
}

type CustomerNotifier struct {
	subscriptions Subscriptions
	notifier      Notifier
}

func NewCustomerNotifier(s Subscriptions, n Notifier) *CustomerNotifier {
	return &CustomerNotifier{s, n}
}

func (c *CustomerNotifier) Notify(product Product) {

	subs := c.subscriptions.FindSubscriptions(product.Name)

	notifications := []CustomerNotification{}

	for _, subscription := range subs {
		notifications = append(notifications, CustomerNotification{subscription.CustomerId, product.ID})
	}

	if len(notifications) != 0 {
		c.notifier.NotifyCustomers(notifications)
	}
}

func (c *CustomerNotifier) Observe(subject any) {
	c.Notify(subject.(Product))
}
