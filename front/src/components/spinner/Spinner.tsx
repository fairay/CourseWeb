import { Image, Img } from "@chakra-ui/react";
import spinner from "./Spinner.gif";
import React from "react";

const Spinner: React.FC = () => {
  return <Image src={spinner} />;
};

export default Spinner;
