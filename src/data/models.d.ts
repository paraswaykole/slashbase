import { DBConnType, ProjectMemberRole } from "./defaults"

export interface User{
    id: string
    name: string
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
}

export interface DBQueryData {
    columns: string[],
    rows: any[],
    count?: number
}

// Result Models

export interface ApiResult<T> { 
    data: T
    success: boolean
    error?: string
}