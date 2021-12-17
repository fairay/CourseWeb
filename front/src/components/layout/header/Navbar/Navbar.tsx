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
            display="flex"
            width="fit-content"
        >
            {/* # TODO Link to component */}
            {items.map(item =>
            <Link 
                key={item.name}
                marginRight="10px"
                width="100%" href={item.ref} textStyle="body"
            >
                {item.name}
            </Link>
            )}
        </Box>
        
        <RoundBox width="190px" height="40px"
        > <Link href="/auth" alignItems="center" display="flex" width="100%">
            <Box marginY="auto" marginX="8px"> <LoginIcon /> </Box>
            <Text marginLeft="2px" textStyle="body">Войти</Text>
        </Link></RoundBox>

    </Box>
    )
}

export default Navbar
