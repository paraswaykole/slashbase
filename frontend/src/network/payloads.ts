

export interface AddDBConnPayload {
    projectId: string
    name: string
    host: string
    port: string
    password: string
    user: string
    dbname: string
    useSSH: string
    loginType: string
    sshHost: string
    sshUser: string
    sshPassword: string
    sshKeyFile: string
}

export interface AddProjectMemberPayload {   
    email: string
    role: string
}