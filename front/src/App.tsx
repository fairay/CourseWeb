import * as React from "react";
import HeaderRouter from "./components/layout/Header";
import Home from "./components/layout/Home";
import Fonts from "./components/Fonts";
import theme from "./extendTheme";
import { ChakraProvider, Box, Text, Menu, Container } from "@chakra-ui/react";

export const App = () => (
  <ChakraProvider theme={theme}>
    <Fonts />
    <HeaderRouter  />
    <Home />
  </ChakraProvider>
);
