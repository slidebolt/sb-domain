package domain

// LockLock commands a lock to lock.
type LockLock struct{}

func (LockLock) ActionName() string { return "lock_lock" }

// LockUnlock commands a lock to unlock.
type LockUnlock struct {
	Code *string `json:"code,omitempty"`
}

func (LockUnlock) ActionName() string { return "lock_unlock" }
