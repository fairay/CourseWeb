import { extendTheme } from "@chakra-ui/react";

const theme = extendTheme({
  colors: {
    primary: "#d4503c",
    "primary-dark": "#a52818",
    secondary: "#eee5cf",
    third: "#ddbc6d",
  },
  fonts: {
    heading: "Unica One",
    body: "Raleway",
  },
});

export default theme;
