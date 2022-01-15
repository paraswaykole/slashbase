export enum DBConnType {
    POSTGRES = "POSTGRES"
}

export enum ProjectMemberRole {
    ADMIN = "ADMIN",
    DEVELOPER = "DEVELOPER",
    ANALYST = "ANALYST"
}

export enum DBConnectionUseSSHType {
    NONE        = "NONE",
	PASSWORD    = "PASSWORD",
	KEYFILE     = "KEYFILE",
	PASSKEYFILE = "PASSKEYFILE",
}

export enum DBConnectionLoginType {
    USE_ROOT        = "USE_ROOT",
	INDIVIDUAL_ACCOUNTS    = "INDIVIDUAL_ACCOUNTS",
}