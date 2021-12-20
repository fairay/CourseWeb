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
import PushStep from "postApi/steps/Push";
import GetLikes from "postApi/likes/Get";
import DeleteLike from "postApi/likes/Delete";
import AddLike from "postApi/likes/Add";
import GetIsLiked from "postApi/likes/GetIs";
import DeleteIngredient from "postApi/ingredients/Delete";
import DeleteStep from "postApi/steps/Delete";
import PutIngredient from "postApi/ingredients/Post";
import DeleteRecipe from "postApi/recipes/Delete";

import {Recipe as RecipeT} from "types/Resipe"
import {Ingredient as IngredientT} from "types/Ingredient";
import {Step as StepT} from "types/Step";

import photoRecipe from "img/photoRecipe.png";
import DeleteIcon from "components/icon/Delete";
import EmptyLike from "components/icon/EmptyLike";
import FullLike from "components/icon/FullLike";
import ClockIcon from "components/icon/Clock";
import PlateIcon from "components/icon/Plate";
import AddIcon from "components/icon/Add";

import RoundBox from "components/base/RoundBox";
import Ingredient from "../menuright/ingredient/Ingredient";
import Step from "../menuright/step/Step";
import IngredientModal from "./IngredientModal";
import Input from "components/base/Input";
import ClockBox from "components/base/box/Clock";



type State = {
    recipe?: RecipeT,
    liked: boolean,
    ingredients: Array<IngredientT>,
    steps: Array<StepT>,
    likes: number,

    isAuthor: boolean,
}

class Recipe extends React.Component<PP, State> {
    id: number;
    
    constructor(props) {
      super(props);
      this.id = Number(this.props.match.id)
      this.state = {
          recipe:undefined, 
          liked:false, 
          ingredients:[], 
          steps:[],
          likes:0,
          isAuthor: false
        }
    }

    async deleteRecipe() {
        var data = await DeleteRecipe(this.id);
        if (data.status == 200) {
            window.location.href = '/';
        }
    }

    async deleteLike() {
        var data = await DeleteLike(this.id);
        if (data.status == 200) {
            this.setState({liked:false});
            this.getLikes();
        }
    }

    async addLike() {
        var data = await AddLike(this.id);
        if (data.status == 200) {
            this.setState({liked:true});
            this.getLikes();
        }
    }

    async getLikes() {
        var data = await GetLikes(this.id)
        if (data.status == 200) {
            this.setState({likes: data.content})
        }
    }

    async getIsLiked() {
        var data = await GetIsLiked(this.id)
        if (data.status == 200) {
            this.setState({liked: data.content})
        }
    }

    async createStep() {
        let step: StepT = {
            description:    '',
            title: '',
            num: 0,
            recipe: this.id
        }

        var data = await PushStep(this.id, step)
        if (data.status == 200) {
            var temp = this.state.steps
            temp.push(data.content)
            this.setState({steps: temp})
        }
    }

    async putIngredient(data: IngredientT) {
        var res = await PutIngredient(this.id, data)
        if (res.status == 201) {
            var ingArr = this.state.ingredients
            ingArr.push(data)
            this.setState({ingredients: ingArr})
        }
    }

    async deleteIngredient(title: string) {
        var data = await DeleteIngredient(this.id, title)
        if (data.status == 200) {
            var ingArr = this.state.ingredients
            ingArr = ingArr.filter(item => item.title != title)
            this.setState({ingredients: ingArr})
        }
    }

    async deleteStep(num: number) {
        var data = await DeleteStep(this.id, num)
        if (data.status == 200) {
            var temp = this.state.steps
            temp = temp.filter(item => item.num != num)
            this.setState({steps: temp})
        }
    }

    componentDidMount() {
        GetRecipe(this.id).then(data => {
            if (data.status == 200) {
                this.setState({recipe: data.content})
                this.setState({isAuthor: this.state.recipe?.author == this.props.cookie.login})

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
        if (this.props.cookie.login) this.getIsLiked();
    }
  
    render() {
        return(
            <VStack display="flex" w="100%">
                <Box display="flex" w="100%">
                    <HStack position="relative" paddingRight="15px">
                        <Image src={photoRecipe} maxWidth="none"
                            borderRadius="10px" 
                            width="336" height="400"
                        />

                        <VStack>
                            <Button display="contents" onClick={() => this.deleteRecipe()}>
                                <Box position="absolute" right="0px" top="0px"> 
                                    {this.state.isAuthor && <DeleteIcon width="50px" height="40px" /> }
                                </Box>
                            </Button>
                            
                            <Button display="contents" onClick={() => {                                
                                {this.state.liked 
                                    && this.deleteLike()
                                    || this.addLike()
                                };
                            }
                            }>
                                <Box position="absolute" right="0px" bottom="0px">
                                    {this.props.cookie.login 
                                    && (this.state.liked 
                                        && <FullLike width="50px" height="40px" /> 
                                        || <EmptyLike width="50px" height="40px" />)
                                    }
                                </Box>
                            </Button>
                        </VStack>
                    </HStack>

                    <Box paddingLeft="40px" w="100%">
                        <HStack display="flex" w="100%" columnGap="30px" paddingLeft="15px">
                            <ClockBox duration={this.state.recipe?.duration}/>

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
                            <HStack paddingTop="20px">
                                <Text textStyle="subtitle" color="title" align="left" maxH="156px"
                                    paddingLeft="15px">
                                    Ингредиенты
                                </Text>

                                {this.state.isAuthor &&
                                 <IngredientModal putCallback={(data: IngredientT) => this.putIngredient(data)}/>
                                }
                            </HStack>

                            <RoundBox width="100%" height="192px" padding="3px"
                                borderColor="add-1" alignItems="flex-start"> 
                                <Box>
                                    {this.state.ingredients.map(item =>
                                        <Ingredient {...item} key={item.title} 
                                        delCallback={
                                            this.state.isAuthor 
                                            ? (title) => this.deleteIngredient(title) 
                                            : undefined
                                        } 
                                        />
                                    )}
                                </Box>
                            </RoundBox>
                        </Box>
                    </Box>
                </Box>

                <Box w="100%">
                    <HStack paddingTop="20px">
                        <Text textStyle="subtitle" color="title" align="left" maxH="156px"
                                paddingTop="15px">
                            Шаги приготовления
                        </Text>

                        <Button display="contents" onClick={() => this.createStep()}>
                            {this.state.isAuthor &&
                                <Box paddingLeft="15px" paddingTop="15px"> <AddIcon /> </Box>}
                        </Button>
                    </HStack>

                    <Box>
                        {this.state.steps.map(item =>
                            <Step {...item} key={item.num}
                            delStepCallback={
                                this.state.isAuthor 
                                ? (num) => this.deleteStep(num)
                                : undefined
                            }
                            />
                        )}
                    </Box>
                </Box>
            </VStack>
        
      )
    }  
}
const RecipeParams = () => WithParams(Recipe)
export default RecipeParams
  