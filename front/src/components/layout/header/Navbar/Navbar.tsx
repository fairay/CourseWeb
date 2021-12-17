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
} from "@chakra-ui/react";
import L1 from 'icon/login.svg';
import { ReactComponent as Logo } from 'icon/login.svg';

interface NavbarProps {}

const Navbar: React.FC<NavbarProps> = (props) => {
    const [items, setItems] = React.useState([
        {
          name: "Рецепты",
          ref: "#",
        },
        {
            name: "Пользователи",
            ref: "#",
        },
      ]);
    
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
            {items.map((item) => (
            <Link 
                marginRight="10px"
                width="100%" href={item.ref} textStyle="body"
            >
                {item.name}
            </Link>
            ))}
        </Box>
        
        <Box width="190px" height="40px">
            <Link display="flex" alignItems="center">
                {/* # TODO How to deal with images */}
                <Logo fill='red' stroke='green'/>
                <Text textStyle="body">Войти</Text>
            </Link>
        </Box>

    </Box>
    )
}

export default Navbar
