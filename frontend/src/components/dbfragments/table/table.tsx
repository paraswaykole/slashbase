import styles from './table.module.scss'
import React, { useState, useRef } from 'react'
import { Cell, useRowSelect, useTable, UseTableInstanceProps } from 'react-table'
import toast from 'react-hot-toast';
import { DBConnection, DBQueryData } from '../../../data/models'
import EditableCell from './editablecell'
import apiService from '../../../network/apiService'
import AddModal from './addmodal';


type ProjectCardPropType = { 
    queryData: DBQueryData,
    dbConnection: DBConnection
    mSchema: string,
    mName: string,
    isEditable: boolean,
    showHeader?: boolean,
    querySort?: string[],
    updateCellData: (oldCtid: string, newCtid: string, columnName: string, newValue: string|null|boolean)=>void,
    onDeleteRows: (indexes: number[]) => void,
    onAddData: (newData: any) => void,
    onFilterChanged: (newFilter: string[]|undefined) => void,
    onSortChanged: (newSort: string[]|undefined) => void,
}

const Table = ({queryData, dbConnection, mSchema, mName, isEditable, showHeader, querySort, updateCellData, onDeleteRows, onAddData, onFilterChanged, onSortChanged}: ProjectCardPropType) => {

    const [editCell, setEditCell] = useState<(string|number)[]>([])
    const [isAdding, setIsAdding] = useState<boolean>(false)

    const filter0Ref = useRef<HTMLSelectElement>(null);
    const filter1Ref = useRef<HTMLSelectElement>(null);
    const filter2Ref = useRef<HTMLInputElement>(null);

    const data = React.useMemo(
        () => queryData.rows,
        [queryData]
    )

    const displayColumns = queryData.columns.filter(col => col !== 'ctid')

    const columns = React.useMemo(
        () => displayColumns.map((col) => ({
            Header: <>{col}{querySort && querySort[0] === col ? 
                querySort[1] === 'ASC' ? 
                <>&nbsp;<i className="fas fa-caret-up"/></>
                :
                <>&nbsp;<i className="fas fa-caret-down"/></>
                 : undefined}</>,
            accessor: col, 
        })),
        [queryData, querySort]
    )

    const defaultColumn = {
        Cell: EditableCell,
    }      

    const resetEditCell = ()=> {
        setEditCell([])
    }

    const onSaveCell = async (ctid: string, columnName: string, newValue: string) => {
        const result = await apiService.updateDBSingleData(dbConnection.id, mSchema, mName, ctid, columnName, newValue)
        if (result.success) {
            updateCellData(ctid, result.data.ctid, columnName, newValue)
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
    const selectedRows: number[] = Object.keys(selectedRowIds).map(x=>parseInt(x))
    const selectedCTIDs = rows.filter((_,i) => selectedRows.includes(i)).map(x => x.original['ctid']).filter(x => x)

    const onDeleteBtnPressed = async () => {
        if (selectedCTIDs.length > 0){
            const result = await apiService.deleteDBData(dbConnection.id, mSchema, mName, selectedCTIDs)
            if (result.success) {
                toast.success('rows deleted');
                onDeleteRows(selectedRows)
            } else {
                toast.error(result.error!);
            }
        }
    }

    const startEditing = (cell: Cell<any, any>)=>{
    if (isEditable)
        setEditCell([cell.row.index, cell.column.id])
    }

    const changeFilter = () => {
        let filter: string[] | undefined = undefined
        if (filter0Ref.current!.value !== 'default' && filter1Ref.current!.value !== 'default'){
            let operator = filter1Ref.current!.value
            if (operator === 'IS NULL' || operator === 'IS NOT NULL' ){
                filter = [filter0Ref.current!.value, operator]
            } else {
                filter = [filter0Ref.current!.value, operator, filter2Ref.current!.value]
            }
        }
        onFilterChanged(filter)
    }

    const changeSort = (newSort: string) => {
        if (!isEditable){
            return
        }
        if (querySort && newSort === querySort[0]) {
            if (querySort[1] === 'ASC') {
                onSortChanged([querySort[0], 'DESC'])
            } else if (querySort[1] === 'DESC') {
                onSortChanged(undefined)
            }
        } else {
            onSortChanged([newSort, 'ASC'])
        }
    }
   
    return (
        <React.Fragment>
            { (showHeader || isEditable) && <div className={styles.tableHeader}>
                <div className="columns">
                    <div className="column is-9">
                        <div className="field has-addons">
                            <p className="control">
                                <span className="select">
                                    <select ref={filter0Ref}>
                                        <option value="default">Select column</option>
                                        {displayColumns.map(x => 
                                            (<option key={x}>{x}</option>)
                                        )}
                                    </select>
                                </span>
                            </p>
                            <p className="control">
                                <span className="select">
                                    <select ref={filter1Ref}>
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
                                <input ref={filter2Ref} className="input" type="text" placeholder="Value"/>
                            </p>
                            <p className="control">
                                <button className="button" onClick={changeFilter}>Filter</button>
                            </p>
                        </div>
                    </div>
                    {isEditable && <React.Fragment>
                        <div className="column is-3 is-flex is-justify-content-flex-end">
                            <button className="button" disabled={selectedCTIDs.length===0} onClick={onDeleteBtnPressed}>
                                <span className="icon is-small">
                                    <i className="fas fa-trash"/>
                                </span>
                            </button>
                            &nbsp;&nbsp;
                            <button className="button is-primary" onClick={()=>{setIsAdding(true)}}>
                                <span className="icon is-small">
                                    <i className="fas fa-plus"/>
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
                    onClose={()=>{setIsAdding(false)}}
                    onAddData={onAddData}/>
            }
            <div className="table-container">
                <table {...getTableProps()} className={"table is-bordered is-striped is-narrow is-hoverable is-fullwidth"}>
                    <thead>
                        {headerGroups.map(headerGroup => (
                            <tr {...headerGroup.getHeaderGroupProps()} key={"header"}>
                                {headerGroup.headers.map(column => (
                                    <th {...column.getHeaderProps()} key={column.id} onClick={()=>{changeSort(column.id)}}>
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
                                <tr {...row.getRowProps()} key={row.id} className={selectedRow.isSelected?'is-selected':''}>
                                    {row.cells.map(cell => {
                                        return (<td {...cell.getCellProps()} onDoubleClick={()=>startEditing(cell)} key={row.id+""+cell.column.id}>
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