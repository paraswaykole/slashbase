import { DBConnType } from "./defaults"

export interface User {
    id: string
    name: string | null
    email: string
    profileImageUrl: string
    isRoot: bool
    isOnboarded: bool
    createdAt: string
    updatedAt: string
}

export interface UserSession {
    id: string
    user: User
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
    role: Role
    user: User
    createdAt: string
    updatedAt: string
}

export interface Role {
    id: string
    name: string
    permissions?: RolePermission[]
    createdAt: string
    updatedAt: string
}

export interface RolePermission {
    id: string
    name: string
    value: bool
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

export interface Tab {
    id: string
    type: TabType
    metadata: {
        schema: string,
        name: string,
        queryId: string,
        query: string,
        queryName: string,
    }
    dbConnectionId: string
    isActive: bool
    createdAt: string
    updatedAt: string
}

export interface DBDataModel {
    name: string
    schemaName: string | null
    fields?: Array<{
        name: string
        type: string
        isPrimary: boolean
        isNullable: boolean
        tags: string[]
    }>
    indexes?: Array<{
        name: string
        indexDef: string
    }>
}

export interface DBQueryData {
    columns: string[]
    rows: any[]
    keys: string[]
    data: any[]
    count?: number
    pkeys?: string[]
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

export interface AddDataResponse {
    newId: string
    data?: any
}