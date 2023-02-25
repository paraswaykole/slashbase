interface EventType {
    [key: string]: {
        REQUEST: string
        RESPONSE: string
    }
}

const Events: EventType = {
    HEALTH_CHECK: {
        REQUEST: "event:check:health",
        RESPONSE: "response:check:health"
    },
    CREATE_PROJECT: {
        REQUEST: "event:create:project",
        RESPONSE: "response:create:project"
    },
    GET_PROJECTS: {
        REQUEST: "event:get:projects",
        RESPONSE: "response:get:projects"
    },
    DELETE_PROJECT: {
        REQUEST: "event:delete:project",
        RESPONSE: "response:delete:project"
    },
    CREATE_DBCONNECTION: {
        REQUEST: "event:create:dbconnection",
        RESPONSE: "response:create:dbconnection"
    },
    GET_DBCONNECTIONS: {
        REQUEST: "event:get:dbconnections",
        RESPONSE: "response:get:dbconnections"
    },
    DELETE_DBCONNECTION: {
        REQUEST: "event:delete:dbconnection",
        RESPONSE: "response:delete:dbconnection"
    },
    GETSINGLE_DBCONNECTION: {
        REQUEST: "event:getsingle:dbconnection",
        RESPONSE: "response:getsingle:dbconnection"
    },
    GET_DBCONNECTIONS_BYPROJECT: {
        REQUEST: "event:get:dbconnections:byproject",
        RESPONSE: "response:get:dbconnections:byproject"
    },
    GETSINGLE_SETTING: {
        REQUEST: "event:getsingle:setting",
        RESPONSE: "response:getsingle:setting:[name]"
    },
    UPDATESINGLE_SETTING: {
        REQUEST: "event:updatesingle:setting",
        RESPONSE: "response:updatesingle:setting:[name]"
    },
    RUN_QUERY: {
        REQUEST: "event:run:query",
        RESPONSE: "response:run:query"
    },
    GET_DATA: {
        REQUEST: "event:get:data",
        RESPONSE: "response:get:data"
    },
    GET_DATAMODELS: {
        REQUEST: "event:get:datamodels",
        RESPONSE: "response:get:datamodels",
    },
    GETSINGLE_DATAMODEL: {
        REQUEST: "event:getsingle:datamodel",
        RESPONSE: "response:getsingle:datamodel"
    },
    ADDSINGLE_DATAMODELFIELD: {
        REQUEST: "event:addsingle:datamodelfield",
        RESPONSE: "response:addsingle:datamodelfield"
    },
    DELETESINGLE_DATAMODELFIELD: {
        REQUEST: "event:deletesingle:datamodelfield",
        RESPONSE: "response:deletesingle:datamodelfield"
    },
    ADD_DATA: {
        REQUEST: "event:add:data",
        RESPONSE: "response:add:data"
    },
    DELETE_DATA: {
        REQUEST: "event:delete:data",
        RESPONSE: "response:delete:data"
    },
    UPDATESINGLE_DATA: {
        REQUEST: "event:updatesingle:data",
        RESPONSE: "response:updatesingle:data"
    },
    ADDSINGLE_DATAMODELINDEX: {
        REQUEST: "event:addsingle:datamodelindex",
        RESPONSE: "response:addsingle:datamodelindex"
    },
    DELETESINGLE_DATAMODELINDEX: {
        REQUEST: "event:deletesingle:datamodelindex",
        RESPONSE: "response:deletesingle:datamodelindex"
    },
    SAVE_DBQUERY: {
        REQUEST: "event:save:dbquery",
        RESPONSE: "response:save:dbquery"
    },
    DELETE_DBQUERY: {
        REQUEST: "event:delete:dbquery",
        RESPONSE: "response:delete:dbquery"
    },
    GET_DBQUERIES_INDBCONNECTION: {
        REQUEST: "event:get:dbqueries:indbconnection",
        RESPONSE: "response:get:dbqueries:indbconnection"
    },
    GETSINGLE_DBQUERY: {
        REQUEST: "event:getsingle:dbquery",
        RESPONSE: "response:getsingle:dbquery"
    },
    GET_QUERYHISTORY_INDBCONNECTION: {
        REQUEST: "event:get:queryhistory:indbconnection",
        RESPONSE: "response:get:queryhistory:indbconnection"
    },
    CREATE_TAB: {
        REQUEST: "event:create:tab",
        RESPONSE: "response:create:tab"
    },
    GET_TABS_BYDBCONNECTION: {
        REQUEST: "event:get:tabs:bydbconnection",
        RESPONSE: "response:get:tabs:bydbconnection"
    },
    UPDATE_TAB: {
        REQUEST: "event:update:tab",
        RESPONSE: "response:update:tab"
    },
    CLOSE_TAB: {
        REQUEST: "event:close:tab",
        RESPONSE: "response:close:tab"
    },
    CONSOLE_RUN_COMMAND: {
        REQUEST: "event:run:cmd",
        RESPONSE: "response:run:cmd"
    },
}

export default Events