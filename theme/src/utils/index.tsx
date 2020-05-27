import api from "../api";

// mathArticleName 匹配url的article名字
// 如/api/articles/test -> test
export function mathArticleName() {
    return window.location.pathname.split("/")[2];
}

export function setStorage(key: string, value: string) {
    window.localStorage.setItem(key, value);
}

export function getStorage(key: string): string | null {
    return window.localStorage.getItem(key)
}

export function getAccessToken(): string {
    return getStorage("token.access_token") || "";
}

export class Auth {
    isLogin() {
        if (getAccessToken() === "") return false;
        return api.checkToken(getAccessToken())
            .then(res => true)
            .catch(res => false);
    }
}