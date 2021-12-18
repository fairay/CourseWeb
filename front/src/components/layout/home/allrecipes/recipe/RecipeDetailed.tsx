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

import photoRecipe from "img/photoRecipe.png";
import DeleteIcon from "components/icon/Delete";
import EmptyLike from "components/icon/EmptyLike";
import FullLike from "components/icon/FullLike";
import ClockIcon from "components/icon/Clock";
import PlateIcon from "components/icon/Plate";

import RoundBox from "components/base/RoundBox";


interface RecipeDtlProps {}

const RecipeDtl: React.FC<RecipeDtlProps> = (props) => {
    const [hide, show] = React.useState(true);

    return (
        <Box display="flex">
            <HStack position="relative" paddingRight="15px">
                <Image src={photoRecipe} 
                    borderRadius="10px" 
                    width="336" height="400"
                />

                <VStack>
                    <Box position="absolute" right="0px" top="0px"> <DeleteIcon width="50px" height="40px" /> </Box>
                    
                    <Button display="contents" onClick={() => show(!hide)}>
                        <Box position="absolute" right="0px" bottom="0px" visibility={hide ? "visible" : "hidden"}>
                            <EmptyLike width="50px" height="40px"/> 
                        </Box>

                        <Box position="absolute" right="0px" bottom="0px" visibility={hide ? "hidden" : "visible"}>
                            <FullLike width="50px" height="40px"/> 
                        </Box>
                    </Button>
                </VStack>
            </HStack>

            <HStack>
                <RoundBox width="120px" height="30px"
                    borderColor="accent-1" alignItems="center"> 
                        <Box marginY="auto" marginX="8px"> <ClockIcon /> </Box>
                        <Text marginLeft="2px" textStyle="body"> TODO </Text>
                </RoundBox>

                <RoundBox width="120px" height="30px"
                    borderColor="accent-1" alignItems="center"> 
                        <Box marginY="auto" marginX="8px"> <FullLike /> </Box>
                        <Text marginLeft="2px" textStyle="body">1 024</Text>
                </RoundBox>

                <RoundBox width="120px" height="30px"
                    borderColor="accent-1" alignItems="center"> 
                        <Box marginY="auto" marginX="8px"> <PlateIcon /> </Box>
                        <Text marginLeft="2px" textStyle="body">TODO</Text>
                </RoundBox>
            </HStack>
            
        </Box>
    )
}

export default RecipeDtl;