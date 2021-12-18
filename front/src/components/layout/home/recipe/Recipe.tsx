import React from "react";
import {
    Box,
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

type State = {
    recipe?: RecipeT
}

class Recipe extends React.Component<PP, State> {
    id: number;
    
    constructor(props) {
      super(props);
      this.id = Number(this.props.match.id)
      this.state = {recipe:undefined}
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
        <Box>
            <h2>Recipe {this.id}</h2>
            {this.state.recipe &&
            <h2>{this.state.recipe.title}</h2>
            }
        </Box>
      )
    }  
}
const RecipeParams = () => WithParams(Recipe)
export default RecipeParams
  