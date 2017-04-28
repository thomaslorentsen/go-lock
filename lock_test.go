package lock

import (
    "testing"
)

func TestInitLock(t *testing.T) {
    l := NewLock()
    if (l.Spoons != 0) {
        t.Errorf("%v Spoons not set when creating new lock", l)
    }
}

func TestUnlockedOnInit(t *testing.T) {
    l := NewLock()
    if (l.Locked()) {
        t.Error("Locked even though lock has not been aquired yet")
    }
}

func TestLock(t *testing.T) {
    l := NewLock()
    if (!l.Lock()) {
        t.Error("Not locked when locking")
    }
}

func TestIsLocked(t *testing.T) {
    l := NewLock()
    l.Lock()
    if (!l.Locked()) {
        t.Error("Not locked when checking if locked")
    }
}

func TestUnlockWithoutLock(t *testing.T) {
    l := NewLock()
    l.Unlock()
    if (l.Locked()) {
        t.Error("Became locked after unlocking")
    }
}

func TestNotLockedIfUnlockedFirst(t *testing.T) {
    l := NewLock()
    l.Unlock()
    if (!l.Lock()) {
        t.Error("Lock not aquired when unlock happened first")
    }
}

func TestStaysLockedWithMultipleLocks(t *testing.T) {
    l := NewLock()
    l.Lock()
    l.Lock()
    l.Unlock()
    if (!l.Locked()) {
        t.Error("Is not locked when locked multiple times")
    }
}
