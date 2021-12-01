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

func PostQuery(client *http.Client, url string, body string) error {
	req, _ := http.NewRequest(http.MethodPost, url, strings.NewReader(body))

	res, err := client.Do(req)
	if err != nil { return err }
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		err = errors.New(fmt.Sprintf("Unexpected status %v", res.StatusCode))
	}
	return err
}

func GetQuery(client *http.Client, url string, obj interface{}) error {
	res, err := client.Get(url)
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

	jar, _ := cookiejar.New(nil)
	client := &http.Client{
		Jar: jar,
	}

	baseUrl := fmt.Sprintf("http://localhost:%d", port)

	acc := dbuilder.AccountMother{}.Obj0().ToDTO()
	acc.Password = acc.Login
	accByte, _ := json.Marshal(acc)
	accStr := string(accByte)

	/// REQUESTS
	// Create account
	url := fmt.Sprintf("%s/accounts", baseUrl)
	err := PostQuery(client, url, accStr)
	if err != nil {
		t.Error("POST account failed:", err)
		return
	}

	// Log in
	url = fmt.Sprintf("%s/accounts/login", baseUrl)
	err = PostQuery(client, url, accStr)
	if err != nil {
		t.Error("Login account failed:", err)
		return
	}

	// Create recipe
	url = fmt.Sprintf("%s/recipes", baseUrl)
	objRcp := dbuilder.RecipeMother{}.Obj0()
	jsonStr, _ := json.Marshal(objRcp)
	err = PostQuery(client, url, string(jsonStr))
	if err != nil {
		t.Error("POST recipe failed:", err)
		return
	}

	// Get recipe
	url = fmt.Sprintf("%s/recipes/%d", baseUrl, objRcp.Id)
	rcpDTO := new(objects.RecipeDTO)
	err = GetQuery(client, url, rcpDTO)
	if err != nil {
		t.Error("GET recipe failed:", err)
		return
	}
	tests.CompareRecipe(t, *objRcp, *rcpDTO.ToModel())

}
