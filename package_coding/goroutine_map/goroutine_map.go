package goroutine_map

type GMap struct {
	m map[string]string
}

func NewGMap(m map[string]string) *GMap {
	return &GMap{
		m: m,
	}
}

/*
写一个并发安全的字典(map)类型
要求如下：
1.写成一个包
2.这个字典类型包含添加、获取、删除、大小功能
3.有单元测试，覆盖率达到100%
*/
func (g *GMap) Add(key, val string) bool {
	_, ok := g.m[key]
	if !ok {
		g.m[key] = val
		return true
	} else {
		return false
	}
}

func (g *GMap) Get(key string) string {
	return g.m[key]
}

func (g *GMap) Del(key string) {
	delete(g.m, key)
}

func (g *GMap) Len() int {
	return len(g.m)
}
