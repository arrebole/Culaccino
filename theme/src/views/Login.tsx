import React from "react";
import Header from "../components/Header";
import api from "../api";
import * as types from "../types";
import { setStorage } from "../utils";

interface State {
    username: string
    password: string
}

export default class Login extends React.Component<{}, State> {
    constructor(props: {}) {
        super(props)
        this.state = {
            username: "",
            password: "",
        }
    }

    login(e: React.FormEvent<HTMLFormElement>) {
        e.preventDefault();
        api.login(this.state)
            .then(res => this.success(res))
            .catch(err => alert("用户名或密码错误"));
    }

    success(token: types.Token) {
        setStorage("token.access_token", `${token.token_type} ${token.access_token}`);
        setStorage("token.expires_at", token.expires_at);
        window.location.href = "/";
    }

    handleGetInputValue(key: "username" | "password", value: string) {
        switch (key) {
            case "username":
                this.setState({ username: value });
                break;
            case "password":
                this.setState({ password: value });
                break
        }
    }

    render() {
        return (
            <div>
                <Header />
                <main className="bg-gray-200 p-6">

                    <div className="p-4 text-center">
                        <h1 className="text-2xl text-gray-600">Sign in to Culaccino</h1>
                    </div>

                    <form
                        className="bg-white w-1/6 mx-auto border border-gray-500 rounded p-3"
                        style={{ minWidth: '300px' }}
                        onSubmit={e => this.login(e)}
                    >
                        <div className="p-2">
                            <label htmlFor="username">UserName</label>
                            <input
                                id="username"
                                className="w-full border border-gray-500 focus:border-blue-500 rounded p-1"
                                value={this.state.username}
                                onChange={e => this.handleGetInputValue("username", e.target.value)}
                            />
                        </div>

                        <div className="p-2">
                            <label htmlFor="password">PassWord</label>
                            <input
                                id="password"
                                className="w-full border border-gray-500 focus:border-blue-500 rounded p-1" type="password"
                                value={this.state.password}
                                onChange={e => this.handleGetInputValue("password", e.target.value)}
                            />
                        </div>

                        <div className="p-2">
                            <button className="w-full mx-auto p-1 bg-green-500 text-white border-2 border-green-500 hover:text-red-200 rounded">
                                Sign in
                            </button>
                        </div>
                    </form>

                </main>
            </div>
        )
    }
}