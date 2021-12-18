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
import {Ingredient as IngredientT} from "types/Ingredient";

interface IngredientProps extends IngredientT{}

const Ingredient: React.FC<IngredientProps> = (props) => {
    const [hide, show] = React.useState(true);

    return (
        <RoundBox display="inline-block" margin="5px" position="relative" borderColor="accent-1" 
            onMouseEnter={() => show(false)}
            onMouseLeave={() => show(true)}> 
            <HStack>
                <VStack alignItems="left" marginLeft="10px" marginRight="20px" minWidth="100px">
                    <Text textStyle="body" color="text" height="30px" m="0px">
                        {props.title}
                    </Text>

                    <Text textStyle="caption" color="accent-3" height="20px"
                        style={{margin: "0"}} >
                        {props.amount}
                    </Text>
                </VStack>

                <Button display="contents" visibility={hide ? "hidden" : "visible"} /*onClick={() => deleteItem(id)}*/>
                    <Box position="absolute" right="0px" top="0px"><CancelIcon /> </Box>
                </Button>

            </HStack>

        </RoundBox>
    )
}

export default Ingredient;