package auth

import (
	"fmt"
	"os"
	"strconv"
	"testing"
)

func TestTokenStructure(t *testing.T) {
	var token0 Token
	fooSession := "foo session"
	token0.SetSession(fooSession)
	if token0.Session != fooSession {
		t.Errorf("SetSession fails")
	}

	token0.MarkValid()
	if token0.Valid != true {
		t.Errorf("MarkValid fails")
	}

	fooInvalidMsg := "foo invalid message"
	token0.MarkInvalid(fmt.Errorf("%s", fooInvalidMsg))
	if token0.Valid != false || token0.ErrorMessage != fooInvalidMsg {
		t.Errorf("MarkInvalid fails")
	}

	pid := strconv.Itoa(os.Getpid())
	token0.ProcessEnter()
	if token0.Counters[pid] != 1 {
		t.Errorf("ProcessEnter fails")
	}

	token0.ProcessEnter()
	if token0.Counters[pid] != 2 {
		t.Errorf("ProcessEnter fails")
	}

	token0.ProcessLeave()
	if token0.Counters[pid] != 1 {
		t.Errorf("ProcessLeave fails")
	}
	token0.ProcessLeave()
	if token0.Counters[pid] != 0 {
		t.Errorf("ProcessLeave fails")
	}
}

func TestTokenset(t *testing.T) {
	var tokenSet TokenSet
	token0 := Token{
		Hostname: "192.168.190.102",
	}
	token1 := Token{
		Hostname: "192.168.190.100",
	}
	tokenSet = append(tokenSet, token0)
	tokenSet = append(tokenSet, token1)
	if tokenSet.TokenPointer("192.168.190.101") != nil {
		t.Errorf("TokenPointer fails")
	}

	token := tokenSet.TokenPointer("192.168.190.102")
	if token == nil || token.Hostname != "192.168.190.102" {
		t.Errorf("TokenPointer fails")
	}

	if PersistTokens(tokenSet) != nil {
		t.Errorf("PersistTokens fails")
	}

	allTokens, err := FetchAllTokens()
	if err != nil {
		t.Errorf("FetchAllTokens fails")
	}
	if allTokens.TokenPointer("192.168.190.102") == nil {
		t.Errorf("FetchAllTokens fails")
	}
}
