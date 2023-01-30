import { EventsEmit } from '../../wailsjs/runtime/runtime'
import { ApiResult, DBConnection, Project } from '../data/models'
import Events from './constants'
import { AddDBConnPayload } from './payloads'
import responseEvent from './responseEvent'

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
    getProjects,
    createNewProject,
    deleteProject,
    addNewDBConn,
    getAllDBConnections,
    getSingleDBConnection,
    deleteDBConnection,
    getDBConnectionsByProject,
    getSingleSetting,
    updateSingleSetting,
}