import { Text, Heading, Divider, Box } from "@chakra-ui/react";
import IngredientsForm from "./ingredientsform/IngredientsForm";
import IngredientsList from "./ingredientslist/IngredientsList";
import SearchResult from "./searchresult/SearchResult";
import React from "react";

interface MenuRightProps {}

const MenuRight: React.FC<MenuRightProps> = (props) => {
  return (
    <Box
      mt="5"
      ml="10"
      d="flex"
      flexDir="column"
      alignItems="start"
      justifyContent="space-around"
    >
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
    </Box>
  );
};

export default MenuRight;
