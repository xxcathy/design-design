package prototype

type Demo struct {
	L []int             `json:"l"`
	M map[string]string `json:"m"`
	I int               `json:"i"`
	S string            `json:"s"`
}

func (demo *Demo) Clone() *Demo {
	d := &Demo{
		L: make([]int, len(demo.L), len(demo.L)),
		//M: make(map[string]string),
		I: demo.I,
		S: demo.S,
	}

	d.L = make([]int, len(demo.L), len(demo.L))
	copy(d.L, demo.L)

	if demo.M != nil {
		d.M = make(map[string]string, len(demo.M))
		for k, v := range demo.M {
			d.M[k] = v
		}
	}

	return d
}