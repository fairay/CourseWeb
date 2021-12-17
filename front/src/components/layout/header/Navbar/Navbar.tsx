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
import LoginIcon from 'components/icon/Login'

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
        
        <Link 
            width="190px" height="40px"
            border={"1px solid #000000"}
            borderRadius="10px"
            display="flex" alignItems="center"
            href="/auth"
        >
                <Box marginY="auto" marginX="8px"> <LoginIcon /> </Box>
                <Text marginLeft="2px" textStyle="body">Войти</Text>
        </Link>

    </Box>
    )
}

export default Navbar
