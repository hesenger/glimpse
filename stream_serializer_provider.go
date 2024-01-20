package glimpse

type StreamSerializerProvider struct {
	serializers map[string]StreamSerializer
}

func NewStreamSerializerProvider() *StreamSerializerProvider {
	return &StreamSerializerProvider{
		serializers: make(map[string]StreamSerializer),
	}
}

func (p *StreamSerializerProvider) Get(streamName string) StreamSerializer {
	return p.serializers[streamName]
}

func (p *StreamSerializerProvider) Register(streamName string, serializer StreamSerializer) {
	p.serializers[streamName] = serializer
}
