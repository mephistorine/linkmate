export type AuthRegisterReqDto = {
    readonly email: string
    readonly name: string
    readonly password: string
}

export type AuthRegisterResDto = {
    readonly userId: number
}
