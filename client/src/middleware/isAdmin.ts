import { Route } from 'vue-router';
import util from '@/util';
import api from "../api/index"

export default{
    noPower,
    hadPower
}

export async function noPower(to: Route, from: Route, next: any) {
    if (util.getCookie() == '') next({ name: "Login" })
    else {
        let isAdmin = await api.token(util.getCookie());
        if (isAdmin.power != "admin") next({ name: "Login" })
        else next()
    }
}

export async function hadPower(to: Route, from: Route, next: any){
    let isAdmin = await api.token(util.getCookie());
    if(isAdmin.power == "admin") next({name:"Admin"})
    else next()
}