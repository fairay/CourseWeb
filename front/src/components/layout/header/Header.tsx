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
import {useCookies} from 'react-cookie';

interface HeaderProps extends TitlesProps {
    role?: string
    addField?: React.FC
}

const Header: React.FC<HeaderProps> = (props) => {
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
            <Navbar />
            <Titles {...props}/>
            {props.addField && <props.addField />}
            
        </Container>
    </Box>
    );
}

export default Header;
