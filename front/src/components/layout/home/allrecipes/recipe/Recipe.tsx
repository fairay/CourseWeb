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
import photoRecipe from "img/photoRecipe.png"
import ClockIcon from "components/icon/Clock";
import FullLike from "components/icon/FullLike";
import {Recipe as RecipeI} from "types/Resipe";

import GetLikes from "postApi/likes/Get";

interface RecipeProps extends RecipeI {}


const Recipe: React.FC<RecipeProps> = (props) => {
    var path = "/recipes/" + props.id;
    const [likes, change] = React.useState(0);

    async function getLikes() {
        var data = await GetLikes(props.id)
        if (data.status == 200) {
            change(data.content)
        }
    }

    getLikes();
    
    return (
        <Link style={{ textDecoration: 'none' }}
            href={path}
        >
            <RoundBox
            w="480px"
            borderColor="title"
            alignItems="stretch"
            >
                <Image src={photoRecipe} 
                    borderRadius={"10px 0 0 10px"} 
                    width="200" height="300"
                />

                <Box 
                    display="flex" 
                    justifyContent="space-between" 
                    flexDirection="column"
                    m="10px"
                >
                    <VStack>
                        <Box w="100%" h="77">
                            <Text textStyle="subtitle" color="title" align="center">
                                {props.title}
                            </Text>
                        </Box>

                        <Box alignSelf="flex-start">
                            <Text  textStyle="alt-body" color="text">
                                {props.description}
                            </Text>
                        </Box>
                    </VStack>

                    <HStack>
                        <RoundBox width="120px" height="30px"
                            borderColor="accent-1" alignItems="center"> 
                                <Box marginY="auto" marginX="8px"> <ClockIcon /> </Box>
                                <Text marginLeft="2px" textStyle="body"> {props.duration} </Text>
                        </RoundBox>

                        <RoundBox width="120px" height="30px"
                            borderColor="accent-1" alignItems="center"> 
                                <Box marginY="auto" marginX="8px"> <FullLike /> </Box>
                                <Text marginLeft="2px" textStyle="body"> {likes} </Text>
                        </RoundBox>
                    </HStack>
                </Box>
            </RoundBox>
        </Link>
    )
}

export default Recipe;