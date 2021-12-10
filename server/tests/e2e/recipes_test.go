package e2e_test

import (
	"api/recipes/objects"
	"api/recipes/objects/dbuilder"
	. "api/recipes/tests"
	"time"

	"encoding/json"
	"fmt"
	"testing"
)

func postRecipe(t *testing.T, client *ClientE2E, baseUrl string) {
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
	CompareRecipe(t, *objRcp, *rcpDTO.ToModel())
}

func TestPostRecipe(t *testing.T) {
	port := StubServer()
	baseUrl := fmt.Sprintf("http://localhost:%d", port)

	client := NewClient()
	postRecipe(t, client, baseUrl)	
}

func BenchmarkPostRecipe(b *testing.B) {
	port := StubServer()
	baseUrl := fmt.Sprintf("http://localhost:%d", port)
	t := new(testing.T)
	
	for i := 0; i < b.N; i++ {
		client := NewClient()
		postRecipe(t, client, baseUrl)
	}
}

func RunStubServer(t *testing.T) {
	StubServer()
	time.Sleep(1e9 * 30)
	fmt.Println("1")
}
