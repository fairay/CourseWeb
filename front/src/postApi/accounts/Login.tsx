import axios from "axios";
import { backUrl } from "../Common";
import { Account } from "types/Account";

export const Login = async function(data: Account, setCookie) {
    const response = await axios.post(backUrl + `/accounts/login`, data, {withCredentials: true}).catch((error) => {
        return {
            status: error.response?.status,
        };
    });

    setCookie("login", data.login, {path: '/'})

    return {
        status: response?.status,
    };
}
