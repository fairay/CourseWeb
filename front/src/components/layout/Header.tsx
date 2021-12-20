import { Box, Container, Divider } from "@chakra-ui/react";
import React from "react";
import AllRecipes from "./home/allrecipes/AllRecipes";
import {BrowserRouter, Routes, Route, useParams, RouteProps, Params} from "react-router-dom"
import RecipeParams from "./home/recipe/Recipe"
import Header from "./header/Header"
import Input from "components/base/Input"
import Author from "components/base/Author";

const searchInput = Input
const authorField = Author

const HeaderRouter: React.FC<{}> = ({}) => {
    return <BrowserRouter>
        <Routes>
            <Route path="/" element={<Header title="Все рецепты" addField={searchInput} />}/>
            <Route path="/me/likes" element={<Header subtitle="Понравилось" title="Мне" />}/>
            <Route path="/me/recipes" element={<Header subtitle="Автор" title="Я" />}/>

            <Route path="/auth/signin" element={<Header title="Вход" undertitle="Добро пожаловать. Снова." />}/>
            <Route path="/auth/signup" element={<Header title="Регистрация" undertitle="Чтобы получить доступ к тысячам новых возможностей в Вашем кулинарном самовыражении!" />}/>
            
            <Route path="/recipes/:id" element={<Header title="" addField={authorField}/>}/>

            <Route path="*" element={<Header title="Страница не найдена"/>}/>
        </Routes>
    </BrowserRouter>
}

export default HeaderRouter