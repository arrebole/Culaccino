import React from 'react';

const headerStyle: React.CSSProperties = {
    display: "flex",
    alignItems: "center",
    height: "56px", 
    backgroundColor: "#1976d2",
    boxShadow: "0px 2px 4px -1px rgba(0,0,0,0.2), 0px 4px 5px 0px rgba(0,0,0,0.14), 0px 1px 10px 0px rgba(0,0,0,0.12)",
}

const Header: React.FC = () => {
    return (
        <header style={ headerStyle } > 
            <nav style={{ padding: "0px 20px" }}>
                <h1 style={{ margin: "0px", fontSize: "1.3rem" }}><a href="/" style={{color: "#fff", textDecoration: "none"}}>Culaccino</a></h1>
            </nav>
        </header>
    )
}

export default Header