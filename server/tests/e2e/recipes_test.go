package e2e_test

import (
	"api/recipes/objects"
	"api/recipes/objects/dbuilder"
	"api/recipes/tests"

	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"strings"
	"testing"
)

type ClientE2E struct {
	client *http.Client;
}

func NewClient() *ClientE2E {
	client := new(ClientE2E)
	jar, _ := cookiejar.New(nil) 
	client.client = &http.Client {
		Jar: jar,
	}
	return client
}

func (this *ClientE2E) PostQuery(url string, body string) error {
	req, _ := http.NewRequest(http.MethodPost, url, strings.NewReader(body))

	res, err := this.client.Do(req)
	if err != nil { return err }
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		err = errors.New(fmt.Sprintf("Unexpected status %v", res.StatusCode))
	}
	return err
}

func (this *ClientE2E) GetQuery(url string, obj interface{}) error {
	res, err := this.client.Get(url)
	if err != nil { return err }
	defer res.Body.Close()
	
	if res.StatusCode != http.StatusOK {
		err = errors.New(fmt.Sprintf("Unexpected status %v", res.StatusCode))
		return err
	}

	return json.NewDecoder(res.Body).Decode(obj)
}

func TestPostRecipe(t *testing.T) {
	port := tests.StubServer()
	client := NewClient()

	baseUrl := fmt.Sprintf("http://localhost:%d", port)

	acc := dbuilder.AccountMother{}.Obj0().ToDTO()
	acc.Password = acc.Login
	accByte, _ := json.Marshal(acc)
	accStr := string(accByte)

	/// REQUESTS
	// Create account
	url := fmt.Sprintf("%s/accounts", baseUrl)
	err := client.PostQuery(url, accStr)
	if err != nil {
		t.Error("POST account failed:", err)
		return
	}

	// Log in
	url = fmt.Sprintf("%s/accounts/login", baseUrl)
	err = client.PostQuery(url, accStr)
	if err != nil {
		t.Error("Login account failed:", err)
		return
	}

	// Create recipe
	url = fmt.Sprintf("%s/recipes", baseUrl)
	objRcp := dbuilder.RecipeMother{}.Obj0()
	jsonStr, _ := json.Marshal(objRcp)
	err = client.PostQuery(url, string(jsonStr))
	if err != nil {
		t.Error("POST recipe failed:", err)
		return
	}

	// Get recipe
	url = fmt.Sprintf("%s/recipes/%d", baseUrl, objRcp.Id)
	rcpDTO := new(objects.RecipeDTO)
	err = client.GetQuery(url, rcpDTO)
	if err != nil {
		t.Error("GET recipe failed:", err)
		return
	}
	tests.CompareRecipe(t, *objRcp, *rcpDTO.ToModel())

}
