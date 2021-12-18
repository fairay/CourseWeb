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

interface TitlesProps {
  subtitle?: string
  title: string
  undertitle?: string
}

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
            <Text  textStyle="subtitle" color={props.undertitle ? "accent-1" : "bg-alt"}>
              {props.subtitle ? props.subtitle : "-"}
            </Text>
          </ListItem>
          <ListItem width='100%'>
            <Text textStyle="header" color="title">
              {props.title}
            </Text>
          </ListItem>
          <ListItem mx="3">
            <Text textStyle="text" color={props.undertitle ? "text" : "bg-alt"} >
              {props.undertitle ? props.undertitle : "-"}
            </Text>
          </ListItem>
        </List>
      </Container>
  );
};

export default Titles;
