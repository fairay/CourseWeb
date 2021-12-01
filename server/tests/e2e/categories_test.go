package e2e_test

import (
	//"api/recipes/objects"
	"api/recipes/objects"
	"api/recipes/objects/dbuilder"
	"api/recipes/tests"
	"strings"

	//"strconv"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/cookiejar"

	//"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type CategorySuite struct {
	suite.Suite
	Title string
}

/* URL:
/recipes/1/categories POST
/recipes/1/categories GET
*/
func TestPostCategory(t *testing.T) {
	port := tests.StubServer()
	url := fmt.Sprintf("http://localhost:%d/accounts", port)

	jar, _ := cookiejar.New(nil)

	client := &http.Client{
		Jar: jar,
	}
	req, _ := http.NewRequest(http.MethodPost, url, strings.NewReader(`{"login": "admin", "password": "admin"}`))

	res, err := client.Do(req)
	if err != nil {
		t.Error("Error POST account:", err)
		return
	}
	if res.StatusCode != http.StatusOK {
		t.Error("Failed POST account:", res.Status)
		return
	}
	res.Body.Close()

	url = fmt.Sprintf("http://localhost:%d/accounts/login", port)

	req, _ = http.NewRequest(http.MethodPost, url, strings.NewReader(`{"login": "admin", "password": "admin"}`))

	res, err = client.Do(req)
	if err != nil {
		t.Error("Error POST account:", err)
		return
	}
	if res.StatusCode != http.StatusOK {
		t.Error("Failed POST account:", res.Status)
		return
	}
	res.Body.Close()

	url = fmt.Sprintf("http://localhost:%d/recipes", port)

	objRcp := dbuilder.RecipeMother{}.Obj0()

	jsonStr, _ := json.Marshal(objRcp)

	req, _ = http.NewRequest(http.MethodPost, url, strings.NewReader(string(jsonStr)))

	res, err = client.Do(req)
	if err != nil {
		t.Error("Error POST recipe:", err)
		return
	}
	if res.StatusCode != http.StatusOK {
		t.Error("Failed POST recipe:", res.Status)
		return
	}
	res.Body.Close()

	url = fmt.Sprintf("http://localhost:%d/recipes/%d", port, objRcp.Id)

	res, err = http.Get(url)
	if err != nil {
		t.Error("Error GET recipe:", err)
		return
	}
	if res.StatusCode != http.StatusOK {
		t.Error("Failed GET recipe:", res.Status)
		return
	}

	fmt.Println("++++++++++")
	fmt.Println(res.Body)
	rcpDTO := new(objects.RecipeDTO)
	err = json.NewDecoder(res.Body).Decode(rcpDTO)
	if err != nil {
		t.Error("Failed Decode:", err)
		return
	}

	res.Body.Close()

	assert.Equal(t, objRcp, rcpDTO.ToModel())
}
