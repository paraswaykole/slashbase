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

export enum DBConnectionLoginType {
    USE_ROOT = "USE_ROOT",
    // ROLE_ACCOUNTS = "ROLE_ACCOUNTS",
}