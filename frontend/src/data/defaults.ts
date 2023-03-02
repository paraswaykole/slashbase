export enum DBConnType {
    POSTGRES = "POSTGRES",
    MONGO = "MONGO",
    MYSQL = "MYSQL"
}

export enum DBConnectionUseSSHType {
    NONE = "NONE",
    PASSWORD = "PASSWORD",
    KEYFILE = "KEYFILE",
    PASSKEYFILE = "PASSKEYFILE",
}

export enum TabType {
    BLANK = "BLANK",
    DATA = "DATA",
    MODEL = "MODEL",
    QUERY = "QUERY",
    HISTORY = "HISTORY",
    CONSOLE = "CONSOLE"
}

export enum DBConnectionLoginType {
    USE_ROOT = "USE_ROOT",
    // ROLE_ACCOUNTS = "ROLE_ACCOUNTS",
}