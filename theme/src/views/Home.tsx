import React, { useState } from 'react';
import Header from '../components/Header'
import MuiPaper from '../components/MuiPaper'
import api, { DashboardData } from "../api"


function Main() {
    const [MuiPaperList, setMuiPaperList] = useState<DashboardData[]>([]);
    if( MuiPaperList.length === 0) api.getDashboard().then(res => { setMuiPaperList(res.data)})
    return (
        <article style={{ display: "flex", flexDirection: "column", alignItems: "center" }}>
            {MuiPaperList.map(item =>
                <section
                    style={{ width: "95%", maxWidth: "1200px", padding: "10px" }}
                    key={item.symbol}>
                    <MuiPaper {...item}></MuiPaper>
                </section>
            )}
        </article>
    )
}

export default class Home extends React.Component {
    render() {
        return (
            <div id="home"><Header /><Main /> </div>
        )
    }
}