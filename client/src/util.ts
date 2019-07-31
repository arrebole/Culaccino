export default {
    setCookie,
    getCookie,
}

export function setCookie(cookie: string) {
        var exp = new Date();
        exp.setTime(exp.getTime() + 0.1 * 24 * 60 * 60 * 1000);
        document.cookie = "user_session=" + cookie + ";expires=" + exp.toUTCString() + ";path=/";
 
}
export function getCookie():string {
    const reg = new RegExp("(^| )user_session=([^;]*)(;|$)");
    let arr = document.cookie.match(new RegExp("(^| )user_session=([^;]*)(;|$)"))
    if(arr == null) return ''
    return arr[2];
}