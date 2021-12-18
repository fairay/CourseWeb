import { Box, Container, Divider } from "@chakra-ui/react";
import React from "react";
import AllRecipes from "./home/allrecipes/AllRecipes";
import {BrowserRouter, Routes, Route} from "react-router-dom"

interface HomeProps {}

const Home: React.FC<HomeProps> = ({}) => {
  return (
    <Box backgroundColor="bg" minH="100%" h="auto">
      <Container maxW="1000px" display="flex" 
        paddingX="0px" paddingY="30px" minH="95%" 
        alignSelf="center" textStyle="body"
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
      <Route path="/auth" element={<Auth />}/>
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
