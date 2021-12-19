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
  Icon,
  LinkBox,
} from "@chakra-ui/react";

import RoundBox from "components/base/RoundBox";
import LoginIcon from 'components/icon/Login'

export interface NavbarProps {
    items: {
        name: string,
        ref: string,
    }[]
}

const Navbar: React.FC<NavbarProps> = (props) => {
    const [items, setItems] = React.useState(props.items);
    
    return (
    <Box
        width="100%"
        display="flex"
        justifyContent = 'space-between'
    >
        <Box
            display="flex" columnGap="10px"
            width="fit-content"
        >
            {items.map(item =>
            <Link 
                key={item.name} href={item.ref}
                width="100%" textStyle="body"
            >
                {item.name}
            </Link>
            )}
        </Box>
        
        <Link href="/auth/signin"> <RoundBox width="190px" height="40px" alignItems="center">
            <Box marginY="auto" marginX="8px"> <LoginIcon /> </Box>
            <Text marginLeft="2px" textStyle="body">Войти</Text>
        </RoundBox></Link>

    </Box>
    )
}

export default Navbar
