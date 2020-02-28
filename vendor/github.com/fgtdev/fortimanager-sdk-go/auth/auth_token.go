package auth

import (
	"fmt"
	"time"
)

const tokenExpiryThreshold = 10
const forceRefreshThreshold = 1

func (a *AuthBlob) GetToken(fmgAddress string, forceLogin bool) (string, error) {
	defer a.LeaveAuthCriticalSection()
	a.EnterAuthCriticalSection()

	current := time.Now()
	alltokens, err := FetchAllTokens()
	needLogin := false
	token := alltokens.TokenPointer(fmgAddress)
	tokenNotPresent := (err != nil) || (token == nil)
	if tokenNotPresent == true {
		newToken := Token{
			Hostname: fmgAddress,
		}
		needLogin = true
		alltokens = append(alltokens, newToken)
		token = alltokens.TokenPointer(fmgAddress)
	}
	// To supress the simutaneous forced login, we set a threshold here
	if needLogin == false && forceLogin == true && current.Sub(token.ForceRefreshTimeStamp).Minutes() >= forceRefreshThreshold {
		needLogin = true
	}
	// examine the timestamp one more time in case the token is expired
	if needLogin == false {
		diff := current.Sub(token.RefreshTimeStamp)
		if diff.Minutes() >= tokenExpiryThreshold {
			needLogin = true
		}
	}
	// relogin in case the fmg service was not ready but, now is able to serve
	if needLogin == false && token.Valid == false {
		needLogin = true
	}

	if needLogin == true {
		session, loginerror := a.loginCallback()
		if loginerror != nil {
			token.MarkInvalid(loginerror)
			token.AppendHistory(fmt.Sprintf("login error:<%v>, token invalidated", loginerror))
		} else {
			token.MarkValid()
			token.SetSession(session)
			if forceLogin == true {
				token.ForceRefreshTimeStamp = current
			}
			token.AppendHistory(fmt.Sprintf("login succeed, new token:%v", session))
		}
	}
	token.RefreshTimestamp()
	token.ProcessEnter()
	PersistTokens(alltokens)

	if token.Valid == false {
		return "", fmt.Errorf("%s", token.ErrorMessage)
	}
	return token.Session, nil
}

func (a *AuthBlob) PutToken(fmgAddress string) error {
	defer a.LeaveAuthCriticalSection()
	a.EnterAuthCriticalSection()

	alltokens, err := FetchAllTokens()
	token := alltokens.TokenPointer(fmgAddress)
	if err != nil || token == nil {
		return fmt.Errorf("token doesn't exist for fmg:%s", fmgAddress)
	}
	token.ProcessLeave()
	if len(token.Counters) == 0 {
		err = a.logoutCallback(token.Session)
		token.MarkInvalid(err)
	}
	PersistTokens(alltokens)
	return nil
}
