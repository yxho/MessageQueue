package mq

type Client struct {
	broker *BrokerImpl
}

func NewClient() *Client {
	return &Client{
		broker: NewBroker(),
	}
}

func (c *Client) SetConditions(capacity int) {
	c.broker.setConditions(capacity)
}

func (c *Client) Publish(topic string, msg interface{}) error {
	return c.broker.publish(topic, msg)
}

func (c *Client) Subscribe(topic string) (<-chan interface{}, error) {
	return c.broker.subscribe(topic)
}

func (c *Client) Unsubscribe(topic string, sub <-chan interface{}) error {
	return c.broker.unsubscribe(topic, sub)
}

func (c *Client) Close() {
	c.broker.close()
}

func (c *Client) GetPayLoad(sub <-chan interface{}) interface{} {
	for val := range sub {
		if val != nil {
			return val
		}
	}
	return nil
}
