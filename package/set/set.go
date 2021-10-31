package set

type group struct {
	m map[string]*set
}

type set struct {
	name string
	m    map[string]struct{}
}

func NewGroup() *group {
	return &group{m: map[string]*set{}}
}

func (g *group) Get(name string) *set {
	if g.m[name] == nil {
		g.m[name] = &set{
			name: name,
			m:    make(map[string]struct{}),
		}
	}
	return g.m[name]
}

func (g *group) Del(name string) {
	g.m[name] = nil
}

func (s *set) Add(str string) {
	s.m[str] = struct{}{}
}

func (s *set) Del(str string) {
	delete(s.m, str)
}

func (s *set) IsExist(str string) bool {
	_, ok := s.m[str]
	return ok
}

func (s *set) Len() int {
	return len(s.m)
}
