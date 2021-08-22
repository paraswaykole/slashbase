import { DBConnType, TeamMemberRole } from "./defaults"

export interface User{
    id: string
    name: string
    email: string
    profileImageUrl: string
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
    role: ProjectMemberRole|null
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
    rows: any[]
}

// Result Models

export interface ApiResult<T> { 
    data: T
    success: boolean
    error?: string
}