import {
    Box,
    Button,
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
    let status = true;

    function changeVisibility(e) {
        /*e.target.style.visibility = "hidden";*/
        console.log(e);
        //console.log(e);
        //document.getElementById("Cancel").style.visibility = "hidden";
        
    }

    return (
        <RoundBox display="inline-block" position="relative" borderColor="accent-1" 
            onMouseEnter={changeVisibility}> 
            <HStack>
                <VStack alignItems="left" marginLeft="10px" marginRight="20px" minWidth="100px">
                    <Text textStyle="body" color="text" height="30px" m="0px">
                        Ингредиент
                    </Text>

                    <Text textStyle="caption" color="accent-3" height="20px"
                        style={{margin: "0"}} >
                        100 грамм
                    </Text>
                </VStack>

                <Button display="contents" id="Cancel" /*onClick={() => deleteItem(id)}*/>
                    <Box position="absolute" right="0px" top="0px"><CancelIcon /> </Box>
                </Button>

            </HStack>

        </RoundBox>
    )
}

export default Ingredient;