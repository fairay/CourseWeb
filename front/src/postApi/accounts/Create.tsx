import axios from "axios";
import { backUrl } from "../Common";
import { Account } from "types/Account";

export const Create = async function(data: Account) {
    console.log(data);
    const response = await axios.post(backUrl + `/accounts`, data).catch((error) => {
        console.log(error.response);    
        return {
            status: error.response?.status,
        };
    });

    console.log(response);
    return {
        status: response?.status,
    };
}
