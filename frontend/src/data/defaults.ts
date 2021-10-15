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