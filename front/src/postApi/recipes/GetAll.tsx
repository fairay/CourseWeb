import axios from "axios";
import { backUrl } from "../Common";

const GetRecipes = async function() {
    const response = await axios.get(backUrl + "/recipes");
    return {
        status: response.status,
        content: response.data
    };
}

export default GetRecipes