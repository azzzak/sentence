package sentence

import (
	"bytes"
	"fmt"
	"sync"
	"text/template"
)

type Sentence struct {
	base      *template.Template
	templates map[string]*template.Template
	mu        sync.RWMutex
}

type options struct {
	lang    *string
	delim   *string
	funcs   template.FuncMap
	mapSize *int
}

const (
	pluralName  = "plural"
	pluralnName = "pluraln"
	anyName     = "any"
	anyfName    = "anyf"
)

var funcsNames = map[string]struct{}{
	pluralName:  {},
	pluralnName: {},
	anyName:     {},
	anyfName:    {},
}

type Option func(options *options) error

const (
	LangEnglish string = "en"
	LangRussian string = "ru"
)

var langsName = map[string]struct{}{
	LangEnglish: {},
	LangRussian: {},
}

var formsNum = map[string]int{
	LangEnglish: 2,
	LangRussian: 3,
}

// WithLang задает язык для согласования с числом. Значение по умолчанию "ru".
func WithLang(lang string) Option {
	return func(opts *options) error {
		if lang == "" {
			return fmt.Errorf("language can't be empty")
		}

		if _, ok := langsName[lang]; !ok {
			return fmt.Errorf("language '%s' is not supported", lang)
		}

		opts.lang = &lang
		return nil
	}
}

// WithDelim задает разделитель элементов в строке. Значение по умолчанию "|".
func WithDelim(delim string) Option {
	return func(opts *options) error {
		if delim == "" {
			return fmt.Errorf("delimiter can't be empty")
		}

		opts.delim = &delim
		return nil
	}
}

// WithFunc передает собственную функцию, которая может быть использована в шаблонах.
func WithFunc(name string, fn interface{}) Option {
	return func(opts *options) error {
		if _, ok := funcsNames[name]; ok {
			return fmt.Errorf("function name '%s' is already used", name)
		}

		opts.funcs[name] = fn
		return nil
	}
}

// WithMapSize задает первоначальный размер мапы для кэша. Значение по умолчанию 10.
func WithMapSize(size int) Option {
	return func(opts *options) error {
		if size < 0 {
			return fmt.Errorf("map size can't be negative")
		}

		opts.mapSize = &size
		return nil
	}
}

// New создает новый объект. Использовать безопасно в разных горутинах.
func New(opts ...Option) (*Sentence, error) {
	var options options
	options.funcs = make(template.FuncMap)

	for _, opt := range opts {
		err := opt(&options)
		if err != nil {
			return nil, err
		}
	}

	lang := "ru"
	if options.lang != nil {
		lang = *options.lang
	}

	delim := "|"
	if options.delim != nil {
		delim = *options.delim
	}

	mapSize := 10
	if options.mapSize != nil {
		mapSize = *options.mapSize
	}

	pluralFn := initPlural(lang, delim)

	funcs := template.FuncMap{
		pluralName:  pluralFn,
		pluralnName: initPluraln(pluralFn),
		anyName:     initAny(delim),
		anyfName:    initAnyf(),
	}

	if len(options.funcs) > 0 {
		for k, v := range options.funcs {
			funcs[k] = v
		}
	}

	return &Sentence{
		base:      template.New("base").Funcs(funcs),
		templates: make(map[string]*template.Template, mapSize),
	}, nil
}

// Render производит насыщение шаблона данными.
func (s *Sentence) Render(pattern string, data interface{}) (string, error) {
	t := s.getTemplate(pattern)

	if t == nil {
		var err error
		t, err = s.parse(pattern)
		if err != nil {
			return "", err
		}

		go s.setTemplate(pattern, t)
	}

	output, err := s.execute(t, data)
	if err != nil {
		return "", err
	}

	return output, nil
}

// MustRender производит насыщение шаблона данными. В случае ошибки вызывается panic.
func (s *Sentence) MustRender(pattern string, data interface{}) string {
	res, err := s.Render(pattern, data)
	if err != nil {
		panic(err)
	}

	return res
}

// Prepare служит для предварительного заполнения кэша данными.
func (s *Sentence) Prepare(patterns ...string) error {
	for _, pattern := range patterns {
		t, err := s.parse(pattern)
		if err != nil {
			return err
		}

		s.setTemplate(pattern, t)
	}

	return nil
}

func (s *Sentence) parse(pattern string) (*template.Template, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	t, err := s.base.Parse(pattern)
	if err != nil {
		return nil, err
	}

	return t, nil
}

func (s *Sentence) execute(t *template.Template, data interface{}) (string, error) {
	var buf bytes.Buffer

	s.mu.RLock()
	defer s.mu.RUnlock()

	if err := t.Execute(&buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}

func (s *Sentence) getTemplate(pattern string) *template.Template {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if t, ok := s.templates[pattern]; ok {
		return t
	}

	return nil
}

func (s *Sentence) setTemplate(pattern string, t *template.Template) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.templates[pattern] = t
}
