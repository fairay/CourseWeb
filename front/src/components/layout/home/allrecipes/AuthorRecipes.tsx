import { Text, Heading, Divider, Box } from "@chakra-ui/react";
import GetRecipes from "postApi/accounts/GetRecipes";
import React from "react";
import { useCookies } from "react-cookie";
import RecipeBox from "./recipe/RecipeBox";

interface AuthorRecipesProps {}

const AuthorRecipes: React.FC<AuthorRecipesProps> = (props) => {
    let [cookie] = useCookies(['role', 'login']);

    return (
    <Box
        d="flex" width="100%"
        flexDir="column"
        alignItems="start"
        justifyContent="space-around"
    >
        <RecipeBox getCall={() => GetRecipes(cookie.login)}/>
    </Box>
    );
};

export default AuthorRecipes;
