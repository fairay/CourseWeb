import { Box, Container, Divider } from "@chakra-ui/react";
import React from "react";
import {BrowserRouter, Routes, Route, useParams, RouteProps, Params} from "react-router-dom"

import AllRecipes from "./home/allrecipes/AllRecipes";
import LikedRecipes from "./home/allrecipes/LikedRecipes";
import AuthorRecipes from "./home/allrecipes/AuthorRecipes";
import RecipeParams from "./home/recipe/Recipe"
import Login from "./home/auth/Login"
import SignUp from "./home/auth/Signup"

interface HomeProps {}

const Home: React.FC<HomeProps> = ({}) => {
  return (
    <Box backgroundColor="bg" minH="100vh" h="auto">
      <Container maxW="1000px" minH="95%"
        display="flex" 
        paddingX="0px" paddingY="30px"  
        alignSelf="center" justifyContent="center"
        textStyle="body"
      > 
        <Routing />
      </Container>
    </Box>
  );
};

function Routing () {
  return <BrowserRouter>
    <Routes>
      <Route path="/" element={<AllRecipes />}/>
      <Route path="/me/likes" element={<LikedRecipes />}/> 
      <Route path="/me/recipes" element={<AuthorRecipes />}/>
      <Route path="/auth/signin" element={<Login />}/>
      <Route path="/auth/signup" element={<SignUp />}/>
      <Route path="/recipes/:id" element={<RecipeParams />}/>
      <Route path="*" element={<NotFound />}/>
    </Routes>
  </BrowserRouter>
}

function Auth () {
  return <h1>Authorization page</h1>
}

function NotFound () {
  return <h1>Page not Found</h1>
}

export default Home;
