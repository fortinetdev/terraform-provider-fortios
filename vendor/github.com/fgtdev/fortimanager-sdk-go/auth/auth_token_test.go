package auth

import (
	"fmt"
	"testing"
	"time"
)

var assignedLoginSession string
var assignedLoginError error

func loginTest() (string, error) {
	return assignedLoginSession, assignedLoginError
}

var assignedLogoutError error

func logoutTest(session string) error {
	return assignedLogoutError
}

func TestToken(t *testing.T) {
	authBlob := FmgAuthInit(loginTest, logoutTest)
	// force an update
	fakeAddress := "FakeFortiManagerAddress"
	assignedLoginSession = "GoodSessionToken0"
	assignedLoginError = nil
	tokenString, err := authBlob.GetToken(fakeAddress, true)
	if err != nil || tokenString != "GoodSessionToken0" {
		t.Errorf("GetToken fails")
	}

	// another following forced GetToken is not actually.
	assignedLoginSession = "GoodSessionToken1"
	tokenString, err = authBlob.GetToken(fakeAddress, true)
	if err != nil || tokenString != "GoodSessionToken0" {
		t.Errorf("GetToken fails")
	}
	// invalidate the token, then next retrieve a new token
	tokenSet, err := FetchAllTokens()
	if err != nil {
		t.Errorf("FetchAllTokens fails")
	}
	token := tokenSet.TokenPointer(fakeAddress)
	if token == nil {
		t.Errorf("TokenPointer fails")
	}
	token.MarkInvalid(fmt.Errorf("invalidated mannually"))
	if PersistTokens(tokenSet) != nil {
		t.Errorf("PersistTokens fails")
	}
	tokenString, err = authBlob.GetToken(fakeAddress, false)
	if err != nil || tokenString != "GoodSessionToken1" {
		t.Errorf("PersistTokens fails")
	}

	// modify the timestamp to be at least 10 minutes earlier than current Now
	tokenSet, err = FetchAllTokens()
	token = tokenSet.TokenPointer(fakeAddress)
	token.RefreshTimeStamp = token.RefreshTimeStamp.Add(-time.Minute * 11)
	PersistTokens(tokenSet)

	assignedLoginSession = "GoodSessionToken2"
	assignedLoginError = fmt.Errorf("dummy error")
	tokenString, err = authBlob.GetToken(fakeAddress, false)
	if err == nil {
		t.Errorf("GetToken fails")
	}

	assignedLoginError = nil
	tokenString, err = authBlob.GetToken(fakeAddress, false)
	if err != nil || tokenString != "GoodSessionToken2" {
		t.Errorf("GetToken fails")
	}

	err = authBlob.PutToken(fakeAddress)
	if err != nil {
		t.Errorf("PutToken fails")
	}
}
