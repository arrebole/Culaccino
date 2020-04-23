import React from "react";
import * as types from "../types";

export default function Banner(props: types.Article) {

    return (
        <div className="bg-white border border-gray-400 shadow rounded-sm p-5 my-8 text-center">
            <h2 className="text-2xl"> {props.name} </h2>
            <div className="text-xs text-gray-600">
                <span className="mx-1">Created: {new Date(props.created_at).toLocaleDateString()}</span>
                <span className="mx-1">Updated: {new Date(props.created_at).toLocaleDateString()}</span>
            </div>
        </div>
    )
}