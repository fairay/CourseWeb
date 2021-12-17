import { Image } from "@chakra-ui/react";
import React from "react";
import l1 from './Ok.svg';
import { ReactComponent as LoginSvg } from './login.svg';

interface IIconProps {
    img: string
}

const IIcon: React.FC<IIconProps> = (props) => {
  return <Image src={props.img} w="50px" />;
}
// export default l1;
export default LoginSvg;