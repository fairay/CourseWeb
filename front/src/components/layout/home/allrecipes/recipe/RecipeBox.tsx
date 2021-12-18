import {
    Box,
  } from "@chakra-ui/react";
import React, { useState } from "react";
import GetRecipes from "postApi/recipes/GetAll"
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
        GetRecipes().then(data => {
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