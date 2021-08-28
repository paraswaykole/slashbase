import Request from './request'
import { UserSession, ApiResult, Project, DBConnection, ProjectMember, DBDataModel, DBQueryData, User } from '../data/models'
import { AddDBConnPayload, AddProjectMemberPayload } from './payloads'

const loginUser = async function(email: string, password: string): Promise<ApiResult<UserSession>> {
    const result: ApiResult<UserSession> = await Request.apiInstance.post('/user/login', { email, password }).then(res => res.data)
    return result
}

const editUser = async function(name: string, profileImageUrl: string): Promise<ApiResult<User>> {
    const result: ApiResult<User> = await Request.apiInstance.post('/user/edit', { name, profileImageUrl }).then(res => res.data)
    return result
}

const createNewProject = async function(projectName: string): Promise<ApiResult<Project>> {
    const result: ApiResult<Project> = await Request.apiInstance.post('/project/create', {name: projectName}).then(res => res.data)
    return result
}

const getProjects = async function(): Promise<ApiResult<Array<Project>>> {
    const result: ApiResult<Array<Project>> = await Request.apiInstance.get('/project/all').then(res => res.data)
    return result
}

const addNewProjectMember = async function(projectId: string, payload: AddProjectMemberPayload): Promise<ApiResult<ProjectMember>> {
    const result: ApiResult<ProjectMember> = await Request.apiInstance.post(`/project/${projectId}/members/create`, payload).then(res => res.data)
    return result
}

const getProjectMembers = async function(projectId: string): Promise<ApiResult<Array<ProjectMember>>> {
    const result: ApiResult<Array<ProjectMember>> = await Request.apiInstance.get(`/project/${projectId}/members`).then(res => res.data)
    return result
}

const addNewDBConn = async function(dbConnPayload: AddDBConnPayload): Promise<ApiResult<DBConnection>> {
    const result: ApiResult<DBConnection> = await Request.apiInstance.post('/dbconnection/create', dbConnPayload).then(res => res.data)
    return result
}

const getAllDBConnections = async function(): Promise<ApiResult<Array<DBConnection>>> {
    const result: ApiResult<Array<DBConnection>> = await Request.apiInstance.get('/dbconnection/all').then(res => res.data)
    return result
}

const getSingleDBConnection = async function(dbConnId: string): Promise<ApiResult<DBConnection>> {
    const result: ApiResult<DBConnection> = await Request.apiInstance.get(`/dbconnection/${dbConnId}`).then(res => res.data)
    return result
}

const getDBConnectionsByProject = async function(projectId: string): Promise<ApiResult<Array<DBConnection>>> {
    const result: ApiResult<Array<DBConnection>> = await Request.apiInstance.get(`/dbconnection/project/${projectId}`).then(res => res.data)
    return result
}

const getDBDataModelsByConnectionId = async function(dbConnId: string): Promise<ApiResult<Array<DBDataModel>>> {
    const result: ApiResult<Array<DBDataModel>> = await Request.apiInstance.get(`/query/datamodels/${dbConnId}`).then(res => res.data)
    return result
}

const getDBDataInDataModel = async function(dbConnId: string,schemaName: string, mName: string, offset: number, fetchCount: boolean): Promise<ApiResult<DBQueryData>> {
    const result: ApiResult<DBQueryData> = await Request.apiInstance.get(`/query/data/${dbConnId}?schema=${schemaName}&name=${mName}&offset=${offset}&count=${fetchCount}`).then(res => res.data)
    return result
}

export default {
    loginUser,
    editUser,
    getProjects,
    createNewProject,
    getProjectMembers,
    getAllDBConnections,
    getSingleDBConnection,
    getDBConnectionsByProject,
    getDBDataModelsByConnectionId,
    getDBDataInDataModel,
    addNewDBConn,
    addNewProjectMember
}