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
    <Box bg="primary" px="10" py="3">
      <Container
        maxW="70%"
        d="flex"
        alignItems="center"
        justifyContent="space-around"
      >
        <Logo />
        <Spacer flex="2" />
        <List display="flex">
          <ListItem mx="3">
            <Link color="secondary">
              <Text fontSize="xl" fontWeight="semibold">
                Home
              </Text>
            </Link>
          </ListItem>
          <ListItem mx="3">
            <Link color="secondary">
              <Text fontSize="xl" fontWeight="semibold" color="secondary">
                About
              </Text>
            </Link>
          </ListItem>
          <ListItem mx="3">
            <Link color="secondary">
              <Text fontSize="xl" fontWeight="semibold" color="secondary">
                Contact Us
              </Text>
            </Link>
          </ListItem>
        </List>
      </Container>
    </Box>
  );
};

export default Navbar;
