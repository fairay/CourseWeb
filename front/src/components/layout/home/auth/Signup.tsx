import { Text, Heading, Divider, Box, Button, Link, FormControl } from "@chakra-ui/react";
import React from "react";
import theme from "extendTheme";
import Input from "components/base/Input";
import RoundButton from "components/base/Button";
import RoundBox from "components/base/RoundBox";
import { Account } from "types/Account";
import { Create as CreateQuery } from "postApi/accounts/Create";
import { NavigateFunction, useNavigate } from "react-router-dom";


function SignUpWrapper(props) {
    let navigate = useNavigate();
    return <SignUp navigate={navigate} {...props}/>
}

type SignUpProps = {
    navigate: NavigateFunction
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

    submit(e: React.MouseEvent) {
        if (this.acc.password != this.repPassword)
            return this.highlightNotMatch()

        e.target.disabled = true
        CreateQuery(this.acc).then(data => {
            if (data.status == 200) {
                this.props.navigate("/")
            } else {
                e.target.disabled = false
                var title = document.getElementById("undertitle")
                if (title)
                    title.innerText = "Ошибка создания аккаунта!"
            }
        });
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
                    onInput={event => this.setLogin(event.target.value)}/>
                </FormControl>
                <FormControl isRequired>
                    <Input name="password" type="password" w="100%" placeholder="Введите пароль"
                    onInput={event => this.setPassword(event.target.value)}/>
                </FormControl>
                <FormControl isRequired>
                    <Input name="rep-password" type="password" w="100%" placeholder="Повторите пароль"
                    onInput={event => this.setRepPassword(event.target.value)}/>
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