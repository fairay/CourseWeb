import React from "react";
import { useCookies } from "react-cookie";
import { Params, useParams } from "react-router-dom";

export function WithParams(T) {
    let [cookie] = useCookies(['role', 'login']);
    return <T match={useParams()} cookie={cookie} />;
}
  
export type PP = {
    match: Readonly<Params<string>>
    cookie: {
        role?: string;
        login?: string;
    }
}