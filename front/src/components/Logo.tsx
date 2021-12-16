import { Image } from "@chakra-ui/react";
import React from "react";
import logo from "../logo.svg";

interface LogoProps {}

const Logo: React.FC<LogoProps> = (props) => {
  return <Image src={logo} w="190px" />;
};

export default Logo;
