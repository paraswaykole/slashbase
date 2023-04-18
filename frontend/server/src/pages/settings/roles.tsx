import React, { FunctionComponent, useEffect, useRef, useState } from 'react'
import toast from 'react-hot-toast'
import { useNavigate } from 'react-router-dom'
import ConfirmModal from '../../components/widgets/confirmModal'
import Constants from '../../constants'
import { Role } from '../../data/models'
import apiService from '../../network/apiService'
import { selectCurrentUser } from '../../redux/currentUserSlice'
import { useAppSelector } from '../../redux/hooks'

const ManageRolesPage: FunctionComponent<{}> = () => {

    const navigate = useNavigate()

    const [roles, setRoles] = useState<Role[]>([])
    const [isDeletingRole, setIsDeletingRole] = useState<Role | undefined>(undefined)
    const [showAddingRole, setShowAddingRole] = useState<boolean>(false)
    const [adding, setAdding] = useState<boolean>(false)

    const newRoleInputRef = useRef<HTMLInputElement>(null)

    const currentUser = useAppSelector(selectCurrentUser)
    useEffect(() => {
        if (currentUser && !currentUser.isRoot) {
            navigate(Constants.APP_PATHS.SETTINGS_ACCOUNT.path)
        }
    }, [currentUser])

    useEffect(() => {
        (async () => {
            const result = await apiService.getRoles()
            setRoles(result.data)
        })()
    }, [])

    const deleteRole = async () => {
        const result = await apiService.deleteRole(isDeletingRole!.id)
        if (result.success)
            setRoles(roles.filter(role => role.id !== isDeletingRole!.id))
        else
            toast.error(result.error!)
        setIsDeletingRole(undefined)
    }

    const addRole = async () => {
        if (newRoleInputRef.current!.value == '') {
            return
        }
        const result = await apiService.addRole(newRoleInputRef.current!.value)
        if (result.success)
            setRoles(roles.concat(result.data))
        else
            toast.error(result.error!)
        setAdding(false)
        setShowAddingRole(false)
    }


    const updateRolePermission = async (e: React.ChangeEvent<HTMLInputElement>) => {
        const roleId = e.target.dataset.roleid!
        const perName = e.target.dataset.name!
        const result = await apiService.updateRolePermission(roleId, perName, e.target.checked)
        if (result.success) {
            const newRoles = [...roles]
            const roleIndex = roles.findIndex(role => role.id === roleId)!
            const role = roles[roleIndex]
            const rp: number | undefined = role.permissions?.findIndex(rp => rp.name === perName)
            if (rp === undefined)
                role.permissions = [result.data, ...role.permissions ?? []]
            else
                role.permissions![rp] = result.data
            newRoles[roleIndex] = role
            setRoles(newRoles)
            toast.success("successfully updated role permission")
        }
        else
            toast.error(result.error!)
    }

    return (
        <React.Fragment>

            <h1>Manage Roles</h1>
            <br />
            {roles.length > 0 && <table className={"table is-bordered is-striped is-narrow is-hoverable"} style={{ minWidth: '200px' }}>
                <thead>
                    <tr>
                        <th colSpan={1} rowSpan={2}>Roles</th>
                        <th colSpan={1}>Permissions</th>
                        <th rowSpan={2}></th>
                    </tr>
                    <tr>
                        <th colSpan={1}>Read-only</th>
                    </tr>
                </thead>
                <tbody>
                    {
                        roles.map(role => (
                            <tr key={role.id}>
                                <td>{role.name}</td>
                                <td>
                                    <input
                                        data-roleid={role.id}
                                        data-name={Constants.ROLES_PERMISSIONS.READ_ONLY}
                                        checked={role.permissions ? role.permissions.find(rp => rp.name === Constants.ROLES_PERMISSIONS.READ_ONLY)?.value : false}
                                        onChange={updateRolePermission}
                                        type='checkbox' />
                                </td>
                                <td style={{ width: '54px' }}>
                                    <button className="button is-danger is-small" disabled={role.name === Constants.ROLES.ADMIN} onClick={() => { setIsDeletingRole(role) }}>
                                        <i className={"fas fa-trash"} />
                                    </button>
                                </td>
                            </tr>
                        ))
                    }
                </tbody>
            </table>}
            {showAddingRole ? <React.Fragment>
                <div className="form-container">
                    <div className="field">
                        <label className="label">New Role:</label>
                        <div className="control">
                            <input
                                ref={newRoleInputRef}
                                className="input"
                                type="text"
                                placeholder="Enter name for new role" />
                        </div>
                    </div>
                    <div className="control">
                        {!adding && <button className="button is-primary" onClick={addRole}>Add</button>}
                        {adding && <button className="button is-primary is-loading">Adding...</button>}
                        {!adding && <>&nbsp;<button className="button " onClick={() => { setShowAddingRole(false) }}>Cancel</button></>}
                    </div>
                </div>
            </React.Fragment> :
                <button className="button" onClick={() => { setShowAddingRole(true) }}>
                    <i className={"fas fa-plus"} />
                    &nbsp;&nbsp;
                    Add New Role
                </button>}
            {isDeletingRole && <ConfirmModal
                message={`Delete Role ${isDeletingRole.name}?`}
                onConfirm={deleteRole}
                onClose={() => { setIsDeletingRole(undefined) }} />}
        </React.Fragment>
    )
}

export default ManageRolesPage
