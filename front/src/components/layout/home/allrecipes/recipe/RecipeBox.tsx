import {
    Box,
  } from "@chakra-ui/react";
import React, { useState } from "react";
import { AllRecipeResp } from "postApi/Common"
import Recipe from "./Recipe";

interface RecipeBoxProps {
    getCall: () => Promise<AllRecipeResp>
}

type State = {
    postContent?: any
}

class RecipeBox extends React.Component<RecipeBoxProps, State> {
    constructor(props) {
        super(props);
        this.state = {
            postContent: Array()
        }
    }

    async getAll() {
        var data = await this.props.getCall()
        if (data.status == 200)
            this.setState({postContent: data.content})
    }

    componentDidMount() {
        this.getAll()
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