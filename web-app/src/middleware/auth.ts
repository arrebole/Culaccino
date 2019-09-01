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
        let res = await api.sessionExist(util.getCookie());
        if (!res.data.islogin) next({ name: "Login" })
        else next()
    }
}

export async function hadPower(to: Route, from: Route, next: any){
    let res = await api.sessionExist(util.getCookie());
    if(res.data.islogin) next({name:"Admin"})
    else next()
}