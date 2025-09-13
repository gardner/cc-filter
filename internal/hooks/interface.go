package hooks

type HookProcessor interface {
	CanHandle(input map[string]interface{}) bool
	Process(input map[string]interface{}) (string, error)
}

type Registry struct {
	processors []HookProcessor
}

func NewRegistry() *Registry {
	return &Registry{
		processors: make([]HookProcessor, 0),
	}
}

func (r *Registry) Register(processor HookProcessor) {
	r.processors = append(r.processors, processor)
}

func (r *Registry) Process(input map[string]interface{}) (string, bool) {
	for _, processor := range r.processors {
		if processor.CanHandle(input) {
			if result, err := processor.Process(input); err == nil {
				return result, true
			}
		}
	}
	return "", false
}