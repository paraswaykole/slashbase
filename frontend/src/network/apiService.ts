import Request from './request'
import { UserSession, ApiResult, Project, DBConnection, ProjectMember, DBDataModel, DBQueryData, User, CTIDResponse, DBQuery, DBQueryResult, DBQueryLog, PaginatedApiResult } from '../data/models'
import { AddDBConnPayload, AddProjectMemberPayload } from './payloads'
import { AxiosResponse } from 'axios'

const loginUser = async function(email: string, password: string): Promise<ApiResult<UserSession>> {
    return await Request.getApiInstance()
                        .post<any, AxiosResponse<ApiResult<UserSession>>>('/user/login', { email, password })
                        .then(res => res.data)
}

const logoutUser = async function(): Promise<ApiResult<null>> {
    return await Request.getApiInstance()
                        .get<ApiResult<null>>('/user/logout')
                        .then(res => res.data)
}

const editUser = async function(name: string, profileImageUrl: string): Promise<ApiResult<User>> {
    return await Request.getApiInstance()
                        .post<any, AxiosResponse<ApiResult<User>>>('/user/edit', { name, profileImageUrl })
                        .then(res => res.data)
}

const changeUserPassword = async function(oldPassword: string, newPassword: string): Promise<ApiResult<undefined>> {
    return await Request.getApiInstance()
                        .post<any, AxiosResponse<ApiResult<undefined>>>('/user/password', { oldPassword, newPassword })
                        .then(res => res.data)
}

const getUsers = async function(offset: number): Promise<PaginatedApiResult<User, number>> {
    return await Request.getApiInstance()
                        .get<PaginatedApiResult<User, number>>(`/user/all?offset=${offset}`)
                        .then(res => res.data)
}


const searchUsers = async function(searchTerm: string,offset: number): Promise<PaginatedApiResult<User, number>> {
    return await Request.getApiInstance()
                        .get<PaginatedApiResult<User, number>>(`/user/all?offset=${offset}&search=${searchTerm}`)
                        .then(res => res.data)
}

const addUser = async function(email: string, password: string): Promise<ApiResult<undefined>> {
    return await Request.getApiInstance()
                        .post<any, AxiosResponse<ApiResult<undefined>>>(`/user/add`, {email, password})
                        .then(res => res.data)
}

const createNewProject = async function(projectName: string): Promise<ApiResult<Project>> {
    return await Request.getApiInstance()
                .post<any, AxiosResponse<ApiResult<Project>>>('/project/create', {name: projectName})
                .then(res => res.data)
}

const getProjects = async function(): Promise<ApiResult<Array<Project>>> {
    return await Request.getApiInstance()
                        .get<ApiResult<Array<Project>>>('/project/all')
                        .then(res => res.data)
}

const addNewProjectMember = async function(projectId: string, payload: AddProjectMemberPayload): Promise<ApiResult<ProjectMember>> {
    return await Request.getApiInstance()
                        .post<any, AxiosResponse<ApiResult<ProjectMember>>>(`/project/${projectId}/members/create`, payload)
                        .then(res => res.data)
    
}

const deleteProjectMember = async function(projectId: string, userId: string): Promise<ApiResult<undefined>> {
    return await Request.getApiInstance()
                        .delete<ApiResult<undefined>>(`/project/${projectId}/members/${userId}`)
                        .then(res => res.data)
    
}

const getProjectMembers = async function(projectId: string): Promise<ApiResult<Array<ProjectMember>>> {
    return await Request.getApiInstance()
                        .get<ApiResult<Array<ProjectMember>>>(`/project/${projectId}/members`)
                        .then(res => res.data)
    
}

const addNewDBConn = async function(dbConnPayload: AddDBConnPayload): Promise<ApiResult<DBConnection>> {
    return await Request.getApiInstance()
                        .post<any, AxiosResponse<ApiResult<DBConnection>>>('/dbconnection/create', dbConnPayload)
                        .then(res => res.data)
}

const getAllDBConnections = async function(): Promise<ApiResult<Array<DBConnection>>> {
    return await Request.getApiInstance()
                        .get<ApiResult<Array<DBConnection>>>('/dbconnection/all')
                        .then(res => res.data)
}

const getSingleDBConnection = async function(dbConnId: string): Promise<ApiResult<DBConnection>> {
    return await Request.getApiInstance()
                        .get<ApiResult<DBConnection>>(`/dbconnection/${dbConnId}`)
                        .then(res => res.data)
}

const deleteDBConnection = async function(dbConnId: string): Promise<ApiResult<undefined>> {
    return await Request.getApiInstance()
                        .delete<ApiResult<undefined>>(`/dbconnection/${dbConnId}`)
                        .then(res => res.data)
}


const getDBConnectionsByProject = async function(projectId: string): Promise<ApiResult<Array<DBConnection>>> {
    return await Request.getApiInstance()
                        .get<ApiResult<Array<DBConnection>>>(`/dbconnection/project/${projectId}`)
                        .then(res => res.data)
}

const getDBDataModelsByConnectionId = async function(dbConnId: string): Promise<ApiResult<Array<DBDataModel>>> {
    return await Request.getApiInstance()
                        .get<ApiResult<Array<DBDataModel>>>(`/query/datamodel/all/${dbConnId}`)
                        .then(res => res.data)
}

const getDBSingleDataModelByConnectionId = async function(dbConnId: string, schemaName: string, mName: string): Promise<ApiResult<DBDataModel>> {
    return await Request.getApiInstance()
                        .get<ApiResult<DBDataModel>>(`/query/datamodel/single/${dbConnId}?schema=${schemaName}&name=${mName}`)
                        .then(res => res.data)
}

const getDBDataInDataModel = async function(dbConnId: string, schemaName: string, mName: string, offset: number, fetchCount: boolean, filter?: string[], sort?: string[]): Promise<ApiResult<DBQueryData>> {
    return await Request.getApiInstance()
                        .get< ApiResult<DBQueryData>>(`/query/data/${dbConnId}`, { 
                            params: {
                                schema: schemaName,
                                name: mName,
                                offset: offset,
                                count: fetchCount,
                                filter,
                                sort,
                            }
                        })
                        .then(res => res.data)
}

const updateDBSingleData = async function(dbConnId: string, schemaName: string, mName: string, ctid: string, columnName: string, value: string): Promise<ApiResult<CTIDResponse>> {
    return await Request.getApiInstance()
                        .post<any, AxiosResponse<ApiResult<CTIDResponse>>>(`/query/data/${dbConnId}/single`, {schema: schemaName, name: mName, ctid, columnName, value})
                        .then(res => res.data)
}

const addDBData = async function(dbConnId: string, schemaName: string, mName: string, data: any): Promise<ApiResult<CTIDResponse>> {
    return await Request.getApiInstance()
                        .post<any, AxiosResponse<ApiResult<CTIDResponse>>>(`/query/data/${dbConnId}/add`, {schema: schemaName, name: mName, data})
                        .then(res => res.data)
}

const deleteDBData = async function(dbConnId: string, schemaName: string, mName: string, ctids: string[]): Promise<ApiResult<DBQueryResult>> {
    return await Request.getApiInstance()
                        .post<any, AxiosResponse<ApiResult<DBQueryResult>>>(`/query/data/${dbConnId}/delete`, {schema: schemaName, name: mName, ctids})
                        .then(res => res.data)
}

const saveDBQuery = async function(dbConnId: string, name: string, query: string, queryId: string): Promise<ApiResult<DBQuery>> {
    return await Request.getApiInstance()
                        .post<any, AxiosResponse<ApiResult<DBQuery>>>(`/query/save/${dbConnId}`, {name, queryId, query})
                        .then(res => res.data)
}

const getDBQueriesInDBConn = async function(dbConnId: string): Promise<ApiResult<DBQuery[]>> {
    return await Request.getApiInstance()
                        .get<ApiResult<DBQuery[]>>(`/query/getall/${dbConnId}`)
                        .then(res => res.data)
}

const getSingleDBQuery = async function(queryId: string): Promise<ApiResult<DBQuery>> {
    return await Request.getApiInstance()
                        .get<ApiResult<DBQuery>>(`/query/get/${queryId}`)
                        .then(res => res.data)
}

const getDBHistory = async function(queryId: string, before?: number): Promise<PaginatedApiResult<DBQueryLog, number>> {
    return await Request.getApiInstance()
                        .get<PaginatedApiResult<DBQueryLog, number>>(`/query/history/${queryId}${before?`?before=${before}`:''}`)
                        .then(res => res.data)
}

const runQuery = async function(dbConnId: string, query: string): Promise<ApiResult<DBQueryData|DBQueryResult>> {
    return await Request.getApiInstance()
                        .post<any, AxiosResponse<ApiResult<DBQueryData|DBQueryResult>>>("/query/run", {dbConnectionId: dbConnId, query})
                        .then(res => res.data)
}

export default {
    loginUser,
    logoutUser,
    editUser,
    changeUserPassword,
    getUsers,
    searchUsers,
    addUser,
    getProjects,
    createNewProject,
    getProjectMembers,
    getAllDBConnections,
    getSingleDBConnection,
    deleteDBConnection,
    getDBConnectionsByProject,
    getDBDataModelsByConnectionId,
    getDBSingleDataModelByConnectionId,
    getDBDataInDataModel,
    addNewDBConn,
    addNewProjectMember,
    deleteProjectMember,
    updateDBSingleData,
    addDBData,
    deleteDBData,
    saveDBQuery,
    getDBQueriesInDBConn,
    getSingleDBQuery,
    getDBHistory,
    runQuery
}