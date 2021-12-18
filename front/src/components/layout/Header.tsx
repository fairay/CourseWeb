import { Box, Container, Divider } from "@chakra-ui/react";
import React from "react";
import AllRecipes from "./home/allrecipes/AllRecipes";
import {BrowserRouter, Routes, Route, useParams, RouteProps, Params} from "react-router-dom"
import RecipeParams from "./home/recipe/Recipe"
import Header from "./header/Header"


const HeaderRouter: React.FC<{}> = ({}) => {
    return <BrowserRouter>
        <Routes>
        <Route path="/" element={<Header title="Все рецепты" />}/>
        <Route path="/auth" element={<Header title="Вход" />}/>
        <Route path="/recipes/:id" element={<Header title="Рецепт"/>}/>
        <Route path="*" element={<Header title="Страница не найдена"/>}/>
        </Routes>
    </BrowserRouter>
}

export default HeaderRouter