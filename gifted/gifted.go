package gifted

import "github.com/disintegration/gift"

type Gifted struct {
	*gift.GIFT
}

func (g *Gifted) Remove(filter gift.Filter) bool {
	for i, f := range g.Filters {
		if f == filter {
			copy(g.Filters[i:], g.Filters[i+1:])
			g.Filters[len(g.Filters)-1] = nil
			g.Filters = g.Filters[:len(g.Filters)-1]
			return true
		}
	}
	return false
}

func (g *Gifted) Replace(old, new gift.Filter) bool {
	for i, f := range g.Filters {
		if f == old {
			g.Filters[i] = new
			return true
		}
	}
	g.Add(new)
	return false
}
