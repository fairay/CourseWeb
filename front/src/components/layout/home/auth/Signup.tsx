import { Text, Heading, Divider, Box, Button, Link, FormControl } from "@chakra-ui/react";
import React from "react";
import theme from "extendTheme";
import Input from "components/base/Input";
import RoundButton from "components/base/Button";
import RoundBox from "components/base/RoundBox";
import { Account } from "types/Account";
import { Create as CreateQuery } from "postApi/accounts/Create";
import { Login as LoginQuery } from "postApi/accounts/Login";
import { NavigateFunction, useNavigate } from "react-router-dom";
import { CookieSetOptions } from "universal-cookie";
import { useCookies } from "react-cookie";


function SignUpWrapper(props) {
    let navigate = useNavigate();
    let [cookie, setCookie] = useCookies(['token', 'role', 'login']);
    return <SignUp navigate={navigate}  cookie={cookie} setCookie={setCookie}  {...props}/>
}

type SignUpProps = {
    navigate: NavigateFunction
    cookie: {
        token?: string;
        role?: string;
        login?: string;
    }
    setCookie: (name: "token" | "role" | "login", value: any, options?: CookieSetOptions | undefined) => void
}

class SignUp extends React.Component<SignUpProps> {
    acc: Account = {login: ""}
    repPassword: string = ""
    constructor(props) {
        super(props)
    }

    setLogin(val: string) {
        this.acc.login = val
    }
    setPassword(val: string) {
        this.acc.password = val
    }
    setRepPassword(val: string) {
        this.repPassword = val
    }

    highlightNotMatch() {
        let node1 = document.getElementsByName("password")[0]
        let node2 = document.getElementsByName("rep-password")[0]

        if (node1.parentElement && node2.parentElement) {
            node1.parentElement.style.borderColor = theme.colors["title"]
            node2.parentElement.style.borderColor = theme.colors["title"]
        }

        var title = document.getElementById("undertitle")
        if (title)
            title.innerText = "Пароли не совпадают!"
    }

    async submit(e: React.MouseEvent<HTMLButtonElement, MouseEvent>) {
        if (this.acc.password != this.repPassword)
            return this.highlightNotMatch()

        e.currentTarget.disabled = true
        var data = await CreateQuery(this.acc)
        if (data.status == 200) {
            await LoginQuery(this.acc, this.props.setCookie)
            this.props.navigate("/")
        } else {
            e.currentTarget.disabled = false
            var title = document.getElementById("undertitle")
            if (title)
                title.innerText = "Ошибка создания аккаунта!"
        };
    }

    render() {
        return <Box
            display="flex" width="70%"
            flexDir="column"
            alignItems="stretch"
            justifyContent="space-around"
            rowGap="70px"
        >
            <Box d="flex" flexDirection="column" rowGap="35px">
                <FormControl isRequired>
                    <Input name="login" w="100%" placeholder="Введите логин" 
                    onInput={event => this.setLogin(event.currentTarget.value)}/>
                </FormControl>
                <FormControl isRequired>
                    <Input name="password" type="password" w="100%" placeholder="Введите пароль"
                    onInput={event => this.setPassword(event.currentTarget.value)}/>
                </FormControl>
                <FormControl isRequired>
                    <Input name="rep-password" type="password" w="100%" placeholder="Повторите пароль"
                    onInput={event => this.setRepPassword(event.currentTarget.value)}/>
                </FormControl>
            </Box>

            <Box d="flex" flexDirection="column" rowGap="15px" alignItems="center">
                <RoundButton type="submit" onClick={event => this.submit(event)}
                    w="100%" bg="button" textColor="bg"
                >
                    Создать аккаунт
                </RoundButton>
                <Link href="/auth/signin">Войти</Link>
            </Box>
        </Box>
    }
}

export default SignUpWrapper