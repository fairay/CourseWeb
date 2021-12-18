import axios from "axios";
import { backUrl } from "../Common";


const GetIngredient = async function(id: number) {
    const response = await axios.get(backUrl + `/recipes/${id}/ingredients`);
    console.log(response);
    return {
        status: response.status,
        content: response.data
    };
}
export default GetIngredient;