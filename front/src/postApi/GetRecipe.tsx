import axios from "axios";
import { backUrl } from "./Common";


export default {
GetRecipe: function() {
    return axios.get(backUrl + "/recipes")
    .then(response => {
        console.log(response)
        return {
        status: response.status,
        content: response.data
        };
    })
}
}