import store from "./store"

// 异步获取vuex的登录状态
export async function localUserStatus(): Promise<{ islogin: boolean, secret: string }> {
    await store.dispatch("fetchUserInfoBycookie");
    if (getCookie() != "" && store.getters.islogin ) {
        return {
            islogin: true,
            secret: store.state.user.secret
        }
    }
    return { islogin: false, secret: "" }
}

export function setCookie(cookie: string) {
    const exp = new Date();
    exp.setTime(exp.getTime() + 0.1 * 24 * 60 * 60 * 1000);
    document.cookie = "user_session=" + cookie + ";expires=" + exp.toUTCString() + ";path=/";

}
export function getCookie(): string {
    const reg = new RegExp("(^| )user_session=([^;]*)(;|$)");
    let arr = document.cookie.match(new RegExp("(^| )user_session=([^;]*)(;|$)"))
    if (arr == null) return ''
    return arr[2];
}