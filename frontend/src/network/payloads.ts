

export interface AddDBConnPayload {
    projectId: string
    name: string
    type: string
    host: string
    port: string
    password: string
    user: string
    dbname: string
    useSSH: string
    sshHost: string
    sshUser: string
    sshPassword: string
    sshKeyFile: string
}

export interface AddProjectMemberPayload {
    email: string
    role: string
}