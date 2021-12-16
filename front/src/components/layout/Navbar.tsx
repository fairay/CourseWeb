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
import Logo from "../Logo";

interface NavbarProps {}

const Navbar: React.FC<NavbarProps> = (props) => {
  return (
    <Box bg="bg-alt" px="10" py="3">
      <Container
        maxW="60%"
        d="flex"
        alignItems="center"
        justifyContent="space-around"
      >
        <List display="block" width='100%' align='center'>
          <ListItem mx="3">
            <Link>
              <Text  textStyle="subtitle" color="accent-1">
              Надзаголовок
              </Text>
            </Link>
          </ListItem>
          <ListItem width='100%'>
            <Link>
              <Text textStyle="header" color="title">
                Заголовок
              </Text>
            </Link>
          </ListItem>
          <ListItem mx="3">
            <Link>
              <Text textStyle="text" color="text">
              Подзаголовок
              </Text>
            </Link>
          </ListItem>
        </List>
      </Container>
    </Box>
  );
};

export default Navbar;
