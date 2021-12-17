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

import RoundBox from "components/base/RoundBox";
import CancelIcon from "components/icon/Cancel";

interface IngredientProps {}

const Ingredient: React.FC<IngredientProps> = (props) => {
    return (
        <RoundBox display="inline-block" borderColor="accent-1"> 
            <VStack alignItems="left" marginLeft="10px" marginRight="20px" minWidth="100px">
                <Text textStyle="body" color="text" height="30px" m="0px">
                    Ингредиент
                </Text>

                <Text textStyle="caption" color="accent-3" height="20px"
                    style={{margin: "0"}} >
                    100 грамм
                </Text>
            </VStack>

        </RoundBox>
    )
}

export default Ingredient;