import { Box, Container, Divider } from "@chakra-ui/react";
import React from "react";
import MenuRight from "./menuright/MenuRight";

interface HomeProps {}

const Home: React.FC<HomeProps> = ({}) => {
  return (
    <Box backgroundColor="bg">
      <Container maxW="1000px" display="flex" paddingX="0px" paddingY="30px">
        <Divider
          h="95%"
          alignSelf="center"
          borderColor="primary-dark"
          orientation="vertical"
        />
        <MenuRight />
      </Container>
    </Box>
  );
};

export default Home;
