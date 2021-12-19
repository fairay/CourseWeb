import { Text, Heading, Divider, Box } from "@chakra-ui/react";
import GetRecipes from "postApi/recipes/GetAll";
import React from "react";
import Recipe from "./recipe/Recipe";
import RecipeBox from "./recipe/RecipeBox";
import RecipeDtl from "./recipe/RecipeDetailed"

interface AllRecipesProps {}

const AllRecipes: React.FC<AllRecipesProps> = (props) => {
  return (
    <Box
      d="flex" width="100%"
      flexDir="column"
      alignItems="start"
      justifyContent="space-around"
    >
      <RecipeBox getCall={GetRecipes}/>
    </Box>
  );
};

export default AllRecipes;
