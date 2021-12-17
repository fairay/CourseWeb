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
import GetRecipe from "postApi/GetRecipe"
import Recipe from "./Recipe";

type State = {
    postContent?: any
}

class RecipeBox extends React.Component {
    public state: State = {
        postContent: null
    }
    constructor(props) {
        super(props);
        this.state = {
            postContent: Array()
        }
    }

    componentDidMount() {
        GetRecipe.GetRecipe().then(data => {
            if (data.status == 200)
                this.setState({postContent: data.content})
        });
    }

    render() {
        return (
            <Box 
                display='flex'
                flexDirection="row"
                flexWrap="wrap"
                justifyContent="space-between"
                rowGap="40px"
            >
                {this.state.postContent.map(item =>
                    <Recipe {...item} key={item.id}/>
                )}
            </Box>
        )
    }
}

export default RecipeBox; 