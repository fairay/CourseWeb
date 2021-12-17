import * as React from "react";
import Header from "./components/layout/header/Header";
import Footer from "./components/layout/Footer";
import Home from "./components/layout/home/Home";
import Fonts from "./components/Fonts";
import theme from "./extendTheme";
import { ChakraProvider, Box, Text, Menu, Container } from "@chakra-ui/react";

export const App = () => (
  <ChakraProvider theme={theme}>
    <Fonts />
    <Header  />
    <Home />
    <Footer />
  </ChakraProvider>
);
