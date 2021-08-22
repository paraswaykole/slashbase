import Request from './request'
import { UserSession, ApiResult, Project, DBConnection, ProjectMember, DBDataModel } from '../data/models'


const loginUser = async function(email: string, password: string): Promise<ApiResult<UserSession>> {
    const result: ApiResult<UserSession> = await Request.apiInstance.post('/user/login', { email, password }).then(res => res.data)
    return result
}

const getProjects = async function(): Promise<ApiResult<Array<Project>>> {
    const result: ApiResult<Array<Project>> = await Request.apiInstance.get('/project/getall').then(res => res.data)
    return result
}

const getProjectMembers = async function(teamId: string): Promise<ApiResult<Array<ProjectMember>>> {
    const result: ApiResult<Array<ProjectMember>> = await Request.apiInstance.get(`/project/getmembers/${teamId}`).then(res => res.data)
    return result
}

const getAllDBConnections = async function(): Promise<ApiResult<Array<DBConnection>>> {
    const result: ApiResult<Array<DBConnection>> = await Request.apiInstance.get('/dbconnection/getall').then(res => res.data)
    return result
}

const getSingleDBConnection = async function(dbConnId: string): Promise<ApiResult<DBConnection>> {
    const result: ApiResult<DBConnection> = await Request.apiInstance.get(`/dbconnection/get/${dbConnId}`).then(res => res.data)
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

export default {
    loginUser,
    getProjects,
    getProjectMembers,
    getAllDBConnections,
    getSingleDBConnection,
    getDBConnectionsByProject,
    getDBDataModelsByConnectionId
}