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

import Titles from "./Titles"
import Navbar from "./Navbar/Navbar"

interface HeaderProps {}

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
            maxW="1000px"
        >
            <Navbar />
            <Titles />
        </Container>
    </Box>
    );
}

export default Header;
