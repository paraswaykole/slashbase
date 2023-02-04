import { EventsEmit } from '../../wailsjs/runtime/runtime'
import { ApiResult, CTIDResponse, DBConnection, DBDataModel, DBQuery, DBQueryData, DBQueryLog, DBQueryResult, PaginatedApiResult, Project } from '../data/models'
import Events from './constants'
import { AddDBConnPayload } from './payloads'
import responseEvent from './responseEvent'

const getHealthCheck = async function (): Promise<any> {
    const response = responseEvent<any>(Events.HEALTH_CHECK.RESPONSE)
    EventsEmit(Events.HEALTH_CHECK.REQUEST, Events.HEALTH_CHECK.RESPONSE)
    return response
}

const createNewProject = async function (projectName: string): Promise<ApiResult<Project>> {
    const response = responseEvent<ApiResult<Project>>(Events.CREATE_PROJECT.RESPONSE)
    EventsEmit(Events.CREATE_PROJECT.REQUEST, Events.CREATE_PROJECT.RESPONSE, projectName)
    return response
}

const deleteProject = async function (projectId: string): Promise<ApiResult<undefined>> {
    const response = responseEvent<ApiResult<undefined>>(Events.DELETE_PROJECT.RESPONSE)
    EventsEmit(Events.DELETE_PROJECT.REQUEST, Events.DELETE_PROJECT.RESPONSE, projectId)
    return response
}

const getProjects = async function (): Promise<ApiResult<Array<Project>>> {
    const response = responseEvent<ApiResult<Array<Project>>>(Events.GET_PROJECTS.RESPONSE)
    EventsEmit(Events.GET_PROJECTS.REQUEST, Events.GET_PROJECTS.RESPONSE)
    return response
}

const addNewDBConn = async function (dbConnPayload: AddDBConnPayload): Promise<ApiResult<DBConnection>> {
    const response = responseEvent<ApiResult<DBConnection>>(Events.CREATE_DBCONNECTION.RESPONSE)
    EventsEmit(Events.CREATE_DBCONNECTION.REQUEST, Events.CREATE_DBCONNECTION.RESPONSE, dbConnPayload)
    return response
}

const getAllDBConnections = async function (): Promise<ApiResult<Array<DBConnection>>> {
    const response = responseEvent<ApiResult<Array<DBConnection>>>(Events.GET_DBCONNECTIONS.RESPONSE)
    EventsEmit(Events.GET_DBCONNECTIONS.REQUEST, Events.GET_DBCONNECTIONS.RESPONSE)
    return response
}

const getSingleDBConnection = async function (dbConnId: string): Promise<ApiResult<DBConnection>> {
    const response = responseEvent<ApiResult<DBConnection>>(Events.GETSINGLE_DBCONNECTION.RESPONSE)
    EventsEmit(Events.GETSINGLE_DBCONNECTION.REQUEST, Events.GETSINGLE_DBCONNECTION.RESPONSE, dbConnId)
    return response
}

const deleteDBConnection = async function (dbConnId: string): Promise<ApiResult<undefined>> {
    const response = responseEvent<ApiResult<undefined>>(Events.DELETE_DBCONNECTION.RESPONSE)
    EventsEmit(Events.DELETE_DBCONNECTION.REQUEST, Events.DELETE_DBCONNECTION.RESPONSE, dbConnId)
    return response
}

const getDBConnectionsByProject = async function (projectId: string): Promise<ApiResult<Array<DBConnection>>> {
    const response = responseEvent<ApiResult<Array<DBConnection>>>(Events.GET_DBCONNECTIONS_BYPROJECT.RESPONSE)
    EventsEmit(Events.GET_DBCONNECTIONS_BYPROJECT.REQUEST, Events.GET_DBCONNECTIONS_BYPROJECT.RESPONSE, projectId)
    return response
}

const getDBDataModelsByConnectionId = async function (dbConnId: string): Promise<ApiResult<Array<DBDataModel>>> {
    const response = responseEvent<ApiResult<Array<DBDataModel>>>(Events.GET_DATAMODELS.RESPONSE)
    EventsEmit(Events.GET_DATAMODELS.REQUEST, Events.GET_DATAMODELS.RESPONSE, dbConnId)
    return response
}

const getDBSingleDataModelByConnectionId = async function (dbConnId: string, schemaName: string, mName: string): Promise<ApiResult<DBDataModel>> {
    const response = responseEvent<ApiResult<DBDataModel>>(Events.GETSINGLE_DATAMODEL.RESPONSE)
    EventsEmit(Events.GETSINGLE_DATAMODEL.REQUEST, Events.GETSINGLE_DATAMODEL.RESPONSE, dbConnId, schemaName, mName)
    return response
}

const addDBSingleDataModelField = async function (dbConnId: string, schemaName: string, mName: string, fieldName: string, dataType: string): Promise<ApiResult<DBQueryResult>> {
    const response = responseEvent<ApiResult<DBQueryResult>>(Events.ADDSINGLE_DATAMODELFIELD.RESPONSE)
    EventsEmit(Events.ADDSINGLE_DATAMODELFIELD.REQUEST, Events.ADDSINGLE_DATAMODELFIELD.RESPONSE, { dbConnectionId: dbConnId, schema: schemaName, name: mName, fieldName, dataType })
    return response
}

const deleteDBSingleDataModelField = async function (dbConnId: string, schemaName: string, mName: string, fieldName: string): Promise<ApiResult<DBQueryResult>> {
    const response = responseEvent<ApiResult<DBQueryResult>>(Events.DELETESINGLE_DATAMODELFIELD.RESPONSE)
    EventsEmit(Events.DELETESINGLE_DATAMODELFIELD.REQUEST, Events.DELETESINGLE_DATAMODELFIELD.RESPONSE, { dbConnectionId: dbConnId, schema: schemaName, name: mName, fieldName })
    return response
}

const addDBSingleDataModelIndex = async function (dbConnId: string, schemaName: string, mName: string, indexName: string, fieldNames: string[], isUnique: boolean): Promise<ApiResult<DBQueryResult>> {
    const response = responseEvent<ApiResult<DBQueryResult>>(Events.ADDSINGLE_DATAMODELINDEX.RESPONSE)
    EventsEmit(Events.ADDSINGLE_DATAMODELINDEX.REQUEST, Events.ADDSINGLE_DATAMODELINDEX.RESPONSE, { dbConnectionId: dbConnId, schema: schemaName, name: mName, indexName, fieldNames, isUnique })
    return response
}

const deleteDBSingleDataModelIndex = async function (dbConnId: string, schemaName: string, mName: string, indexName: string): Promise<ApiResult<DBQueryResult>> {
    const response = responseEvent<ApiResult<DBQueryResult>>(Events.DELETESINGLE_DATAMODELINDEX.RESPONSE)
    EventsEmit(Events.DELETESINGLE_DATAMODELINDEX.REQUEST, Events.DELETESINGLE_DATAMODELINDEX.RESPONSE, { dbConnectionId: dbConnId, schema: schemaName, name: mName, indexName })
    return response
}

const getDBDataInDataModel = async function (dbConnId: string, schemaName: string, mName: string, limit: number, offset: number, fetchCount: boolean, filter?: string[], sort?: string[]): Promise<ApiResult<DBQueryData>> {
    const response = responseEvent<ApiResult<DBQueryData>>(Events.GET_DATA.RESPONSE)
    EventsEmit(Events.GET_DATA.REQUEST, Events.GET_DATA.RESPONSE, { dbConnectionId: dbConnId, schema: schemaName, name: mName, fetchCount, limit, offset, filter, sort })
    return response
}

const updateDBSingleData = async function (dbConnId: string, schemaName: string, mName: string, id: string, columnName: string, value: string): Promise<ApiResult<CTIDResponse>> {
    const response = responseEvent<ApiResult<CTIDResponse>>(Events.UPDATESINGLE_DATA.RESPONSE)
    EventsEmit(Events.UPDATESINGLE_DATA.REQUEST, Events.UPDATESINGLE_DATA.RESPONSE, { dbConnectionId: dbConnId, schema: schemaName, name: mName, id, columnName, value })
    return response
}

const addDBData = async function (dbConnId: string, schemaName: string, mName: string, data: any): Promise<ApiResult<any>> {
    const response = responseEvent<ApiResult<any>>(Events.ADD_DATA.RESPONSE)
    EventsEmit(Events.ADD_DATA.REQUEST, Events.ADD_DATA.RESPONSE, { dbConnectionId: dbConnId, schema: schemaName, name: mName, data })
    return response
}

const deleteDBData = async function (dbConnId: string, schemaName: string, mName: string, ids: string[]): Promise<ApiResult<DBQueryResult>> {
    const response = responseEvent<ApiResult<DBQueryResult>>(Events.DELETE_DATA.RESPONSE)
    EventsEmit(Events.DELETE_DATA.REQUEST, Events.DELETE_DATA.RESPONSE, { dbConnectionId: dbConnId, schema: schemaName, name: mName, ids })
    return response
}

const saveDBQuery = async function (dbConnId: string, name: string, query: string, queryId: string): Promise<ApiResult<DBQuery>> {
    const response = responseEvent<ApiResult<DBQuery>>(Events.SAVE_DBQUERY.RESPONSE)
    EventsEmit(Events.SAVE_DBQUERY.REQUEST, Events.SAVE_DBQUERY.RESPONSE, { dbConnectionId: dbConnId, name, queryId, query })
    return response
}

const deleteDBQuery = async function (queryId: string): Promise<ApiResult<undefined>> {
    const response = responseEvent<ApiResult<undefined>>(Events.DELETE_DBQUERY.RESPONSE)
    EventsEmit(Events.DELETE_DBQUERY.REQUEST, Events.DELETE_DBQUERY.RESPONSE, queryId)
    return response
}

const getDBQueriesInDBConn = async function (dbConnId: string): Promise<ApiResult<DBQuery[]>> {
    const response = responseEvent<ApiResult<DBQuery[]>>(Events.GET_DBQUERIES_INDBCONNECTION.RESPONSE)
    EventsEmit(Events.GET_DBQUERIES_INDBCONNECTION.REQUEST, Events.GET_DBQUERIES_INDBCONNECTION.RESPONSE, dbConnId)
    return response
}

const getSingleDBQuery = async function (queryId: string): Promise<ApiResult<DBQuery>> {
    const response = responseEvent<ApiResult<DBQuery>>(Events.GETSINGLE_DBQUERY.RESPONSE)
    EventsEmit(Events.GETSINGLE_DBQUERY.REQUEST, Events.GETSINGLE_DBQUERY.RESPONSE, queryId)
    return response
}

const getDBHistory = async function (queryId: string, before?: number): Promise<PaginatedApiResult<DBQueryLog, number>> {
    const response = responseEvent<PaginatedApiResult<DBQueryLog, number>>(Events.GET_QUERYHISTORY_INDBCONNECTION.RESPONSE)
    EventsEmit(Events.GET_QUERYHISTORY_INDBCONNECTION.REQUEST, Events.GET_QUERYHISTORY_INDBCONNECTION.RESPONSE, queryId, before)
    return response
}

const runQuery = async function (dbConnId: string, query: string): Promise<ApiResult<DBQueryData | DBQueryResult>> {
    const response = responseEvent<ApiResult<DBQueryData | DBQueryResult>>(Events.RUN_QUERY.RESPONSE)
    EventsEmit(Events.RUN_QUERY.REQUEST, Events.RUN_QUERY.RESPONSE, dbConnId, query)
    return response
}

const getSingleSetting = async function (name: string): Promise<ApiResult<any>> {
    const response = responseEvent<ApiResult<any>>(Events.GETSINGLE_SETTING.RESPONSE.replaceAll("[name]", name))
    EventsEmit(Events.GETSINGLE_SETTING.REQUEST, Events.GETSINGLE_SETTING.RESPONSE.replaceAll("[name]", name), name)
    return response
}

const updateSingleSetting = async function (name: string, value: string): Promise<ApiResult<undefined>> {
    const response = responseEvent<ApiResult<any>>(Events.UPDATESINGLE_SETTING.RESPONSE.replaceAll("[name]", name))
    EventsEmit(Events.UPDATESINGLE_SETTING.REQUEST, Events.UPDATESINGLE_SETTING.RESPONSE.replaceAll("[name]", name), name, value)
    return response
}

export default {
    getHealthCheck,
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
}