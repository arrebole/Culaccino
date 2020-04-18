import React from "react"
import * as types from "../types";

export default function Stage(props: types.Article) {
    return (
        <article className="w-auto bg-white border border-gray-400 shadow-lg rounded-sm my-3" key={props.name}>
            <div className="text-center pt-3">
                <div className="text-2xl text-indigo-700"><a href={`/articles/${props.name}`}>{props.name}</a></div>
                <div className="text-xs text-gray-600">
                    <span className="mx-1">Created: {new Date(props.created_at).toLocaleDateString()}</span>
                    <span className="mx-1">Updated: {new Date(props.created_at).toLocaleDateString()}</span>
                </div>
            </div>
            <img className="mx-auto my-3" src={props.cover} alt="cover"/>
            <div className="text-center p-4 text-gray-700">{props.summary}</div>
        </article>
    )
}