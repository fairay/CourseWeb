import { Box, Container, Divider } from "@chakra-ui/react";
import React from "react";
import AllRecipes from "./home/allrecipes/AllRecipes";

interface HomeProps {}

const Home: React.FC<HomeProps> = ({}) => {
  return (
    <Box backgroundColor="bg">
      <Container maxW="1000px" display="flex" paddingX="0px" paddingY="30px" h="95%" alignSelf="center"> 
        <AllRecipes />
      </Container>
    </Box>
  );
};

export default Home;
