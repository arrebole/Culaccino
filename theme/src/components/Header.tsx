import React from 'react';
import { Auth } from '../utils';

export default class Header extends React.Component {

    async gotoAdmin(e: React.MouseEvent<HTMLAnchorElement, MouseEvent>) {
        e.preventDefault();
        const isLogin = await new Auth().isLogin();
        if (isLogin) {
            return window.location.href = "/admin";
        }
        return window.location.href = "/login";
    }

    render() {
        return (
            <header>
                <nav className="bg-gray-900 p-3">
                    <a href="/admin" onClick={e => this.gotoAdmin(e)}>
                        <span className="iconfont icon-category text-white cursor-pointer" style={{ fontSize: "24px" }}></span>
                    </a>
                    <a href="/" className="font-sans font-semibold text-xl text-gray-200 mx-4 hover:text-gray-400">Culaccino</a>
                </nav>
            </header>
        )
    }
}