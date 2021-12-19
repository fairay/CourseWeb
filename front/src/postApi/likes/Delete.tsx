import axios from "axios";
import { backUrl } from "../Common";


const DeleteLike = async function(id: number) {
    const response = await axios.delete(backUrl + `/recipes/${id}/like`);
    return {
        status: response.status,
    };
}
export default DeleteLike;