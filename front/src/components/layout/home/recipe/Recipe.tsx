import React from "react";
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

import {WithParams, PP} from "../../Warpers"
import GetRecipe from "postApi/recipes/Get"
import {Recipe as RecipeT} from "types/Resipe"

import photoRecipe from "img/photoRecipe.png";
import DeleteIcon from "components/icon/Delete";
import EmptyLike from "components/icon/EmptyLike";
import FullLike from "components/icon/FullLike";
import ClockIcon from "components/icon/Clock";
import PlateIcon from "components/icon/Plate";

import RoundBox from "components/base/RoundBox";

type State = {
    recipe?: RecipeT,
    liked: boolean
}

class Recipe extends React.Component<PP, State> {
    id: number;
    
    constructor(props) {
      super(props);
      this.id = Number(this.props.match.id)
      this.state = {recipe:undefined, liked:false}
    }

    componentDidMount() {
        GetRecipe(this.id).then(data => {
            if (data.status == 200) {
                this.setState({recipe: data.content})

                var title = document.getElementById("title")
                if (title && this.state.recipe)
                    title.innerText = this.state.recipe.title
            }
        });
    }
  
    render() {
        return(
            <Box display="flex" w="100%">
            <HStack position="relative" paddingRight="15px">
                <Image src={photoRecipe} 
                    borderRadius="10px" 
                    width="336" height="400"
                />

                <VStack>
                    <Box position="absolute" right="0px" top="0px"> <DeleteIcon width="50px" height="40px" /> </Box>
                    
                    <Button display="contents" onClick={() => this.setState({liked:!this.state.liked})}>
                        <Box position="absolute" right="0px" bottom="0px">
                            {!this.state.liked && <EmptyLike width="50px" height="40px" />}
                            {this.state.liked && <FullLike width="50px" height="40px" />}
                        </Box>
                    </Button>
                </VStack>
            </HStack>

            <Box>
                <HStack display="flex" justify-content="space-between">
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
            
        </Box>
      )
    }  
}
const RecipeParams = () => WithParams(Recipe)
export default RecipeParams
  