import axios from "axios";
import { backUrl } from "../Common";


const GetSteps = async function(id: number) {
    const response = await axios.get(backUrl + `/recipes/${id}/steps`);
    console.log(response);
    return {
        status: response.status,
        content: response.data
    };
}
export default GetSteps;