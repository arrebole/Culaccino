import { Route } from 'vue-router';
import { localUserStatus } from "../util"

// 没有登录不允许进入
export async function notAllowedNoLogin(to: Route, from: Route, next: any) {
        const status = await localUserStatus()
        if (!status.islogin) next({ name: "Login" })
        else next()
}

// 已经登录不允许进入
export async function notAllowedHadLogin(to: Route, from: Route, next: any) {
        const status = await localUserStatus()
        if (status.islogin) next({ name: "ManageRepos", params: { domain: status.secret } })
        else next()
}