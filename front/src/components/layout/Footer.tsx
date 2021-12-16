import {
  Box,
  Container,
  Spacer,
  List,
  ListItem,
  Link,
  Text,
} from "@chakra-ui/react";
import React from "react";
import Logo from "../Logo";

interface FooterProps {}

const Footer: React.FC<FooterProps> = (props) => {
  return (
    <Box as="footer" bg="primary" px="10" py="3">
      <Container
        maxW="70%"
        d="flex"
        flexDirection="column"
        alignItems="center"
        justifyContent="space-around"
      >
        <Link color="secondary">
          <Text fontSize="md" fontWeight="semibold" color="secondary">
            Found a bug or want to file a request ? Contact the devs
          </Text>
        </Link>
      </Container>
    </Box>
  );
};

export default Footer;
