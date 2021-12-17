import { Box, Container, Divider } from "@chakra-ui/react";
import React from "react";
import MenuRight from "./menuright/MenuRight";

interface HomeProps {}

const Home: React.FC<HomeProps> = ({}) => {
  return (
    <Box backgroundColor="bg">
      <Container maxW="1000px" d="flex" py={4} padding="0px">
        <Divider
          h="95%"
          alignSelf="center"
          borderColor="primary-dark"
          orientation="vertical"
        />
        <Box marginX="0px" marginTop="30px" flex={[2, 3, 4]}>
          <MenuRight />
        </Box>
      </Container>
    </Box>
  );
};

export default Home;
