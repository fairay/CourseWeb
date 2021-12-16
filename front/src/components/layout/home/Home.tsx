import { Box, Container, Divider } from "@chakra-ui/react";
import React from "react";
import MenuLeft from "./menuleft/MenuLeft";
import MenuRight from "./menuright/MenuRight";

interface HomeProps {}

const Home: React.FC<HomeProps> = ({}) => {
  return (
    <Box backgroundColor="secondary">
      <Container maxW="60%" d="flex" py={4}>
        <Box minW="30%" flex="1">
          <MenuLeft />
        </Box>
        <Divider
          h="95%"
          alignSelf="center"
          borderColor="primary-dark"
          orientation="vertical"
        />
        <Box flex={[2, 3, 4]}>
          <MenuRight />
        </Box>
      </Container>
    </Box>
  );
};

export default Home;
