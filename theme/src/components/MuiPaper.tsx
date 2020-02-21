import React from 'react';
import { Link } from "react-router-dom"

const MuiPaperStyle:React.CSSProperties = {
    padding: "16px",
    border:"1px solid #c4c4c4",
    boxShadow: "0px 2px 1px -1px rgba(0,0,0,0.2), 0px 1px 1px 0px rgba(0,0,0,0.14), 0px 1px 3px 0px rgba(0,0,0,0.12)",
}

export interface IProps{
    title: string,
    type: string,
    summary: string,
}

const MuiPaper: React.FC<IProps> = (props) => {
    return (
        <div style = { MuiPaperStyle }>
        <h2 style={{ margin:"0px" }}><Link style={{ fontSize: "1.25rem", color: "#1976d2", textDecoration: "none" }} to={props.title}>{ props.title }</Link></h2>
        <p style={{ fontSize: "0.9rem" }}>{ props.summary }</p>
        </div>
    )
}

export default MuiPaper