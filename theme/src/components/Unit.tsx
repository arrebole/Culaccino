import React from 'react';
import { Paper } from "../api/index"
import { Link } from "react-router-dom"

const Unit: React.FC<Paper> = (props) => {
    return (
        <div className="d-flex m-0 p-4 box-shadow-large rounded-2 bg-white flex-column">
            <div className="text-align-center">
                <div className="my-1">
                    <Link className="h4 text-decoration-none text-blue" to={props.title}>{ props.title }</Link>
                </div>
                <div className="f6 d-flex flex-justify-center text-gray-light">
                    <div className="px-1" >Create on: </div> <time>{props.create_at.split(" ")[0]}</time>
                    <div className="px-1" >|</div> 
                    <div className="px-1" >Update on: </div> <time>{props.update_at.split(" ")[0]}</time>
                </div>
            </div>
            <div className="p-2">
                <img  style={{width:"100%"}} src={ props.cover } alt="img"/>
            </div>
            <p className="f4 text-gray text-align-center">{ props.summary }</p>
        </div>
    )
}

export default Unit