interface EventType {
    [key: string]: {
        REQUEST: string
        RESPONSE: string
    }
}

const Events: EventType = {
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
    }
}

export default Events