import styles from './table.module.scss'
import React, { useState, useRef, useContext } from 'react'
import { Cell, useRowSelect, useTable } from 'react-table'
import toast from 'react-hot-toast'
import { DBConnection, DBQueryData, Tab } from '../../../data/models'
import EditableCell from './editablecell'
import AddModal from './addmodal'
import ConfirmModal from '../../widgets/confirmModal'
import { useAppDispatch, useAppSelector } from '../../../redux/hooks'
import { deleteDBData, setQueryData, updateDBSingleData } from '../../../redux/dataModelSlice'
import { DBConnType } from '../../../data/defaults'
import TabContext from '../../layouts/tabcontext'


type TablePropType = {
    queryData: DBQueryData,
    dbConnection: DBConnection
    mSchema: string,
    mName: string,
    isEditable: boolean,
    showHeader?: boolean,
    querySort?: string[],
    onFilterChanged: (newFilter: string[] | undefined) => void,
    onSortChanged: (newSort: string[] | undefined) => void,
}

const Table = ({ queryData, dbConnection, mSchema, mName, isEditable, showHeader, querySort, onFilterChanged, onSortChanged }: TablePropType) => {

    const dispatch = useAppDispatch()

    const activeTab: Tab = useContext(TabContext)!

    const [editCell, setEditCell] = useState<(string | number)[]>([])
    const [isAdding, setIsAdding] = useState<boolean>(false)
    const [isDeleting, setIsDeleting] = useState<boolean>(false)

    const [filterValue, setFilterValue] = useState<string[]>(['default', 'default', ''])

    const data = React.useMemo(
        () => queryData.rows,
        [queryData]
    )

    const displayColumns = queryData.columns.filter(col => col !== 'ctid')
    const ctidExists = queryData.columns.length !== displayColumns.length

    const columns = React.useMemo(
        () => displayColumns.map((col, i) => ({
            Header: <>{col}{querySort && querySort[0] === col ?
                querySort[1] === 'ASC' ?
                    <>&nbsp;<i className="fas fa-caret-up" /></>
                    :
                    <>&nbsp;<i className="fas fa-caret-down" /></>
                : undefined}</>,
            accessor: (ctidExists ? i + 1 : i).toString(),
        })),
        [queryData, querySort]
    )

    const defaultColumn = {
        Cell: EditableCell,
    }

    const resetEditCell = () => {
        setEditCell([])
    }

    const onSaveCell = async (ctid: string, columnIdx: string, newValue: string) => {
        const columnName = queryData.columns[parseInt(columnIdx)]
        const result = await dispatch(updateDBSingleData({ dbConnectionId: dbConnection.id, schemaName: mSchema, name: mName, id: ctid, columnName, newValue, columnIdx })).unwrap()
        if (result.success) {
            const rowIdx = queryData!.rows.findIndex(x => x["0"] === ctid)
            if (rowIdx) {
                const newQueryData: DBQueryData = { ...queryData!, rows: [...queryData!.rows] }
                console.log(result, newQueryData)
                newQueryData!.rows[rowIdx] = { ...newQueryData!.rows[rowIdx], ctid: result.data.ctid }
                newQueryData!.rows[rowIdx][columnIdx] = newValue
                dispatch(setQueryData({ data: newQueryData, tabId: activeTab.id }))
            } else {
                // fetchData(false)
            }
            resetEditCell()
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
    } = useTable<any>({
        columns,
        data,
        defaultColumn,
        ...{ editCell, resetEditCell, onSaveCell }
    },
        useRowSelect,
        hooks => {
            if (isEditable)
                hooks.visibleColumns.push(columns => [
                    {
                        id: 'selection',
                        Header: HeaderSelectionComponent,
                        Cell: CellSelectionComponent,
                    },
                    ...columns,
                ]
                )
        }
    )

    const newState: any = state // temporary typescript hack
    const selectedRowIds: any = newState.selectedRowIds
    const selectedRows: number[] = Object.keys(selectedRowIds).map(x => parseInt(x))
    const selectedCTIDs = rows.filter((_, i) => selectedRows.includes(i)).map(x => x.original['0']).filter(x => x)

    const deleteRows = async () => {
        if (selectedCTIDs.length > 0) {
            const result = await dispatch(deleteDBData({ dbConnectionId: dbConnection.id, schemaName: mSchema, name: mName, selectedIDs: selectedCTIDs })).unwrap()
            if (result.success) {
                toast.success('rows deleted')
                const filteredRows = queryData!.rows.filter((_, i) => !selectedRows.includes(i))
                const newQueryData: DBQueryData = { ...queryData!, rows: filteredRows }
                dispatch(setQueryData({ data: newQueryData, tabId: activeTab.id }))
            } else {
                toast.error(result.error!)
            }
        }
        setIsDeleting(false)
    }

    const startEditing = (cell: Cell<any, any>) => {
        if (isEditable)
            setEditCell([cell.row.index, cell.column.id])
    }

    const onFilter = () => {
        let filter: string[] | undefined = undefined
        if (filterValue[0] !== 'default' && filterValue[1] !== 'default') {
            let operator = filterValue[1]
            if (operator === 'IS NULL' || operator === 'IS NOT NULL') {
                filter = [filterValue[0], operator]
            } else {
                filter = [filterValue[0], operator, filterValue[2]]
            }
        }
        onFilterChanged(filter)
    }

    const changeSort = (newSortIdx: string) => {
        if (!isEditable) {
            return
        }
        const newSortName: string = displayColumns.find((_, i) => {
            const colIdx = ctidExists ? i + 1 : i
            return colIdx.toString() === newSortIdx
        })!
        if (querySort && newSortName === querySort[0]) {
            if (querySort[1] === 'ASC') {
                onSortChanged([querySort[0], 'DESC'])
            } else if (querySort[1] === 'DESC') {
                onSortChanged(undefined)
            }
        } else {
            onSortChanged([newSortName, 'ASC'])
        }
    }

    const handleFilterChange = (e: React.ChangeEvent<HTMLInputElement | HTMLSelectElement>, index: number) => {
        const value = e.target.value
        const newFilterValue = [...filterValue]
        newFilterValue[index] = value
        setFilterValue(newFilterValue)
    }

    return (
        <React.Fragment>
            {(showHeader || isEditable) && <div className={styles.tableHeader}>
                <div className="columns">
                    <div className="column is-9">
                        <div className="field has-addons">
                            <p className="control">
                                <span className="select">
                                    <select value={filterValue[0]} onChange={e => handleFilterChange(e, 0)}>
                                        <option value="default">Select column</option>
                                        {displayColumns.map(x =>
                                            (<option key={x}>{x}</option>)
                                        )}
                                    </select>
                                </span>
                            </p>
                            <p className="control">
                                <span className="select">
                                    <select value={filterValue[1]} onChange={e => handleFilterChange(e, 1)}>
                                        <option value="default">Select operator</option>
                                        <option value="=">=</option>
                                        <option value="!=">≠</option>
                                        <option value="<">&lt;</option>
                                        <option value=">">&gt;</option>
                                        <option value=">=">≥</option>
                                        <option value="<=">≤</option>
                                        <option value="IS NULL">is null</option>
                                        <option value="IS NOT NULL">not null</option>
                                        <option value="LIKE">like</option>
                                        <option value="NOT LIKE">not like</option>
                                    </select>
                                </span>
                            </p>
                            <p className="control">
                                <input className="input" type="text" placeholder="Value" value={filterValue[2]} onChange={e => handleFilterChange(e, 2)} />
                            </p>
                            <p className="control">
                                <button className="button" onClick={onFilter}>Filter</button>
                            </p>
                        </div>
                    </div>
                    {isEditable && <React.Fragment>
                        <div className="column is-3 is-flex is-justify-content-flex-end">
                            <button className="button" disabled={dbConnection.type === DBConnType.MYSQL || selectedCTIDs.length === 0} onClick={() => { setIsDeleting(true) }}>
                                <span className="icon is-small">
                                    <i className="fas fa-trash" />
                                </span>
                            </button>
                            &nbsp;&nbsp;
                            <button className="button is-primary" onClick={() => { setIsAdding(true) }}>
                                <span className="icon is-small">
                                    <i className="fas fa-plus" />
                                </span>
                            </button>
                        </div>
                    </React.Fragment>}
                </div>
            </div>}
            {isAdding &&
                <AddModal
                    queryData={queryData}
                    dbConnection={dbConnection}
                    mSchema={mSchema}
                    mName={mName}
                    onClose={() => { setIsAdding(false) }} />
            }
            {isDeleting && <ConfirmModal
                message={`Are you sure you want to delete selected rows?`}
                onConfirm={deleteRows}
                onClose={() => { setIsDeleting(false) }} />}
            <div className="table-container">
                <table {...getTableProps()} className={"table is-bordered is-striped is-narrow is-hoverable is-fullwidth"}>
                    <thead>
                        {headerGroups.map(headerGroup => (
                            <tr {...headerGroup.getHeaderGroupProps()} key={"header"}>
                                {headerGroup.headers.map(column => (
                                    <th {...column.getHeaderProps()} key={column.id} onClick={() => { changeSort(column.id) }}>
                                        {column.render('Header')}
                                    </th>
                                ))}
                            </tr>
                        ))}
                    </thead>
                    <tbody {...getTableBodyProps()}>
                        {rows.map(row => {
                            prepareRow(row)
                            const selectedRow: any = row // temp type hack 
                            return (
                                <tr {...row.getRowProps()} key={row.id} className={selectedRow.isSelected ? 'is-selected' : ''}>
                                    {row.cells.map(cell => {
                                        return (<td {...cell.getCellProps()} onDoubleClick={() => startEditing(cell)} key={row.id + "" + cell.column.id}>
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


export default Table