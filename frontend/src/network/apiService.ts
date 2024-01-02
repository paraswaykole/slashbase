import Request from './request'
import { EventsEmit } from '../../wailsjs/runtime/runtime'
import Events from './constants'
import responseEvent from './responseEvent'
import { AxiosResponse } from 'axios'
import { ApiResult, User, UserSession, Project, DBConnection, DBDataModel, DBQueryData, CTIDResponse, DBQuery, DBQueryResult, DBQueryLog, Tab, PaginatedApiResult, Role, RolePermission, ProjectMember } from '../data/models'
import { AddDBConnPayload, AddProjectMemberPayload } from './payloads'
import { TabType } from '../data/defaults'
import Constants from '../constants'

const isapiService = (() => {
    if (Constants.Build === 'server') {
        return false
    }
    return true
})()

const getHealthCheck = async function (): Promise<any> {
    if (isapiService) {
        const response = responseEvent<any>(Events.HEALTH_CHECK.RESPONSE)
        EventsEmit(Events.HEALTH_CHECK.REQUEST, Events.HEALTH_CHECK.RESPONSE)
        return response
    }
    return await Request.apiInstance
        .get<ApiResult<undefined>>('/health')
        .then(res => res.data)
}

const loginUser = async function (email: string, password: string): Promise<ApiResult<UserSession>> {
    if (isapiService) {
        return Promise.reject("only api service")
    }
    return await Request.apiInstance
        .post<any, AxiosResponse<ApiResult<UserSession>>>('/user/login', { email, password })
        .then(res => res.data)
}

const isUserAuthenticated = async function (): Promise<boolean> {
    if (isapiService) {
        return Promise.reject("only api service")
    }
    return await Request.apiInstance
        .get<ApiResult<undefined>>('/user/checkauth')
        .then(res => res.data.success)
}

const logoutUser = async function (): Promise<ApiResult<null>> {
    if (isapiService) {
        return Promise.reject("only api service")
    }
    return await Request.apiInstance
        .get<ApiResult<null>>('/user/logout')
        .then(res => res.data)
}

const editUser = async function (name: string, profileImageUrl: string): Promise<ApiResult<User>> {
    if (isapiService) {
        return Promise.reject("only api service")
    }
    return await Request.apiInstance
        .post<any, AxiosResponse<ApiResult<User>>>('/user/edit', { name, profileImageUrl })
        .then(res => res.data)
}

const getUsers = async function (offset: number): Promise<PaginatedApiResult<User, number>> {
    if (isapiService) {
        return Promise.reject("only api service")
    }
    return await Request.apiInstance
        .get<PaginatedApiResult<User, number>>(`/user/all?offset=${offset}`)
        .then(res => res.data)
}

const searchUsers = async function (searchTerm: string, offset: number): Promise<PaginatedApiResult<User, number>> {
    if (isapiService) {
        return Promise.reject("only api service")
    }
    return await Request.apiInstance
        .get<PaginatedApiResult<User, number>>(`/user/all?offset=${offset}&search=${searchTerm}`)
        .then(res => res.data)
}

const addUsers = async function (email: string, password: string): Promise<ApiResult<undefined>> {
    if (isapiService) {
        return Promise.reject("only api service")
    }
    return await Request.apiInstance
        .post<any, AxiosResponse<ApiResult<undefined>>>(`/user/add`, { email, password })
        .then(res => res.data)
}

const createNewProject = async function (projectName: string): Promise<ApiResult<Project>> {
    if (isapiService) {
        const response = responseEvent<ApiResult<Project>>(Events.CREATE_PROJECT.RESPONSE)
        EventsEmit(Events.CREATE_PROJECT.REQUEST, Events.CREATE_PROJECT.RESPONSE, projectName)
        return response
    }
    return await Request.apiInstance
        .post<any, AxiosResponse<ApiResult<Project>>>('/project/create', { name: projectName })
        .then(res => res.data)
}

const deleteProject = async function (projectId: string): Promise<ApiResult<undefined>> {
    if (isapiService) {
        const response = responseEvent<ApiResult<undefined>>(Events.DELETE_PROJECT.RESPONSE)
        EventsEmit(Events.DELETE_PROJECT.REQUEST, Events.DELETE_PROJECT.RESPONSE, projectId)
        return response
    }
    return await Request.apiInstance
        .delete<any, AxiosResponse<ApiResult<undefined>>>(`/project/${projectId}`)
        .then(res => res.data)
}

const getProjects = async function (): Promise<ApiResult<Array<Project>>> {
    if (isapiService) {
        const response = responseEvent<ApiResult<Array<Project>>>(Events.GET_PROJECTS.RESPONSE)
        EventsEmit(Events.GET_PROJECTS.REQUEST, Events.GET_PROJECTS.RESPONSE)
        return response
    }
    return await Request.apiInstance
        .get<ApiResult<Array<Project>>>('/project/all')
        .then(res => res.data)
}

const addNewProjectMember = async function (projectId: string, payload: AddProjectMemberPayload): Promise<ApiResult<ProjectMember>> {
    if (isapiService) {
        return Promise.reject("only api service")
    }
    return await Request.apiInstance
        .post<any, AxiosResponse<ApiResult<ProjectMember>>>(`/project/${projectId}/members/create`, payload)
        .then(res => res.data)

}

const deleteProjectMember = async function (projectId: string, userId: string): Promise<ApiResult<undefined>> {
    if (isapiService) {
        return Promise.reject("only api service")
    }
    return await Request.apiInstance
        .delete<ApiResult<undefined>>(`/project/${projectId}/members/${userId}`)
        .then(res => res.data)

}

const getProjectMembers = async function (projectId: string): Promise<ApiResult<Array<ProjectMember>>> {
    if (isapiService) {
        return Promise.reject("only api service")
    }
    return await Request.apiInstance
        .get<ApiResult<Array<ProjectMember>>>(`/project/${projectId}/members`)
        .then(res => res.data)

}

const addNewDBConn = async function (dbConnPayload: AddDBConnPayload): Promise<ApiResult<DBConnection>> {
    if (isapiService) {
        const response = responseEvent<ApiResult<DBConnection>>(Events.CREATE_DBCONNECTION.RESPONSE)
        EventsEmit(Events.CREATE_DBCONNECTION.REQUEST, Events.CREATE_DBCONNECTION.RESPONSE, dbConnPayload)
        return response
    }
    return await Request.apiInstance
        .post<any, AxiosResponse<ApiResult<DBConnection>>>('/dbconnection/create', dbConnPayload)
        .then(res => res.data)
}

const getAllDBConnections = async function (): Promise<ApiResult<Array<DBConnection>>> {
    if (isapiService) {
        const response = responseEvent<ApiResult<Array<DBConnection>>>(Events.GET_DBCONNECTIONS.RESPONSE)
        EventsEmit(Events.GET_DBCONNECTIONS.REQUEST, Events.GET_DBCONNECTIONS.RESPONSE)
        return response
    }
    return await Request.apiInstance
        .get<ApiResult<Array<DBConnection>>>('/dbconnection/all')
        .then(res => res.data)
}

const getSingleDBConnection = async function (dbConnId: string): Promise<ApiResult<DBConnection>> {
    if (isapiService) {
        const response = responseEvent<ApiResult<DBConnection>>(Events.GETSINGLE_DBCONNECTION.RESPONSE)
        EventsEmit(Events.GETSINGLE_DBCONNECTION.REQUEST, Events.GETSINGLE_DBCONNECTION.RESPONSE, dbConnId)
        return response
    }
    return await Request.apiInstance
        .get<ApiResult<DBConnection>>(`/dbconnection/${dbConnId}`)
        .then(res => res.data)
}

const deleteDBConnection = async function (dbConnId: string): Promise<ApiResult<undefined>> {
    if (isapiService) {
        const response = responseEvent<ApiResult<undefined>>(Events.DELETE_DBCONNECTION.RESPONSE)
        EventsEmit(Events.DELETE_DBCONNECTION.REQUEST, Events.DELETE_DBCONNECTION.RESPONSE, dbConnId)
        return response
    }
    return await Request.apiInstance
        .delete<ApiResult<undefined>>(`/dbconnection/${dbConnId}`)
        .then(res => res.data)
}


const getDBConnectionsByProject = async function (projectId: string): Promise<ApiResult<Array<DBConnection>>> {
    if (isapiService) {
        const response = responseEvent<ApiResult<Array<DBConnection>>>(Events.GET_DBCONNECTIONS_BYPROJECT.RESPONSE)
        EventsEmit(Events.GET_DBCONNECTIONS_BYPROJECT.REQUEST, Events.GET_DBCONNECTIONS_BYPROJECT.RESPONSE, projectId)
        return response
    }
    return await Request.apiInstance
        .get<ApiResult<Array<DBConnection>>>(`/dbconnection/project/${projectId}`)
        .then(res => res.data)
}

const getDBDataModelsByConnectionId = async function (dbConnId: string): Promise<ApiResult<Array<DBDataModel>>> {
    if (isapiService) {
        const response = responseEvent<ApiResult<Array<DBDataModel>>>(Events.GET_DATAMODELS.RESPONSE)
        EventsEmit(Events.GET_DATAMODELS.REQUEST, Events.GET_DATAMODELS.RESPONSE, dbConnId)
        return response
    }
    return await Request.apiInstance
        .get<ApiResult<Array<DBDataModel>>>(`/query/datamodel/all/${dbConnId}`)
        .then(res => res.data)
}

const getDBSingleDataModelByConnectionId = async function (dbConnId: string, schemaName: string, mName: string): Promise<ApiResult<DBDataModel>> {
    if (isapiService) {
        const response = responseEvent<ApiResult<DBDataModel>>(Events.GETSINGLE_DATAMODEL.RESPONSE)
        EventsEmit(Events.GETSINGLE_DATAMODEL.REQUEST, Events.GETSINGLE_DATAMODEL.RESPONSE, dbConnId, schemaName, mName)
        return response
    }
    return await Request.apiInstance
        .get<ApiResult<DBDataModel>>(`/query/datamodel/single/${dbConnId}?schema=${schemaName}&name=${mName}`)
        .then(res => res.data)
}

const addDBSingleDataModelField = async function (dbConnId: string, schemaName: string, mName: string, fieldName: string, dataType: string): Promise<ApiResult<DBQueryResult>> {
    if (isapiService) {
        const response = responseEvent<ApiResult<DBQueryResult>>(Events.ADDSINGLE_DATAMODELFIELD.RESPONSE)
        EventsEmit(Events.ADDSINGLE_DATAMODELFIELD.REQUEST, Events.ADDSINGLE_DATAMODELFIELD.RESPONSE, { dbConnectionId: dbConnId, schema: schemaName, name: mName, fieldName, dataType })
        return response
    }
    return await Request.apiInstance
        .post<ApiResult<DBQueryResult>>(`/query/datamodel/single/addfield`, { dbConnectionId: dbConnId, schema: schemaName, name: mName, fieldName, dataType })
        .then(res => res.data)
}

const deleteDBSingleDataModelField = async function (dbConnId: string, schemaName: string, mName: string, fieldName: string): Promise<ApiResult<DBQueryResult>> {
    if (isapiService) {
        const response = responseEvent<ApiResult<DBQueryResult>>(Events.DELETESINGLE_DATAMODELFIELD.RESPONSE)
        EventsEmit(Events.DELETESINGLE_DATAMODELFIELD.REQUEST, Events.DELETESINGLE_DATAMODELFIELD.RESPONSE, { dbConnectionId: dbConnId, schema: schemaName, name: mName, fieldName })
        return response
    }
    return await Request.apiInstance
        .post<ApiResult<DBQueryResult>>(`/query/datamodel/single/deletefield`, { dbConnectionId: dbConnId, schema: schemaName, name: mName, fieldName })
        .then(res => res.data)
}

const addDBSingleDataModelIndex = async function (dbConnId: string, schemaName: string, mName: string, indexName: string, fieldNames: string[], isUnique: boolean): Promise<ApiResult<DBQueryResult>> {
    if (isapiService) {
        const response = responseEvent<ApiResult<DBQueryResult>>(Events.ADDSINGLE_DATAMODELINDEX.RESPONSE)
        EventsEmit(Events.ADDSINGLE_DATAMODELINDEX.REQUEST, Events.ADDSINGLE_DATAMODELINDEX.RESPONSE, { dbConnectionId: dbConnId, schema: schemaName, name: mName, indexName, fieldNames, isUnique })
        return response
    }
    return await Request.apiInstance
        .post<ApiResult<DBQueryResult>>(`/query/datamodel/single/addindex`, { dbConnectionId: dbConnId, schema: schemaName, name: mName, indexName, fieldNames, isUnique })
        .then(res => res.data)
}

const deleteDBSingleDataModelIndex = async function (dbConnId: string, schemaName: string, mName: string, indexName: string): Promise<ApiResult<DBQueryResult>> {
    if (isapiService) {
        const response = responseEvent<ApiResult<DBQueryResult>>(Events.DELETESINGLE_DATAMODELINDEX.RESPONSE)
        EventsEmit(Events.DELETESINGLE_DATAMODELINDEX.REQUEST, Events.DELETESINGLE_DATAMODELINDEX.RESPONSE, { dbConnectionId: dbConnId, schema: schemaName, name: mName, indexName })
        return response
    }
    return await Request.apiInstance
        .post<ApiResult<DBQueryResult>>(`/query/datamodel/single/deleteindex`, { dbConnectionId: dbConnId, schema: schemaName, name: mName, indexName })
        .then(res => res.data)
}

const getDBDataInDataModel = async function (dbConnId: string, schemaName: string, mName: string, limit: number, offset: number, isFirstFetch: boolean, filter?: string[], sort?: string[]): Promise<ApiResult<DBQueryData>> {
    if (isapiService) {
        const responseEventName = Events.GET_DATA.RESPONSE.replaceAll("[schema.name]", schemaName + "." + mName)
        const response = responseEvent<ApiResult<DBQueryData>>(responseEventName)
        EventsEmit(Events.GET_DATA.REQUEST, responseEventName, { dbConnectionId: dbConnId, schema: schemaName, name: mName, isFirstFetch, limit, offset, filter, sort })
        return response
    }
    return await Request.apiInstance
        .get<ApiResult<DBQueryData>>(`/query/data/${dbConnId}`, {
            params: {
                schema: schemaName,
                name: mName,
                limit: limit,
                offset: offset,
                count: isFirstFetch,
                filter,
                sort,
            }
        })
        .then(res => res.data)
}

const updateDBSingleData = async function (dbConnId: string, schemaName: string, mName: string, id: string, columnName: string, value: string): Promise<ApiResult<CTIDResponse>> {
    if (isapiService) {
        const response = responseEvent<ApiResult<CTIDResponse>>(Events.UPDATESINGLE_DATA.RESPONSE)
        EventsEmit(Events.UPDATESINGLE_DATA.REQUEST, Events.UPDATESINGLE_DATA.RESPONSE, { dbConnectionId: dbConnId, schema: schemaName, name: mName, id, columnName, value })
        return response
    }
    return await Request.apiInstance
        .post<any, AxiosResponse<ApiResult<CTIDResponse>>>(`/query/data/${dbConnId}/single`, { schema: schemaName, name: mName, id, columnName, value })
        .then(res => res.data)
}

const addDBData = async function (dbConnId: string, schemaName: string, mName: string, data: any): Promise<ApiResult<any>> {
    if (isapiService) {
        const response = responseEvent<ApiResult<any>>(Events.ADD_DATA.RESPONSE)
        EventsEmit(Events.ADD_DATA.REQUEST, Events.ADD_DATA.RESPONSE, { dbConnectionId: dbConnId, schema: schemaName, name: mName, data })
        return response
    }
    return await Request.apiInstance
        .post<any, AxiosResponse<ApiResult<any>>>(`/query/data/${dbConnId}/add`, { schema: schemaName, name: mName, data })
        .then(res => res.data)
}

const deleteDBData = async function (dbConnId: string, schemaName: string, mName: string, ids: string[]): Promise<ApiResult<DBQueryResult>> {
    if (isapiService) {
        const response = responseEvent<ApiResult<DBQueryResult>>(Events.DELETE_DATA.RESPONSE)
        EventsEmit(Events.DELETE_DATA.REQUEST, Events.DELETE_DATA.RESPONSE, { dbConnectionId: dbConnId, schema: schemaName, name: mName, ids })
        return response
    }
    return await Request.apiInstance
        .post<any, AxiosResponse<ApiResult<DBQueryResult>>>(`/query/data/${dbConnId}/delete`, { schema: schemaName, name: mName, ids })
        .then(res => res.data)
}

const saveDBQuery = async function (dbConnId: string, name: string, query: string, queryId: string): Promise<ApiResult<DBQuery>> {
    if (isapiService) {
        const response = responseEvent<ApiResult<DBQuery>>(Events.SAVE_DBQUERY.RESPONSE)
        EventsEmit(Events.SAVE_DBQUERY.REQUEST, Events.SAVE_DBQUERY.RESPONSE, { dbConnectionId: dbConnId, name, queryId, query })
        return response
    }
    return await Request.apiInstance
        .post<any, AxiosResponse<ApiResult<DBQuery>>>(`/query/save/${dbConnId}`, { name, queryId, query })
        .then(res => res.data)
}

const deleteDBQuery = async function (queryId: string): Promise<ApiResult<undefined>> {
    if (isapiService) {
        const response = responseEvent<ApiResult<undefined>>(Events.DELETE_DBQUERY.RESPONSE)
        EventsEmit(Events.DELETE_DBQUERY.REQUEST, Events.DELETE_DBQUERY.RESPONSE, queryId)
        return response
    }
    return await Request.apiInstance
        .delete<any, AxiosResponse<ApiResult<undefined>>>(`/query/delete/${queryId}`)
        .then(res => res.data)
}

const getDBQueriesInDBConn = async function (dbConnId: string): Promise<ApiResult<DBQuery[]>> {
    if (isapiService) {
        const response = responseEvent<ApiResult<DBQuery[]>>(Events.GET_DBQUERIES_INDBCONNECTION.RESPONSE)
        EventsEmit(Events.GET_DBQUERIES_INDBCONNECTION.REQUEST, Events.GET_DBQUERIES_INDBCONNECTION.RESPONSE, dbConnId)
        return response
    }
    return await Request.apiInstance
        .get<ApiResult<DBQuery[]>>(`/query/getall/${dbConnId}`)
        .then(res => res.data)
}

const getSingleDBQuery = async function (queryId: string): Promise<ApiResult<DBQuery>> {
    if (isapiService) {
        const response = responseEvent<ApiResult<DBQuery>>(Events.GETSINGLE_DBQUERY.RESPONSE)
        EventsEmit(Events.GETSINGLE_DBQUERY.REQUEST, Events.GETSINGLE_DBQUERY.RESPONSE, queryId)
        return response
    }
    return await Request.apiInstance
        .get<ApiResult<DBQuery>>(`/query/get/${queryId}`)
        .then(res => res.data)
}

const getDBHistory = async function (queryId: string, before?: number): Promise<PaginatedApiResult<DBQueryLog, number>> {
    if (isapiService) {
        const response = responseEvent<PaginatedApiResult<DBQueryLog, number>>(Events.GET_QUERYHISTORY_INDBCONNECTION.RESPONSE)
        EventsEmit(Events.GET_QUERYHISTORY_INDBCONNECTION.REQUEST, Events.GET_QUERYHISTORY_INDBCONNECTION.RESPONSE, queryId, before)
        return response
    }
    return await Request.apiInstance
        .get<PaginatedApiResult<DBQueryLog, number>>(`/query/history/${queryId}${before ? `?before=${before}` : ''}`)
        .then(res => res.data)
}

const runQuery = async function (dbConnId: string, query: string): Promise<ApiResult<DBQueryData | DBQueryResult>> {
    if (isapiService) {
        const response = responseEvent<ApiResult<DBQueryData | DBQueryResult>>(Events.RUN_QUERY.RESPONSE)
        EventsEmit(Events.RUN_QUERY.REQUEST, Events.RUN_QUERY.RESPONSE, dbConnId, query)
        return response
    }
    return await Request.apiInstance
        .post<any, AxiosResponse<ApiResult<DBQueryData | DBQueryResult>>>("/query/run", { dbConnectionId: dbConnId, query })
        .then(res => res.data)
}

const getSingleSetting = async function (name: string): Promise<ApiResult<any>> {
    if (isapiService) {
        const response = responseEvent<ApiResult<any>>(Events.GETSINGLE_SETTING.RESPONSE.replaceAll("[name]", name))
        EventsEmit(Events.GETSINGLE_SETTING.REQUEST, Events.GETSINGLE_SETTING.RESPONSE.replaceAll("[name]", name), name)
        return response
    }
    return await Request.apiInstance
        .get<any, AxiosResponse<ApiResult<any>>>(`/setting/single?name=${name}`)
        .then(res => res.data)
}

const updateSingleSetting = async function (name: string, value: string): Promise<ApiResult<undefined>> {
    if (isapiService) {
        const response = responseEvent<ApiResult<any>>(Events.UPDATESINGLE_SETTING.RESPONSE.replaceAll("[name]", name))
        EventsEmit(Events.UPDATESINGLE_SETTING.REQUEST, Events.UPDATESINGLE_SETTING.RESPONSE.replaceAll("[name]", name), name, value)
        return response
    }
    return await Request.apiInstance
        .post<any, AxiosResponse<ApiResult<any>>>(`/setting/single`, { name, value })
        .then(res => res.data)
}

const createTab = async function (dbConnectionId: string, tabType: string, mSchema: string, mName: string, queryId: string, query: string): Promise<ApiResult<Tab>> {
    if (isapiService) {
        const response = responseEvent<ApiResult<Tab>>(Events.CREATE_TAB.RESPONSE)
        EventsEmit(Events.CREATE_TAB.REQUEST, Events.CREATE_TAB.RESPONSE, dbConnectionId, tabType, mSchema, mName, queryId, query)
        return response
    }
    return await Request.apiInstance
        .post<any, AxiosResponse<ApiResult<Tab>>>(`/tab/create`, { dbConnectionId, tabType, modelschema: mSchema, modelname: mName, queryId, query })
        .then(res => res.data)
}

const getTabsByDBConnection = async function (dbConnectionId: string): Promise<ApiResult<Array<Tab>>> {
    if (isapiService) {
        const response = responseEvent<ApiResult<Array<Tab>>>(Events.GET_TABS_BYDBCONNECTION.RESPONSE)
        EventsEmit(Events.GET_TABS_BYDBCONNECTION.REQUEST, Events.GET_TABS_BYDBCONNECTION.RESPONSE, dbConnectionId)
        return response
    }
    return await Request.apiInstance
        .get<any, AxiosResponse<ApiResult<Array<Tab>>>>(`/tab/getall/${dbConnectionId}`)
        .then(res => res.data)
}

const updateTab = async function (dbConnectionId: string, tabId: string, tabType: TabType, metadata: Object): Promise<ApiResult<Tab>> {
    if (isapiService) {
        const response = responseEvent<ApiResult<Tab>>(Events.UPDATE_TAB.RESPONSE)
        EventsEmit(Events.UPDATE_TAB.REQUEST, Events.UPDATE_TAB.RESPONSE, dbConnectionId, tabId, tabType, metadata)
        return response
    }
    return await Request.apiInstance
        .post<any, AxiosResponse<ApiResult<Tab>>>(`/tab/update`, { dbConnectionId, tabId, tabType, metadata })
        .then(res => res.data)
}

const closeTab = async function (dbConnectionId: string, tabId: string): Promise<ApiResult<undefined>> {
    if (isapiService) {
        const response = responseEvent<ApiResult<undefined>>(Events.CLOSE_TAB.RESPONSE)
        EventsEmit(Events.CLOSE_TAB.REQUEST, Events.CLOSE_TAB.RESPONSE, dbConnectionId, tabId)
        return response
    }
    return await Request.apiInstance
        .delete<any, AxiosResponse<ApiResult<any>>>(`/tab/close/${dbConnectionId}/${tabId}`)
        .then(res => res.data)
}

const runConsoleCommand = async function (dbConnectionId: string, cmdString: string): Promise<ApiResult<string>> {
    if (isapiService) {
        const response = responseEvent<ApiResult<string>>(Events.CONSOLE_RUN_COMMAND.RESPONSE)
        EventsEmit(Events.CONSOLE_RUN_COMMAND.REQUEST, Events.CONSOLE_RUN_COMMAND.RESPONSE, dbConnectionId, cmdString)
        return response
    }
    return await Request.apiInstance
        .post<any, AxiosResponse<ApiResult<string>>>(`/console/runcmd`, { dbConnectionId, cmd: cmdString })
        .then(res => res.data)
}

const checkConnection = async function (dbConnectionId: string): Promise<ApiResult<undefined>> {
    if (isapiService) {
        const response = responseEvent<ApiResult<undefined>>(Events.CHECK_DBCONNECTION.RESPONSE.replaceAll("[dbid]", dbConnectionId))
        EventsEmit(Events.CHECK_DBCONNECTION.REQUEST, Events.CHECK_DBCONNECTION.RESPONSE.replaceAll("[dbid]", dbConnectionId), dbConnectionId)
        return response
    }
    return await Request.apiInstance
        .get<any, AxiosResponse<ApiResult<undefined>>>(`/dbconnection/check/${dbConnectionId}`)
        .then(res => res.data)
}

const getRoles = async function (): Promise<ApiResult<Role[]>> {
    if (isapiService) {
        return Promise.reject("only api service")
    }
    return await Request.apiInstance
        .get<any, AxiosResponse<ApiResult<Role[]>>>(`/role/all`)
        .then(res => res.data)
}

const addRole = async function (name: string): Promise<ApiResult<Role>> {
    if (isapiService) {
        return Promise.reject("only api service")
    }
    return await Request.apiInstance
        .post<any, AxiosResponse<ApiResult<Role>>>(`/role/add`, { name })
        .then(res => res.data)
}

const deleteRole = async function (roleId: string): Promise<ApiResult<Role>> {
    if (isapiService) {
        return Promise.reject("only api service")
    }
    return await Request.apiInstance
        .delete<any, AxiosResponse<ApiResult<Role>>>(`/role/${roleId}`)
        .then(res => res.data)
}

const updateRolePermission = async function (roleId: string, name: string, value: boolean): Promise<ApiResult<RolePermission>> {
    if (isapiService) {
        return Promise.reject("only api service")
    }
    return await Request.apiInstance
        .post<any, AxiosResponse<ApiResult<RolePermission>>>(`/role/${roleId}/permission`, { roleId, name, value })
        .then(res => res.data)
}

const generateSQL = async function (dbConnectionId: string, text: string): Promise<ApiResult<string>> {
    if (isapiService) {
        const response = responseEvent<ApiResult<string>>(Events.AI_GENSQL.RESPONSE)
        EventsEmit(Events.AI_GENSQL.REQUEST, Events.AI_GENSQL.RESPONSE, dbConnectionId, text)
        return response
    }
    return await Request.apiInstance
        .post<any, AxiosResponse<ApiResult<string>>>(`/ai/gensql`, { dbConnectionId, text })
        .then(res => res.data)
}

const listSupportedAIModels = async function (): Promise<ApiResult<string[]>> {
    if (isapiService) {
        const response = responseEvent<ApiResult<string[]>>(Events.AI_LIST_SUPPORTEDMODELS.RESPONSE)
        EventsEmit(Events.AI_LIST_SUPPORTEDMODELS.REQUEST, Events.AI_LIST_SUPPORTEDMODELS.RESPONSE)
        return response
    }
    return await Request.apiInstance
        .get<any, AxiosResponse<ApiResult<string[]>>>(`/ai/listmodels`)
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
    addUsers,
    getProjects,
    createNewProject,
    deleteProject,
    getProjectMembers,
    addNewProjectMember,
    deleteProjectMember,
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
    checkConnection,
    getRoles,
    addRole,
    deleteRole,
    updateRolePermission,
    generateSQL,
    listSupportedAIModels
}