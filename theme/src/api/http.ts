// 传输接口，实现http方法
export default interface Http {
    GET<T>(url: string): Promise<T>
    POST<T>(url: string, body: string): Promise<T>
    PATCH<T>(url: string, body: string): Promise<T>
    DELETE<T>(url: string): Promise<T>
}
