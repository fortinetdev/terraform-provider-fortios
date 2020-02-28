package auth

import (
	"fmt"
	"github.com/nightlyone/lockfile"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type AuthBlob struct {
	processGuard   lockfile.Lockfile
	routineGuard   sync.Mutex
	loginCallback  func() (string, error)
	logoutCallback func(string) error
}

const lockFile = "/tmp/.fmg_terraform_auth.lck"
const waitAndPollDuration = 100 // in millisecond

func (authBlob *AuthBlob) TryEnterAuthCriticalSection() error {
	authBlob.routineGuard.Lock()

	err := authBlob.processGuard.TryLock()
	if err != nil {
		authBlob.routineGuard.Unlock()
		return err
	}
	return nil
}

func (authBlob *AuthBlob) LeaveAuthCriticalSection() error {
	defer authBlob.routineGuard.Unlock()

	// The inter-process lock must be owned by current process
	process, err := authBlob.processGuard.GetOwner()
	if err != nil {
		// The lock does't exist.
		return err
	}

	if process.Pid != os.Getpid() {
		return fmt.Errorf("The lock is not owned by current process")
	}

	//authBlob.routineGuard.Unlock()
	return authBlob.processGuard.Unlock()
}

func (authBlob *AuthBlob) EnterAuthCriticalSection() {
	for {
		err := authBlob.TryEnterAuthCriticalSection()
		if err == nil {
			break
		}
		time.Sleep(waitAndPollDuration * time.Millisecond)
	}
}

func FmgAuthInit(loginCallback func() (string, error), logoutCallback func(string) error) *AuthBlob {
	var authBlob AuthBlob
	lockFileAbsPath, _ := filepath.Abs(lockFile)
	processGuard, err := lockfile.New(lockFileAbsPath)
	if err != nil {
		panic(fmt.Sprintf("not able to lock:%s, error:%s", lockFile, err))
	}
	authBlob.processGuard = processGuard
	authBlob.loginCallback = loginCallback
	authBlob.logoutCallback = logoutCallback
	return &authBlob
}
