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
import GetIngredient from "postApi/ingredients/Get";
import GetSteps from "postApi/steps/Get";
import GetLikes from "postApi/likes/Get";
import DeleteLike from "postApi/likes/Delete";
import AddLike from "postApi/likes/Add";

import {Recipe as RecipeT} from "types/Resipe"
import {Ingredient as IngredientT} from "types/Ingredient";
import {Step as StepT} from "types/Step";

import photoRecipe from "img/photoRecipe.png";
import DeleteIcon from "components/icon/Delete";
import EmptyLike from "components/icon/EmptyLike";
import FullLike from "components/icon/FullLike";
import ClockIcon from "components/icon/Clock";
import PlateIcon from "components/icon/Plate";

import RoundBox from "components/base/RoundBox";
import Ingredient from "../menuright/ingredient/Ingredient";
import Step from "../menuright/step/Step";

type State = {
    recipe?: RecipeT,
    liked: boolean,
    ingredients: Array<IngredientT>,
    steps: Array<StepT>,
    likes: number
}

class Recipe extends React.Component<PP, State> {
    id: number;
    
    constructor(props) {
      super(props);
      this.id = Number(this.props.match.id)
      this.state = {recipe:undefined, liked:false, ingredients:[], steps:[], likes:0}
    }

    deleteLike() {
        DeleteLike(this.id);
    }

    addLike() {
        AddLike(this.id);
    }

    getLikes() {
        GetLikes(this.id).then(data => {
            if (data.status == 200) {
                this.setState({likes: data.content})
            }
        });
    }

    componentDidMount() {
        GetRecipe(this.id).then(data => {
            if (data.status == 200) {
                this.setState({recipe: data.content})

                var elem = document.getElementById("title")
                if (elem && this.state.recipe)
                    elem.innerText = this.state.recipe.title

                elem = document.getElementById("author")
                if (elem && this.state.recipe)
                    elem.innerText = this.state.recipe.author
            }
        });

        GetIngredient(this.id).then(data => {
            if (data.status == 200) {
                this.setState({ingredients: data.content})
            }
        });

        GetSteps(this.id).then(data => {
            if (data.status == 200) {
                this.setState({steps: data.content})
            }
        });

        this.getLikes();
    }
  
    render() {
        var stringDuration = ""
        if (!this.state.recipe)
            stringDuration = "---"
        else if (this.state.recipe.duration < 90)
            stringDuration += this.state.recipe.duration + " мин"
        else if (this.state.recipe.duration < 60 * 24)
            stringDuration += Math.round(this.state.recipe.duration / 60) + " ч"
        else
            stringDuration += Math.round(this.state.recipe.duration / (60 * 24)) + " д"

        return(
            <VStack display="flex" w="100%">
                <Box display="flex" w="100%">
                    <HStack position="relative" paddingRight="15px">
                        <Image src={photoRecipe} maxWidth="none"
                            borderRadius="10px" 
                            width="336" height="400"
                        />

                        <VStack>
                            <Box position="absolute" right="0px" top="0px"> <DeleteIcon width="50px" height="40px" /> </Box>
                            
                            <Button display="contents" onClick={() => {                                
                                {this.state.liked && this.deleteLike()};
                                {!this.state.liked && this.addLike()};
                                this.setState({liked:!this.state.liked});
                                this.getLikes();
                            }
                            }>
                                <Box position="absolute" right="0px" bottom="0px">
                                    {!this.state.liked && <EmptyLike width="50px" height="40px" />}
                                    {this.state.liked && <FullLike width="50px" height="40px" />}
                                </Box>
                            </Button>
                        </VStack>
                    </HStack>

                    <Box paddingLeft="40px" w="100%">
                        <HStack display="flex" w="100%" columnGap="30px" paddingLeft="15px">
                            <RoundBox width="120px" height="30px" 
                                borderColor="accent-1" alignItems="center"> 
                                    <Box marginY="auto" marginX="8px"> <ClockIcon /> </Box>
                                    <Text marginLeft="2px" textStyle="body"> {stringDuration}  </Text>
                            </RoundBox>

                            <RoundBox width="120px" height="30px" 
                                borderColor="accent-1" alignItems="center"> 
                                    <Box marginY="auto" marginX="8px"> <FullLike /> </Box>
                                    <Text marginLeft="2px" textStyle="body"> {this.state.likes} </Text>
                            </RoundBox>

                            <RoundBox width="120px" height="30px"
                                borderColor="accent-1" alignItems="center"> 
                                    <Box marginY="auto" marginX="8px"> <PlateIcon /> </Box>
                                    <Text marginLeft="2px" textStyle="body"> {this.state.recipe?.portion_num} шт</Text>
                            </RoundBox>
                        </HStack>

                        <Text textStyle="alt-body" color="text" 
                            align="left" h="156px" 
                            paddingTop="20px" paddingLeft="15px">
                            {this.state.recipe?.description}
                        </Text>

                        <Box>
                            <Text textStyle="subtitle" color="title" align="left" maxH="156px"
                                paddingTop="20px" paddingLeft="15px">
                                Ингредиенты
                            </Text>

                            <RoundBox width="100%" height="192px" padding="3px"
                                borderColor="add-1" alignItems="flex-start"> 
                                <Box>
                                    {this.state.ingredients.map(item =>
                                        <Ingredient {...item} key={item.title}/>
                                    )}
                                </Box>
                            </RoundBox>
                        </Box>
                    </Box>
                </Box>

                <Box w="100%">
                    <Text textStyle="subtitle" color="title" align="left" maxH="156px"
                            paddingTop="15px">
                        Шаги приготовления
                    </Text>

                    <Box>
                        {this.state.steps.map(item =>
                            <Step {...item} key={item.num}/>
                        )}
                    </Box>
                </Box>
            </VStack>
        
      )
    }  
}
const RecipeParams = () => WithParams(Recipe)
export default RecipeParams
  