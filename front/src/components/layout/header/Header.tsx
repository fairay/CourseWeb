import React from "react";
import {
  Flex,
  Box,
  Heading,
  Text,
  List,
  ListItem,
  Link,
  Container,
  Spacer,
} from "@chakra-ui/react";

import {Titles, TitlesProps} from "./Titles"
import Navbar from "./Navbar/Navbar"
import Input from "components/base/Input"

interface HeaderProps extends TitlesProps {
    role?: string
}

const Header: React.FC<HeaderProps> = (props) => {
    const navItems = [
        { name: "Рецепты", ref: "#" },
        { name: "Пользователи", ref: "#" },
    ]
    
    return (
    <Box 
        bg="bg-alt"
        px="10" py="1"
        alignItems="center"
    >
        <Container
            py = "0"
            marginTop="20px"
            marginBottom="15px"
            padding={0}
            maxW="1000px"
        >
            <Navbar items={navItems}/>
            <Titles {...props}/>
            <Input width="100%" bg='bg' marginTop="5px"/>
        </Container>
    </Box>
    );
}

export default Header;
