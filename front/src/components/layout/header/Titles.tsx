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

interface TitlesProps {}

const Titles: React.FC<TitlesProps> = (props) => {
  return (
      <Container
        width="100%"
        d="flex"
        alignItems="center"
        justifyContent="space-around"
      >
        <List display="block" width='100%' align='center'>
          <ListItem mx="3">
            <Text  textStyle="subtitle" color="accent-1">
            Надзаголовок
            </Text>
          </ListItem>
          <ListItem width='100%'>
            <Text textStyle="header" color="title">
              Заголовок
            </Text>
          </ListItem>
          <ListItem mx="3">
            <Text textStyle="text" color="text">
            Подзаголовок
            </Text>
          </ListItem>
        </List>
      </Container>
  );
};

export default Titles;
