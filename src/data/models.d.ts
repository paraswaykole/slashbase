import { DBConnType, ProjectMemberRole } from "./defaults"

export interface User{
    id: string
    name: string | null
    email: string
    profileImageUrl: string
    isRoot: bool
    createdAt: string
    updatedAt: string
}

export interface UserSession {
    id: string
    user: User
    token: string
    isActive: boolean
}

export interface Project {
    id: string
    name: string
    currentMember?: ProjectMember
    createdAt: string
    updatedAt: string
}

export interface ProjectMember {
    id: string
    role: ProjectMemberRole
    user: User
    createdAt: string
    updatedAt: string
}

export interface DBConnection {
    id: string
    name: string
    type: DBConnType
    projectId: string
    createdAt: string
    updatedAt: string
}

export interface DBDataModel {
    name: string
    schemaName: string|null
    fields?: Array<{
        name: string
        type: string
        isPrimary: boolean
        isNullable: boolean
        charMaxLength: number|null
        default: string|null
    }>
}

export interface DBQueryData {
    columns: string[]
    rows: any[]
    count?: number
}

export interface DBQueryResult {
    message: string
}

export interface DBQuery {
    id: string
    name: string
    query: string
    dbConnectionId: string
}

export interface DBQueryLog {
    id: string
    query: string
    user: User
    dbConnectionId: string
    createdAt: string
}

// Result Models

export interface ApiResult<T> { 
    data: T
    success: boolean
    error?: string
}

export interface PaginatedApiResult<T, N> {
    data: {
        list: T[]
        next: N
    }
    success: boolean
    error?: string
}

export interface CTIDResponse { 
    ctid: string
}