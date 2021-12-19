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
import { useCookies } from "react-cookie";
import theme from "extendTheme";

import RoundBox from "components/base/RoundBox";
import LoginIcon from 'components/icon/Login'
import AuthorIcon from "components/icon/Author";
import DownArrowIcon from "components/icon/DownArrow";
import FullLikeIcon from "components/icon/FullLike";
import RecipeIcon from "components/icon/Recipes";
import LogoutIcon from "components/icon/Logout";

export interface NavbarProps {}

const navItems = {
    '': [
        { name: "Рецепты", ref: "/" },
    ],
    'admin': [
        { name: "Рецепты", ref: "/" },
        { name: "Пользователи", ref: "#" },
    ],
    'user': [
        { name: "Рецепты", ref: "/" },
    ],
}

const Navbar: React.FC<NavbarProps> = (props) => {
    let [cookie] = useCookies(['role', 'login']);
    let role = cookie.login ? 'admin' : ''

    const [items, setItems] = React.useState(navItems[role]);
    const [expanded, setExpanded] = React.useState(false);
    
    const RenderNoAuth = () => {
        return (
        <Link href="/auth/signin"> <RoundBox width="190px" height="40px" alignItems="center">
            <Box marginY="auto" marginX="8px"> <LoginIcon /> </Box>
            <Text marginLeft="2px" textStyle="body">Войти</Text>
        </RoundBox></Link>
        )
    }

    const RenderAuth = () => {
        if (expanded) {
        return (
        <RoundBox width="190px" height="40px" alignItems="center" justifyContent="space-between">
            <Link href="/me/likes"><Box marginY="auto" marginX="8px"> 
                <FullLikeIcon width="30" height="24"  fill={theme.colors.title} /> 
            </Box></Link>
            <Link href="/me/recipes"><Box marginY="auto" marginX="8px"> 
                <RecipeIcon  fill={theme.colors.title} /> 
            </Box></Link>
            <Link href="/auth/logout"><Box marginY="auto" marginX="8px"> 
                <LogoutIcon  fill={theme.colors.title} /> 
            </Box></Link>
            <Box marginY="auto" marginX="8px" onClick={() => setExpanded(false)} cursor="pointer"> 
                <DownArrowIcon fill={theme.colors.title} /> 
            </Box>
        </RoundBox>
        ) 
        } else {
        return (
        <RoundBox width="190px" height="40px" alignItems="center" justifyContent="space-between">
            <Box marginY="auto" marginX="8px"> 
                <AuthorIcon width="30" height="24"  fill={theme.colors.title} /> 
            </Box>
            <Text marginLeft="2px" textStyle="body">{cookie.login}</Text>
            <Box marginY="auto" marginX="8px" onClick={() => setExpanded(true)} cursor="pointer"> 
                <DownArrowIcon fill={theme.colors.title} style={{transform: "rotate(180deg)"}}/> 
            </Box>
        </RoundBox>
        )
        }
    }

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

        { role === '' && <RenderNoAuth /> || <RenderAuth />}
    </Box>
    )
}

export default Navbar
