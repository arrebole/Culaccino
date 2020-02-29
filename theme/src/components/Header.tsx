import React from 'react';

const Header: React.FC = () => {
    return (
        <header className="box-shadow bg-blue p-3" > 
            <nav>
                <h1 className="m-0 f3"> <a className="text-white text-decoration-none" href="/">Culaccino</a></h1>
            </nav>
        </header>
    )
}

export default Header