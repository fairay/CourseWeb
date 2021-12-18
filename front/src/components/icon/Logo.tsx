import { BoxProps, extendTheme, Image } from "@chakra-ui/react";
import theme from "extendTheme";
import React, { SVGProps } from "react";
import logo from "../logo.svg";

interface LogoProps extends SVGProps{
  width?: string,
  height?: string,
  fill?: string,
  visibility?: string
}


export default LogoProps
