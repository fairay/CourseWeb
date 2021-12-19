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

export interface TitlesProps {
  subtitle?: string
  title: string
  undertitle?: string
}

export const Titles: React.FC<TitlesProps> = (props) => {
  return (
      <Container
        width="100%" maxWidth="none"
        d="flex"
        alignItems="center"
        justifyContent="space-around"
      >
        <List display="block" width='100%' align='center'>
          <ListItem mx="3">
            <Text  textStyle="subtitle" color="accent-1" id="subtitle">
              {props.subtitle ? props.subtitle : "ㅤ"}
            </Text>
          </ListItem>
          <ListItem width='100%'>
            <Text textStyle="header" color="title" id="title">
              {props.title}
            </Text>
          </ListItem>
          <ListItem mx="3">
            <Text textStyle="text" color="text" id="undertitle">
              {props.undertitle ? props.undertitle : "ㅤ"}
            </Text>
          </ListItem>
        </List>
      </Container>
  );
};
