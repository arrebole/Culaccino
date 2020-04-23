import React from "react";


export default function Pager(props: { onClick: Function }) {
    return (
        <div className="py-3">
            <button
                className="w-full text-xs font-semibold text-blue-500 p-1 hover:bg-gray-200 border border-gray-400 shadow-lg"
                onClick={() => props.onClick()}
            >Load more ...</button>
        </div>
    )
}