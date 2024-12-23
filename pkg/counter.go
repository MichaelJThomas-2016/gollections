package pkg

type Counter map[interface{}]int

func (c Counter) Add(k interface{}) {
	if !c.HasKey(k) {
		c[k] = 1
	} else {
		c[k]++
	}
}

func (c Counter) HasKey(key interface{}) bool {
	_, exists := c[key]
	return exists
}

func (c Counter) mostCommon(n int) {

	type Items struct {
		k interface{}
		v interface{}
	}

}
