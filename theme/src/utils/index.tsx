// mathArticleName 匹配url的article名字
// 如/api/articles/test -> test
export function mathArticleName() {
    return window.location.pathname.split("/")[2];
}