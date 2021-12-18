import { Text, Heading, Divider, Box } from "@chakra-ui/react";
import IngredientsForm from "../menuright/ingredientsform/IngredientsForm";
import IngredientsList from "../menuright/ingredientslist/IngredientsList";
import SearchResult from "../menuright/searchresult/SearchResult";
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
      <RecipeBox />
      <Heading color="primary-dark">What's in your fridge :</Heading>
      <Divider borderColor="primary-dark" orientation="horizontal" mb={2} />
      <Text color="primary-dark" fontWeight="bold">
        We help you get the best dinner you can with what you have available.
      </Text>
      <Text color="primary-dark" fontWeight="bold">
        All you have to do is tell us what you can cook with.
      </Text>
      <IngredientsForm />
      <Text mt={3} color="primary-dark" fontWeight="bold">
        Your ingredients :
      </Text>
      <IngredientsList />
      <Heading color="primary-dark">Results :</Heading>
      <Divider borderColor="primary-dark" orientation="horizontal" mb={2} />
      <SearchResult />
      <RecipeDtl />
    </Box>
  );
};

export default AllRecipes;
