

export interface AddDBConnPayload {
    projectId: string
    name: string
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