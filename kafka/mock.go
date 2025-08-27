package kafka

type MockProducer struct {
    Messages []string
}

func (m *MockProducer) Publish(topic, message string) error {
    m.Messages = append(m.Messages, message)
    return nil
}

func (m *MockProducer) Close() error {
    return nil
}
