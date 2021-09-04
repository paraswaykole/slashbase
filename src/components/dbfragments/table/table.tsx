import styles from './table.module.scss'
import React, { useState } from 'react'
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
    heading?: string,
    updateCellData: (oldCtid: string, newCtid: string, columnName: string, newValue: string|null|boolean)=>void,
    onDeleteRows: (indexes: number[]) => void,
    onAddData: (newData: any) => void,
}

const Table = ({queryData, dbConnection, mSchema, mName, isEditable, heading, updateCellData, onDeleteRows, onAddData}: ProjectCardPropType) => {

    const [editCell, setEditCell] = useState<(string|number)[]>([])
    const [isAdding, setIsAdding] = useState<boolean>(false)

    const data = React.useMemo(
        () => queryData.rows,
        [queryData]
    )

    const columns = React.useMemo(
        () => queryData.columns.filter(col => col !== 'ctid').map((col) => ({
            Header: col,
            accessor: col, 
        })),
        [queryData]
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
                        Header: ({ getToggleAllRowsSelectedProps }: any) => (
                            <div>
                                <IndeterminateCheckbox {...getToggleAllRowsSelectedProps()} />
                            </div>
                        ),
                        Cell: ({ row }: any) => (
                            <div>
                                <IndeterminateCheckbox {...row.getToggleRowSelectedProps()} />
                            </div>
                        ),
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
   
    return (
        <React.Fragment>
            { (heading || isEditable) && <div className={styles.tableHeader}>
                <div className="columns">
                    <div className="column is-9">
                        {heading && <h1>{heading}</h1> }
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
                            <tr {...headerGroup.getHeaderGroupProps()} key={headerGroup.id}>
                                {headerGroup.headers.map(column => (
                                    <th {...column.getHeaderProps()} key={column.id}>
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

const IndeterminateCheckbox = React.forwardRef<HTMLInputElement, {
    indeterminate: boolean,
}>(
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
  


export default Table