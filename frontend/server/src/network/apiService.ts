import Request from './request'
import { ApiResult, User, UserSession, Project, DBConnection, DBDataModel, DBQueryData, CTIDResponse, DBQuery, DBQueryResult, DBQueryLog, Tab, PaginatedApiResult } from '../data/models'
import { AddDBConnPayload } from './payloads'
import { AxiosResponse } from 'axios'
import { TabType } from '../data/defaults'

const getHealthCheck = async function (): Promise<any> {
    return await Request.apiInstance
        .get<ApiResult<undefined>>('/health')
        .then(res => res.data)
}

const loginUser = async function (email: string, password: string): Promise<ApiResult<UserSession>> {
    return await Request.apiInstance
        .post<any, AxiosResponse<ApiResult<UserSession>>>('/user/login', { email, password })
        .then(res => res.data)
}

const isUserAuthenticated = async function (): Promise<boolean> {
    return await Request.apiInstance
        .get<ApiResult<undefined>>('/user/checkauth')
        .then(res => res.data.success)
}

const logoutUser = async function (): Promise<ApiResult<null>> {
    return await Request.apiInstance
        .get<ApiResult<null>>('/user/logout')
        .then(res => res.data)
}

const editUser = async function (name: string, profileImageUrl: string): Promise<ApiResult<User>> {
    return await Request.apiInstance
        .post<any, AxiosResponse<ApiResult<User>>>('/user/edit', { name, profileImageUrl })
        .then(res => res.data)
}

const getUsers = async function (offset: number): Promise<PaginatedApiResult<User, number>> {
    return await Request.apiInstance
        .get<PaginatedApiResult<User, number>>(`/user/all?offset=${offset}`)
        .then(res => res.data)
}


const searchUsers = async function (searchTerm: string, offset: number): Promise<PaginatedApiResult<User, number>> {
    return await Request.apiInstance
        .get<PaginatedApiResult<User, number>>(`/user/all?offset=${offset}&search=${searchTerm}`)
        .then(res => res.data)
}

const createNewProject = async function (projectName: string): Promise<ApiResult<Project>> {
    return await Request.apiInstance
        .post<any, AxiosResponse<ApiResult<Project>>>('/project/create', { name: projectName })
        .then(res => res.data)
}

const deleteProject = async function (projectId: string): Promise<ApiResult<undefined>> {
    return await Request.apiInstance
        .delete<any, AxiosResponse<ApiResult<undefined>>>(`/project/${projectId}`)
        .then(res => res.data)
}

const getProjects = async function (): Promise<ApiResult<Array<Project>>> {
    return await Request.apiInstance
        .get<ApiResult<Array<Project>>>('/project/all')
        .then(res => res.data)
}

const addNewDBConn = async function (dbConnPayload: AddDBConnPayload): Promise<ApiResult<DBConnection>> {
    return await Request.apiInstance
        .post<any, AxiosResponse<ApiResult<DBConnection>>>('/dbconnection/create', dbConnPayload)
        .then(res => res.data)
}

const getAllDBConnections = async function (): Promise<ApiResult<Array<DBConnection>>> {
    return await Request.apiInstance
        .get<ApiResult<Array<DBConnection>>>('/dbconnection/all')
        .then(res => res.data)
}

const getSingleDBConnection = async function (dbConnId: string): Promise<ApiResult<DBConnection>> {
    return await Request.apiInstance
        .get<ApiResult<DBConnection>>(`/dbconnection/${dbConnId}`)
        .then(res => res.data)
}

const deleteDBConnection = async function (dbConnId: string): Promise<ApiResult<undefined>> {
    return await Request.apiInstance
        .delete<ApiResult<undefined>>(`/dbconnection/${dbConnId}`)
        .then(res => res.data)
}


const getDBConnectionsByProject = async function (projectId: string): Promise<ApiResult<Array<DBConnection>>> {
    return await Request.apiInstance
        .get<ApiResult<Array<DBConnection>>>(`/dbconnection/project/${projectId}`)
        .then(res => res.data)
}

const getDBDataModelsByConnectionId = async function (dbConnId: string): Promise<ApiResult<Array<DBDataModel>>> {
    return await Request.apiInstance
        .get<ApiResult<Array<DBDataModel>>>(`/query/datamodel/all/${dbConnId}`)
        .then(res => res.data)
}

const getDBSingleDataModelByConnectionId = async function (dbConnId: string, schemaName: string, mName: string): Promise<ApiResult<DBDataModel>> {
    return await Request.apiInstance
        .get<ApiResult<DBDataModel>>(`/query/datamodel/single/${dbConnId}?schema=${schemaName}&name=${mName}`)
        .then(res => res.data)
}

const addDBSingleDataModelField = async function (dbConnId: string, schemaName: string, mName: string, fieldName: string, dataType: string): Promise<ApiResult<DBQueryResult>> {
    return await Request.apiInstance
        .post<ApiResult<DBQueryResult>>(`/query/datamodel/single/addfield`, { dbConnectionId: dbConnId, schema: schemaName, name: mName, fieldName, dataType })
        .then(res => res.data)
}

const deleteDBSingleDataModelField = async function (dbConnId: string, schemaName: string, mName: string, fieldName: string): Promise<ApiResult<DBQueryResult>> {
    return await Request.apiInstance
        .post<ApiResult<DBQueryResult>>(`/query/datamodel/single/deletefield`, { dbConnectionId: dbConnId, schema: schemaName, name: mName, fieldName })
        .then(res => res.data)
}

const addDBSingleDataModelIndex = async function (dbConnId: string, schemaName: string, mName: string, indexName: string, fieldNames: string[], isUnique: boolean): Promise<ApiResult<DBQueryResult>> {
    return await Request.apiInstance
        .post<ApiResult<DBQueryResult>>(`/query/datamodel/single/addindex`, { dbConnectionId: dbConnId, schema: schemaName, name: mName, indexName, fieldNames, isUnique })
        .then(res => res.data)
}

const deleteDBSingleDataModelIndex = async function (dbConnId: string, schemaName: string, mName: string, indexName: string): Promise<ApiResult<DBQueryResult>> {
    return await Request.apiInstance
        .post<ApiResult<DBQueryResult>>(`/query/datamodel/single/deleteindex`, { dbConnectionId: dbConnId, schema: schemaName, name: mName, indexName })
        .then(res => res.data)
}

const getDBDataInDataModel = async function (dbConnId: string, schemaName: string, mName: string, limit: number, offset: number, fetchCount: boolean, filter?: string[], sort?: string[]): Promise<ApiResult<DBQueryData>> {
    return await Request.apiInstance
        .get<ApiResult<DBQueryData>>(`/query/data/${dbConnId}`, {
            params: {
                schema: schemaName,
                name: mName,
                limit: limit,
                offset: offset,
                count: fetchCount,
                filter,
                sort,
            }
        })
        .then(res => res.data)
}

const updateDBSingleData = async function (dbConnId: string, schemaName: string, mName: string, id: string, columnName: string, value: string): Promise<ApiResult<CTIDResponse>> {
    return await Request.apiInstance
        .post<any, AxiosResponse<ApiResult<CTIDResponse>>>(`/query/data/${dbConnId}/single`, { schema: schemaName, name: mName, id, columnName, value })
        .then(res => res.data)
}

const addDBData = async function (dbConnId: string, schemaName: string, mName: string, data: any): Promise<ApiResult<any>> {
    return await Request.apiInstance
        .post<any, AxiosResponse<ApiResult<any>>>(`/query/data/${dbConnId}/add`, { schema: schemaName, name: mName, data })
        .then(res => res.data)
}

const deleteDBData = async function (dbConnId: string, schemaName: string, mName: string, ids: string[]): Promise<ApiResult<DBQueryResult>> {
    return await Request.apiInstance
        .post<any, AxiosResponse<ApiResult<DBQueryResult>>>(`/query/data/${dbConnId}/delete`, { schema: schemaName, name: mName, ids })
        .then(res => res.data)
}

const saveDBQuery = async function (dbConnId: string, name: string, query: string, queryId: string): Promise<ApiResult<DBQuery>> {
    return await Request.apiInstance
        .post<any, AxiosResponse<ApiResult<DBQuery>>>(`/query/save/${dbConnId}`, { name, queryId, query })
        .then(res => res.data)
}

const deleteDBQuery = async function (queryId: string): Promise<ApiResult<undefined>> {
    return await Request.apiInstance
        .delete<any, AxiosResponse<ApiResult<undefined>>>(`/query/delete/${queryId}`)
        .then(res => res.data)
}

const getDBQueriesInDBConn = async function (dbConnId: string): Promise<ApiResult<DBQuery[]>> {
    return await Request.apiInstance
        .get<ApiResult<DBQuery[]>>(`/query/getall/${dbConnId}`)
        .then(res => res.data)
}

const getSingleDBQuery = async function (queryId: string): Promise<ApiResult<DBQuery>> {
    return await Request.apiInstance
        .get<ApiResult<DBQuery>>(`/query/get/${queryId}`)
        .then(res => res.data)
}

const getDBHistory = async function (queryId: string, before?: number): Promise<PaginatedApiResult<DBQueryLog, number>> {
    return await Request.apiInstance
        .get<PaginatedApiResult<DBQueryLog, number>>(`/query/history/${queryId}${before ? `?before=${before}` : ''}`)
        .then(res => res.data)
}

const runQuery = async function (dbConnId: string, query: string): Promise<ApiResult<DBQueryData | DBQueryResult>> {
    return await Request.apiInstance
        .post<any, AxiosResponse<ApiResult<DBQueryData | DBQueryResult>>>("/query/run", { dbConnectionId: dbConnId, query })
        .then(res => res.data)
}

const getSingleSetting = async function (name: string): Promise<ApiResult<any>> {
    return await Request.apiInstance
        .get<any, AxiosResponse<ApiResult<any>>>(`/setting/single?name=${name}`)
        .then(res => res.data)
}

const updateSingleSetting = async function (name: string, value: string): Promise<ApiResult<undefined>> {
    return await Request.apiInstance
        .post<any, AxiosResponse<ApiResult<any>>>(`/setting/single`, { name, value })
        .then(res => res.data)
}

const createTab = async function (dbConnectionId: string, tabType: string, mSchema: string, mName: string, queryId: string): Promise<ApiResult<Tab>> {
    return await Request.apiInstance
        .post<any, AxiosResponse<ApiResult<Tab>>>(`/tab/create`, { dbConnectionId, tabType, modelschema: mSchema, modelname: mName, queryId })
        .then(res => res.data)
}

const getTabsByDBConnection = async function (dbConnectionId: string): Promise<ApiResult<Array<Tab>>> {
    return await Request.apiInstance
        .get<any, AxiosResponse<ApiResult<Array<Tab>>>>(`/tab/getall/${dbConnectionId}`)
        .then(res => res.data)
}

const updateTab = async function (dbConnectionId: string, tabId: string, tabType: TabType, metadata: Object): Promise<ApiResult<Tab>> {
    return await Request.apiInstance
        .post<any, AxiosResponse<ApiResult<Tab>>>(`/tab/update`, { dbConnectionId, tabId, tabType, metadata })
        .then(res => res.data)
}

const closeTab = async function (dbConnectionId: string, tabId: string): Promise<ApiResult<undefined>> {
    return await Request.apiInstance
        .delete<any, AxiosResponse<ApiResult<any>>>(`/tab/close/${dbConnectionId}/${tabId}`)
        .then(res => res.data)
}

const runConsoleCommand = async function (dbConnectionId: string, cmdString: string): Promise<ApiResult<string>> {
    return await Request.apiInstance
        .post<any, AxiosResponse<ApiResult<string>>>(`/console/runcmd`, { dbConnectionId, cmd: cmdString })
        .then(res => res.data)
}

const checkConnection = async function (dbConnectionId: string): Promise<ApiResult<undefined>> {
    return await Request.apiInstance
        .get<any, AxiosResponse<ApiResult<undefined>>>(`/dbconnection/check/${dbConnectionId}`)
        .then(res => res.data)
}


export default {
    getHealthCheck,
    loginUser,
    isUserAuthenticated,
    logoutUser,
    editUser,
    getUsers,
    searchUsers,
    getProjects,
    createNewProject,
    deleteProject,
    getAllDBConnections,
    getSingleDBConnection,
    deleteDBConnection,
    getDBConnectionsByProject,
    getDBDataModelsByConnectionId,
    getDBSingleDataModelByConnectionId,
    addDBSingleDataModelField,
    deleteDBSingleDataModelField,
    addDBSingleDataModelIndex,
    deleteDBSingleDataModelIndex,
    getDBDataInDataModel,
    addNewDBConn,
    updateDBSingleData,
    addDBData,
    deleteDBData,
    saveDBQuery,
    deleteDBQuery,
    getDBQueriesInDBConn,
    getSingleDBQuery,
    getDBHistory,
    runQuery,
    getSingleSetting,
    updateSingleSetting,
    createTab,
    getTabsByDBConnection,
    updateTab,
    closeTab,
    runConsoleCommand,
    checkConnection
}