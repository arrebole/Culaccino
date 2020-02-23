import React from 'react';
import {DashboardData} from "../api/index"
import { Link } from "react-router-dom"

const MuiPaperStyle:React.CSSProperties = {
    display: "flex",
    flexDirection:"column",
    alignItems: "center",
    padding: "20px",
    margin: "0px 0px",
    border:"1px solid #c4c4c4",
    backgroundColor:"#fff",
    boxShadow: "0 2px 2px 0 rgba(0,0,0,0.12), 0 3px 1px -2px rgba(0,0,0,0.06), 0 1px 5px 0 rgba(0,0,0,0.12)",
}

const MuiPaperHeader:React.CSSProperties = {
    textAlign: "center",
}

const MuiPaper: React.FC<DashboardData> = (props) => {
    return (
        <div style = { MuiPaperStyle }>
            <div style={ MuiPaperHeader }>
                <div style={{ margin:"5px 0px" }}>
                    <Link style={{ fontSize: "1.4rem", color: "#1976d2", textDecoration: "none" }} to={props.title}>{ props.title }</Link>
                </div>
                <div style={ {color:"rgb(151, 151, 151)", fontSize:"12px", display:"flex", justifyContent:"center"} }>
                    <div style={{padding:"0px 5px"}}>Create on: </div> <time>{props.create_at.split(" ")[0]}</time>
                    <div style={{padding:"0px 5px"}}>|</div> 
                    <div style={{padding:"0px 5px"}}>Update on: </div> <time>{props.create_at.split(" ")[0]}</time>
                </div>
            </div>
            <div style={{padding:"10px"}}>
                <img  style={{width:"100%"}} src={ props.cover } alt="img"/>
            </div>
            <p style={{ fontSize: "0.9rem"}}>{ props.summary }</p>
        </div>
    )
}

export default MuiPaper