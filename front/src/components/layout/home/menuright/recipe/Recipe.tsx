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
import Spinner from "../../../../spinner/Spinner";
import photoRecipe from "../../../../../img/photoRecipe.png";

interface RecipeProps {}

const Recipe: React.FC<RecipeProps> = (props) => {
    return (
        <HStack
        w="100%"
        borderWidth="1px"
        borderStyle="solid"
        borderRadius="10px"
        borderColor="title"
        alignItems="stretch"
        >
            <Image src={photoRecipe} style={{width: 200,
                                            height: 300,
                                            borderTopLeftRadius: 10,
                                            borderBottomLeftRadius: 10,
            }} 
            />

            <Box display="flex" justifyContent="space-between" flexDirection="column">
                <VStack>
                    <Box w="100%" h="77">
                        <Text textStyle="subtitle" color="title" align="center">
                        Торт Графские развалины
                        </Text>
                    </Box>

                    <Box>
                        <Text  textStyle="alt-body" color="text">
                        Необычайно вкусный торт, не пожалеете Приготовить безе. В миску вылить белки и добавить 240 грамм сахара. Установить миску на водяную баню и, энергично помешивая массу венчиком, дать сахару бла-бла
                        </Text>
                    </Box>
                </VStack>

                <Text>
                    Важная инфа в два овала
                </Text>
            </Box>

            

            



        </HStack>
    )
}

export default Recipe;