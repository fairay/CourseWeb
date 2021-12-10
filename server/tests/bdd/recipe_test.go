package bdd_test

import (
	"encoding/json"
	"fmt"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"api/recipes/objects"
	"api/recipes/objects/dbuilder"
	. "api/recipes/tests"
)

var _ = Describe("Recipe", func() {
	port := StubServer()
	client := NewClient()

	baseUrl := fmt.Sprintf("http://localhost:%d", port)

	acc := dbuilder.AccountMother{}.Obj0().ToDTO()
	acc.Password = acc.Login
	accByte, _ := json.Marshal(acc)
	accStr := string(accByte)
	var objRcp *objects.Recipe;

	Describe("Creating new account", func() {
		Context("With user's role", func() {
			It("Should add successfuly", func() {
				url := fmt.Sprintf("%s/accounts", baseUrl)
				err := client.PostQuery(url, accStr)
				Expect(err).To(BeNil())
			})
		})
	})

	Describe("Logging in new account", func() {
		Context("With correct data", func() {
			It("Shoud add token to cookie", func() {
				url := fmt.Sprintf("%s/accounts/login", baseUrl)
				err := client.PostQuery(url, accStr)
				Expect(err).To(BeNil())
			})
		})
	})

	Describe("Creating new recipe", func() {
		Context("With correct data, author exists", func() {
			It("Shoud add successfuly", func() {
				url := fmt.Sprintf("%s/recipes", baseUrl)
				objRcp = dbuilder.RecipeMother{}.Obj0()
				jsonStr, _ := json.Marshal(objRcp)
				err := client.PostQuery(url, string(jsonStr))
				Expect(err).To(BeNil())
			})
		})
	})

	Describe("Get recipe", func() {
		Context("With existing id", func() {
			It("Shoud get expected recipe", func() {
				url := fmt.Sprintf("%s/recipes/%d", baseUrl, objRcp.Id)
				rcpDTO := new(objects.RecipeDTO)

				err := client.GetQuery(url, rcpDTO)
				rcpModel := rcpDTO.ToModel()
				objRcp.CreatedAt = time.Time{}
				rcpModel.CreatedAt = time.Time{}

				Expect(err).To(BeNil())
				Expect(objRcp).To(Equal(rcpModel))				
			})
		})
	})
})
