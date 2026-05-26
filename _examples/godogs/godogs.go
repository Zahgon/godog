package godogs

// Godogs is an example behavior holder.
type Godogs int

// Add increments Godogs count.
func (g *Godogs) Add(n int) { _ = "STUB: not implemented"; return }

// Eat decrements Godogs count or fails if there is not enough available.
func (g *Godogs) Eat(n int) error { _ = "STUB: not implemented"; return nil }

// Available returns the number of currently available Godogs.
func (g *Godogs) Available() int { _ = "STUB: not implemented"; return 0 }
