import React, { useState } from 'react';
import Header from '../components/Header'
import MuiPaper from '../components/MuiPaper'
import api, { DashboardData } from "../api"


function Main() {
    const [MuiPaperList, setMuiPaperList] = useState<DashboardData[]>([]);
    if( MuiPaperList.length === 0) api.getDashboard().then(res => { setMuiPaperList(res.data)})
    return (
        <main style={{ display: "flex", flexDirection: "column", alignItems: "center" }}>
            {MuiPaperList.map(item =>
                <article style={{ width: "95%", maxWidth: "900px", padding: "10px" }} key={item.title}>
                    <MuiPaper {...item}></MuiPaper>
                </article>
            )}
        </main>
    )
}

export default class Home extends React.Component {
    render() {
        return (
            <div id="home" style={{backgroundColor: "#eee"}}><Header /><Main /> </div>
        )
    }
}