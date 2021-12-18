import React from "react";
import { Params, useParams } from "react-router-dom";

export function WithParams(T) {
    return <T match={useParams()} />;
}
  
export type PP = {
    match: Readonly<Params<string>>
  }