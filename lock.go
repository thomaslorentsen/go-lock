package lock

type Lock struct {
    Spoons int
}

func NewLock() *Lock {
    return &Lock{Spoons: 0}
}

func (l *Lock) Lock() bool {
    l.Spoons = l.Spoons + 1
    return true
}

func (l *Lock) Unlock() bool {
    if (l.Spoons == 0) {
        return l.Locked()
    }
    l.Spoons = l.Spoons - 1
    return l.Locked()
}

func (l *Lock) Locked() bool {
    return l.Spoons > 0
}
