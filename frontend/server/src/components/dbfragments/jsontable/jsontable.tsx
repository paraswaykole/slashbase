import styles from './jsontable.module.scss'
import React, { useContext, useRef, useState } from 'react'
import { useRowSelect, useTable } from 'react-table'
import { DBConnection, DBQueryData, Tab } from '../../../data/models'
import JsonCell from './jsoncell'
import AddModal from './addmodel'
import toast from 'react-hot-toast'
import ConfirmModal from '../../widgets/confirmModal'
import { useAppDispatch } from '../../../redux/hooks'
import { deleteDBData, setQueryData, updateDBSingleData } from '../../../redux/dataModelSlice'
import TabContext from '../../layouts/tabcontext'

type JsonTablePropType = {
    queryData: DBQueryData,
    dbConnection: DBConnection
    mName: string,
    isInteractive: boolean,
    isReadOnly: boolean,
    showHeader?: boolean,
    onRefresh: () => void,
    onFilterChanged: (newFilter: string[] | undefined) => void,
    onSortChanged: (newSort: string[] | undefined) => void,
}

const JsonTable = ({ queryData, dbConnection, mName, isInteractive, isReadOnly, showHeader, onFilterChanged, onSortChanged, onRefresh }: JsonTablePropType) => {

    const dispatch = useAppDispatch()

    const activeTab: Tab = useContext(TabContext)!

    const [isAdding, setIsAdding] = useState<boolean>(false)
    const [isEditing, setIsEditing] = useState<boolean>(false)
    const [isDeleting, setIsDeleting] = useState<boolean>(false)
    const [editingCellIndex, setEditingCellIndex] = useState<(number | null)>(null)

    const filterRef = useRef<HTMLInputElement>(null);
    const sortRef = useRef<HTMLInputElement>(null);

    const data = React.useMemo(
        () => queryData.data,
        [queryData]
    )

    const displayColumns = ["data"]

    const columns = React.useMemo(
        () => displayColumns.map((col, i) => ({
            Header: <>{col}</>,
            accessor: (i).toString(),
        })),
        [queryData]
    )

    const defaultColumn = {
        Cell: JsonCell,
    }

    const startEditing = (index: number | null) => {
        if (!(isInteractive && isEditing)) {
            return
        }
        setEditingCellIndex(index)
    }

    const changeFilter = () => {
        let filter: string[] | undefined = undefined
        const filterText = filterRef.current!.value.trim()
        if (filterText !== '' && filterText.startsWith("{") && filterText.endsWith("}")) {
            filter = [filterText]
        }
        onFilterChanged(filter)
    }

    const changeSort = () => {
        if (!isInteractive) {
            return
        }
        let sort: string[] | undefined = undefined
        const sortText = sortRef.current!.value.trim()
        if (sortText !== '' && sortText.startsWith("{") && sortText.endsWith("}")) {
            sort = [sortText]
        }
        onSortChanged(sort)
    }


    const onSaveCell = async (underscoreId: string, newData: string) => {
        const result = await dispatch(updateDBSingleData({ dbConnectionId: dbConnection.id, schemaName: "", name: mName, id: underscoreId, columnName: "", newValue: newData })).unwrap()
        if (result.success) {
            const rowIdx = queryData!.data.findIndex(x => x["_id"] == underscoreId)
            if (rowIdx) {
                const newQueryData: DBQueryData = { ...queryData!, data: [...queryData!.data] }
                newQueryData!.data[rowIdx] = { _id: underscoreId, ...JSON.parse(newData) }
                dispatch(setQueryData({ data: newQueryData, tabId: activeTab.id }))
            } else {
                // fetchData(false)
            }
            startEditing(null)
            toast.success('1 row updated');
        } else {
            toast.error(result.error!);
        }
    }

    const {
        getTableProps,
        getTableBodyProps,
        headerGroups,
        rows,
        prepareRow,
        state,
    } = useTable({
        columns,
        data,
        defaultColumn,
        initialState: { selectedRowIds: {} },
        ...{ editingCellIndex, startEditing, onSaveCell }
    }, useRowSelect, hooks => {
        if (isInteractive && isEditing)
            hooks.visibleColumns.push(columns => [
                {
                    id: 'selection',
                    Header: HeaderSelectionComponent,
                    Cell: CellSelectionComponent,
                },
                ...columns,
            ]
            )
    })

    const newState: any = state // temporary typescript hack
    const selectedRows: number[] = Object.keys(newState.selectedRowIds).map(x => parseInt(x))
    const selectedUnderscoreIDs = rows.filter((_, i) => selectedRows.includes(i)).map(x => (x.original as any)['_id']).filter(x => x)

    const deleteRows = async () => {
        if (selectedUnderscoreIDs.length > 0) {
            const result = await dispatch(deleteDBData({ dbConnectionId: dbConnection.id, schemaName: "", name: mName, selectedIDs: selectedUnderscoreIDs })).unwrap()
            if (result.success) {
                toast.success('rows deleted');
                const filteredRows = queryData!.data.filter((_, i) => !selectedRows.includes(i))
                const newQueryData: DBQueryData = { ...queryData!, data: filteredRows }
                dispatch(setQueryData({ data: newQueryData, tabId: activeTab.id }))
            } else {
                toast.error(result.error!);
            }
        }
        setIsDeleting(false)
    }

    return (
        <React.Fragment>
            {(showHeader || (isInteractive && isEditing)) && <div className={styles.tableHeader}>
                <div className="columns">
                    <div className="column is-3">
                        <div className="field has-addons">
                            <p className="control">
                                <input ref={filterRef} className="input" type="text" placeholder="{ field: 'Value'}" />
                            </p>
                            <p className="control">
                                <button className="button" onClick={changeFilter}>Filter</button>
                            </p>
                        </div>
                    </div>
                    <div className="column is-6">
                        <div className="field has-addons">
                            <p className="control">
                                <input ref={sortRef} className="input" type="text" placeholder="{ field: 1 or -1}" />
                            </p>
                            <p className="control">
                                <button className="button" onClick={changeSort}>Sort</button>
                            </p>
                        </div>
                    </div>
                    {isInteractive && !isEditing && <React.Fragment>
                        <div className="column is-3 is-flex is-justify-content-flex-end">
                            <button className="button is-secondary" onClick={onRefresh}>
                                <span className="icon is-small">
                                    <i className="fas fa-refresh" />
                                </span>
                            </button>
                            &nbsp;&nbsp;
                            {!isReadOnly && <button className="button is-primary" onClick={() => { setIsEditing(true) }}>
                                <span className="icon is-small">
                                    <i className="fas fa-pen" />
                                </span>
                            </button>}
                        </div>
                    </React.Fragment>}
                    {isInteractive && isEditing && <React.Fragment>
                        <div className="column is-3 is-flex is-justify-content-flex-end">
                            <button className="button" disabled={selectedUnderscoreIDs.length === 0} onClick={() => { setIsDeleting(true) }}>
                                <span className="icon is-small">
                                    <i className="fas fa-trash" />
                                </span>
                            </button>
                            &nbsp;&nbsp;
                            <button className="button is-secondary" onClick={() => { setIsAdding(true) }}>
                                <span className="icon is-small">
                                    <i className="fas fa-plus" />
                                </span>
                            </button>
                            &nbsp;&nbsp;
                            <button className="button is-primary" onClick={() => { setIsEditing(false) }}>
                                <span className="icon is-small">
                                    <i className="fas fa-check" />
                                </span>
                            </button>
                        </div>
                    </React.Fragment>}
                </div>
            </div>}
            {isAdding &&
                <AddModal
                    dbConnection={dbConnection}
                    mName={mName}
                    onClose={() => { setIsAdding(false) }} />
            }
            {isDeleting && <ConfirmModal
                message={`Are you sure you want to delete selected documents?`}
                onConfirm={deleteRows}
                onClose={() => { setIsDeleting(false) }} />}
            <div className="table-container">
                <table {...getTableProps()} className={"table is-bordered is-striped is-narrow is-hoverable is-fullwidth"}>
                    <thead>
                        {headerGroups.map(headerGroup => (
                            <tr {...headerGroup.getHeaderGroupProps()} key={"header"}>
                                {headerGroup.headers.map(column => (
                                    <th {...column.getHeaderProps()} key={column.id}>
                                        {column.render('Header')}
                                    </th>
                                ))}
                            </tr>
                        ))}
                    </thead>
                    <tbody {...getTableBodyProps()}>
                        {rows.map((row, rowIndex) => {
                            prepareRow(row)
                            const selectedRow: any = row // temp type hack 
                            return (
                                <tr {...row.getRowProps()} key={row.id} className={selectedRow.isSelected ? 'is-selected' : ''}>
                                    {row.cells.map(cell => {
                                        return (<td {...cell.getCellProps()} onDoubleClick={() => startEditing(rowIndex)} key={row.id + "" + cell.column.id}>
                                            {cell.render('Cell')}
                                        </td>
                                        )
                                    })}
                                </tr>
                            )
                        })}
                    </tbody>
                </table>
            </div>
        </React.Fragment>
    )
}

const IndeterminateCheckbox = React.forwardRef<HTMLInputElement, { indeterminate: boolean }>(
    ({ indeterminate, ...rest }, ref) => {
        const defaultRef = React.useRef()
        const resolvedRef: any = ref || defaultRef

        React.useEffect(() => {
            resolvedRef.current.indeterminate = indeterminate
        }, [resolvedRef, indeterminate])

        return (
            <>
                <input type="checkbox" ref={resolvedRef} {...rest} />
            </>
        )
    }
)
IndeterminateCheckbox.displayName = 'IndeterminateCheckbox';

const HeaderSelectionComponent = ({ getToggleAllRowsSelectedProps }: any) => (
    <div>
        <IndeterminateCheckbox {...getToggleAllRowsSelectedProps()} />
    </div>
)

const CellSelectionComponent = ({ row }: any) => (
    <div>
        <IndeterminateCheckbox {...row.getToggleRowSelectedProps()} />
    </div>
)

export default JsonTable
