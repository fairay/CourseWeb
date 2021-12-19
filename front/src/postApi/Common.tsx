import { Recipe } from "types/Resipe";

export const backUrl = "http://localhost:8000";

export type AllRecipeResp = {
    status: number,
    content: Recipe[] | string
}
