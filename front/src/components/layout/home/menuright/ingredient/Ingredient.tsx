import {
    Box,
    Heading,
    HStack,
    Image,
    Link,
    Text,
    VStack,
  } from "@chakra-ui/react";
import React, { useState } from "react";

interface IngredientProps {}

const Ingredient: React.FC<IngredientProps> = (props) => {
    return (
        <VStack alignItems="left">
            <Text textStyle="body" color="text">
                Ингредиент
            </Text>

            <Text textStyle="caption" color="accent-3">
                Ингредиент
            </Text>
        </VStack>
    )
}

export default Ingredient;