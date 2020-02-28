package auth

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"time"
)

type Token struct {
	Hostname              string
	Session               string
	RefreshTimeStamp      time.Time
	ForceRefreshTimeStamp time.Time
	Counters              map[string]int

	Valid        bool
	ErrorMessage string
	History      []string
}

type TokenSet []Token

const persistencyFile = "/tmp/.fmg_terraform_tokens"

func PersistTokens(tokens TokenSet) error {
	jsonString, err := json.Marshal(tokens)
	if err != nil {
		return err
	}

	tokenFile, err := os.Create(persistencyFile)
	if err != nil {
		return err
	}
	_, err = tokenFile.Write(jsonString)
	if err != nil {
		return err
	}

	tokenFile.Close()
	return nil
}

func FetchAllTokens() (TokenSet, error) {
	data, err := ioutil.ReadFile(persistencyFile)
	if err != nil {
		return nil, err
	}
	var tokens TokenSet
	err = json.Unmarshal(data, &tokens)
	return tokens, err
}

func (tokenSet TokenSet) TokenPointer(hostname string) *Token {
	for idx := 0; idx < len(tokenSet); idx++ {
		if tokenSet[idx].Hostname == hostname {
			return &tokenSet[idx]
		}
	}
	return nil
}
func (token *Token) MarkValid() {
	token.Valid = true
	token.ErrorMessage = ""
}

func (token *Token) MarkInvalid(errormessage error) {
	token.Valid = false
	token.ErrorMessage = fmt.Sprintf("%v", errormessage)
}

func (token *Token) SetSession(session string) {
	token.Session = session
}

func (token *Token) RefreshTimestamp() {
	token.RefreshTimeStamp = time.Now()
}

func (token *Token) ProcessEnter() {
	if token.Counters == nil {
		token.Counters = make(map[string]int)
	}
	token.Counters[strconv.Itoa(os.Getpid())]++
}

func (token *Token) ProcessLeave() {
	if token.Counters == nil {
		token.Counters = make(map[string]int)
	}
	pid := strconv.Itoa(os.Getpid())
	token.Counters[pid]--
	if token.Counters[pid] <= 0 {
		delete(token.Counters, pid)
	}
}

func (token *Token) AppendHistory(msg string) {
	fmtMsg := fmt.Sprintf("%s: %s", time.Now(), msg)
	token.History = append(token.History, fmtMsg)
}
