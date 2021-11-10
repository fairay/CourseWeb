package dbuilder

import (
	"api/recipes/objects"
	"time"
)

type RecipeBuilder struct {
	Id          int				
	Author      string			
	Title       string			
	CreatedAt   time.Time		
	Description string     		
	Duration    int 			
	PortionNum  int
}
func newRecipeBuilder() *RecipeBuilder { 
	return new(RecipeBuilder) 
}
func (this *RecipeBuilder) Build() *objects.Recipe { 
	return &objects.Recipe{Id: this.Id, Author: this.Author, 
		Title: this.Title, CreatedAt: this.CreatedAt,
		Description: this.Description, Duration: this.Duration,
		PortionNum: this.PortionNum} 
}

type RecipeMother struct{}
func (RecipeMother) Obj0() *objects.Recipe {
	b := newRecipeBuilder()
	b.Id = 1
	b.Author = "test1"
	b.Title = "Блины"
	b.Description = "Низкокалорийные"
	b.Duration = 1
	b.PortionNum = 3

	return b.Build()
}
func (RecipeMother) Obj1() *objects.Recipe {
	b := newRecipeBuilder()
	b.Id = 2
	b.Author = "test1"
	b.Title = "Сырники"
	b.Description = "На завтрак"
	b.Duration = 1
	b.PortionNum = 2

	return b.Build()
}
func (RecipeMother) Obj2() *objects.Recipe {
	b := newRecipeBuilder()
	b.Id = 3
	b.Author = "test2"
	b.Title = "Оладьи"
	b.Description = "Из остатков кефира"
	b.Duration = 25
	b.PortionNum = 7

	return b.Build()
}

