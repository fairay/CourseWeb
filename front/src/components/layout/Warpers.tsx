import React from "react";
import { useCookies } from "react-cookie";
import { NavigateFunction, Params, useNavigate, useParams } from "react-router-dom";

export function WithParams(T) {
    let [cookie] = useCookies(['role', 'login']);
    let navigate = useNavigate()
    return <T match={useParams()} navigate={navigate} cookie={cookie} />;
}
  
export type PP = {
    match: Readonly<Params<string>>
    navigate: NavigateFunction
    cookie: {
        role?: string;
        login?: string;
    }
}