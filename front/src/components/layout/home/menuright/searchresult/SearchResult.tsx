import { ExternalLinkIcon } from "@chakra-ui/icons";
import {
  Box,
  Heading,
  HStack,
  Image,
  Link,
  Text,
  VStack,
} from "@chakra-ui/react";
import React, { useState } from "react";
import Spinner from "../../../../spinner/Spinner";

interface SearchResultProps {}

const SearchResult: React.FC<SearchResultProps> = (props) => {
  const [state, setState] = useState({
    isLoading: false,
  });

  const { isLoading } = state;

  return (
    <>
      {isLoading ? (
        <Spinner />
      ) : (
        <>
          <Text my={3} color="primary-dark" fontWeight="bold">
            Our top picks for you :
          </Text>
          <VStack w="100%">
            <HStack
              w="100%"
              borderWidth="1px"
              borderRadius="lg"
              bg="white"
              overflow="hidden"
            >
              <Image src="https://dummyimage.com/150x150/000/fff" />

              <Box p={6} mt="1" fontWeight="semibold" w="100%">
                <Heading fontSize="2xl">A first recipe</Heading>
                <Link d="flex" flexDirection="row" alignItems="baseline">
                  <Text>Checkout how it is made</Text>
                  <ExternalLinkIcon ml={1} />
                </Link>
              </Box>
            </HStack>
            <HStack
              w="100%"
              borderWidth="1px"
              borderRadius="lg"
              bg="white"
              overflow="hidden"
            >
              <Image src="https://dummyimage.com/150x150/000/fff" />

              <Box p={6} mt="1" fontWeight="semibold" w="100%">
                <Heading fontSize="2xl">A second recipe</Heading>
                <Link d="flex" flexDirection="row" alignItems="baseline">
                  <Text>Checkout how it is made</Text>
                  <ExternalLinkIcon ml={1} />
                </Link>
              </Box>
            </HStack>
            <HStack
              w="100%"
              borderWidth="1px"
              borderRadius="lg"
              bg="white"
              overflow="hidden"
            >
              <Image src="https://dummyimage.com/150x150/000/fff" />

              <Box p={6} mt="1" fontWeight="semibold" w="100%">
                <Heading fontSize="2xl">A third recipe</Heading>
                <Link d="flex" flexDirection="row" alignItems="baseline">
                  <Text>Checkout how it is made</Text>
                  <ExternalLinkIcon ml={1} />
                </Link>
              </Box>
            </HStack>
          </VStack>
        </>
      )}
    </>
  );
};

export default SearchResult;
